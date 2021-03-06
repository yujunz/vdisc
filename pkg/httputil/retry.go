// Copyright © 2019 NVIDIA Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httputil

import (
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
	"go.uber.org/zap"
)

type RetryOptions struct {
	InitialInterval time.Duration
	MaxInterval     time.Duration
	MaxElapsedTime  time.Duration
}

func WithRetries(base http.RoundTripper, options ...RetryOptions) http.RoundTripper {
	t := &retryingRoundTripper{rt: base}
	if len(options) > 0 {
		t.opts = options[0]
	}
	return t
}

type retryFunc func() (retry bool)

func retry(doFn retryFunc, options ...RetryOptions) {
	b := backoff.NewExponentialBackOff()
	if len(options) > 0 {
		if options[0].MaxElapsedTime != 0 {
			b.MaxElapsedTime = options[0].MaxElapsedTime
		}
		if options[0].MaxInterval != 0 {
			b.MaxInterval = options[0].MaxInterval
		}
		if options[0].InitialInterval != 0 {
			b.InitialInterval = options[0].InitialInterval
		}
	}
	for {
		retry := doFn()
		delay := b.NextBackOff()
		if !retry || delay == backoff.Stop {
			return
		}
		time.Sleep(delay)
	}
}

type retryingRoundTripper struct {
	rt   http.RoundTripper
	opts RetryOptions
}

func (r *retryingRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	retry(func() (retry bool) {
		resp, err = r.rt.RoundTrip(req)
		retry = err != nil || resp.StatusCode >= 500 || resp.StatusCode == 429
		if retry {
			if err != nil {
				warnf("retrying %s: %s\n", req.URL.String(), err)
			} else {
				warnf("retrying %s: http %d\n", req.URL.String(), resp.StatusCode)
			}
		}
		return
	}, r.opts)
	return
}

func warnf(v string, a ...interface{}) {
	zap.L().Sugar().Named("httputil").Warnf(v, a...)
}
