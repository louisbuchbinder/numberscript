package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/louisbuchbinder/core/lib/util"
)

func extract(tempdir string, distTarball io.Reader) error {
	gz, err := gzip.NewReader(distTarball)
	if err != nil {
		return err
	}
	defer gz.Close()
	r := tar.NewReader(gz)
	for {
		header, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		target := filepath.Join(tempdir, header.Name)
		switch header.Typeflag {
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o700); err != nil {
				return err
			}
			f := util.Must(os.Create(target))
			n := util.Must(io.Copy(f, r))
			if n != header.Size {
				return fmt.Errorf("expected to write %d bytes, but wrote %d bytes", header.Size, n)
			}
			if err := f.Close(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown type: %v in %s", header.Typeflag, header.Name)
		}
	}
	return nil
}

var (
	distRlocation string
	port          int
)

func init() {
	flag.IntVar(&port, "port", 8000, "Port to serve on")
	flag.Parse()
}

func main() {
	distTarball := util.Must(os.Open(util.Must(runfiles.Rlocation(distRlocation))))
	tempdir, cleanup, err := util.MkdirTemp("", "louisbuchbinder.com-*")
	util.Must0(err)
	defer func() {
		err := cleanup()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to clean up temp dir %s, with error: %s\n", tempdir, err.Error())
		}
	}()
	util.Must0(extract(tempdir, distTarball))
	handler := http.FileServerFS(os.DirFS(tempdir))
	http.Handle("/", handler)
	fmt.Printf("Starting server on :%d\n", port)
	util.Must0(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
