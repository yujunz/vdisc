// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockdriver

import context "context"
import driver "github.com/NVIDIA/vdisc/pkg/storage/driver"
import mock "github.com/stretchr/testify/mock"
import os "os"

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, url
func (_m *Driver) Create(ctx context.Context, url string) (driver.ObjectWriter, error) {
	ret := _m.Called(ctx, url)

	var r0 driver.ObjectWriter
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.ObjectWriter); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.ObjectWriter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Open provides a mock function with given fields: ctx, url, size
func (_m *Driver) Open(ctx context.Context, url string, size int64) (driver.Object, error) {
	ret := _m.Called(ctx, url, size)

	var r0 driver.Object
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) driver.Object); ok {
		r0 = rf(ctx, url, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, url, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readdir provides a mock function with given fields: ctx, url
func (_m *Driver) Readdir(ctx context.Context, url string) ([]os.FileInfo, error) {
	ret := _m.Called(ctx, url)

	var r0 []os.FileInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) []os.FileInfo); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, url
func (_m *Driver) Remove(ctx context.Context, url string) error {
	ret := _m.Called(ctx, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stat provides a mock function with given fields: ctx, url
func (_m *Driver) Stat(ctx context.Context, url string) (os.FileInfo, error) {
	ret := _m.Called(ctx, url)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) os.FileInfo); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
