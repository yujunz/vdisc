package vdisc_cli

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/NVIDIA/vdisc/pkg/iso9660"
	"github.com/NVIDIA/vdisc/pkg/vdisc"
)

// CsvCmd type
type CsvCmd struct {
	Url  string `short:"u" help:"The URL of the vdisc" required:"true"`
	Path string `short:"p" help:"The path in the vdisc to list" default:"/"`
}

// Run CsvCmd
func (cmd *CsvCmd) Run(globals *Globals) error {
	v, err := vdisc.Load(cmd.Url, globalCache(&globals.Cache))
	if err != nil {
		zap.L().Fatal("loading vdisc", zap.Error(err))
	}
	defer v.Close()

	w := iso9660.NewWalker(v.Image())
	cmd.printCsv(w, cmd.Path, v)

	return nil
}

func (cmd *CsvCmd) printCsv(walker *iso9660.Walker, path string, urlMap extentURLMapper) {
	finfos, err := walker.ReadDir(path)
	if err != nil {
		zap.L().Fatal("", zap.Error(err))
	}

	for _, fi := range finfos {
		name := fi.Name()
		if name == "." || name == ".." {
			continue
		}
		abspath := filepath.Join(path, name)

		if fi.IsDir() {
			cmd.printCsv(walker, abspath, urlMap)
		} else if fi.Mode()&os.ModeSymlink != 0 {
			zap.L().Fatal("Symlink not supported")
		} else {
			url, err := urlMap.ExtentURL(fi.Extent())
			if err != nil {
				zap.L().Fatal("extent url lookup", zap.Uint32("lba", uint32(fi.Extent())), zap.Error(err))
			}
			fmt.Println(fmt.Sprintf("%q,%q,%d", abspath, url, fi.Size()))
		}
	}
}
