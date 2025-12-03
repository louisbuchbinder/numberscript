package main

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/louisbuchbinder/core/lib/build"
	"github.com/louisbuchbinder/core/lib/util"
	"golang.org/x/sync/errgroup"
)

var tarReaders []*tar.Reader

var (
	archives stringSliceFlag
	output   string
)

type stringSliceFlag []string

func (s *stringSliceFlag) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func initArchives() error {
	for _, a := range archives {
		f, err := os.Open(a)
		if err != nil {
			return err
		}
		tarReaders = append(tarReaders, tar.NewReader(f))

	}
	return nil
}

func init() {
	flag.Var(&archives, "archive", "[required] the archives to be used to generate the manifest")
	flag.StringVar(&output, "output", "", "[required] the output manifest file")
	flag.Parse()
	if len(archives) == 0 ||
		output == "" {
		flag.Usage()
		os.Exit(1)
	}
	util.Must0(initArchives())
}

func process() (map[string]build.ManifestData, error) {
	c := make(chan build.ManifestData, 1)
	group := new(errgroup.Group)
	group.Go(
		func() error {
			defer close(c)
			for _, r := range tarReaders {
				_ = r
				for {
					h, err := r.Next()
					if err == io.EOF {
						break
					}
					if err != nil {
						return err
					}
					if filepath.Base(h.Name) == "manifest.jsonl" {
						dir := filepath.Dir(h.Name)
						b := make([]byte, h.Size)
						n, err := r.Read(b)
						if err != nil && err != io.EOF {
							return err
						}
						if n != int(h.Size) {
							return fmt.Errorf("failed to read entire file, expected '%d', got '%d", h.Size, n)
						}
						for _, line := range bytes.Split(b, []byte("\n")) {
							if len(line) == 0 {
								continue
							}
							var d build.ManifestData
							if err := json.Unmarshal(line, &d); err != nil {
								return err
							}
							d.Filename = filepath.Join("/", dir, d.Filename)
							d.Sha256Filename = filepath.Join("/", dir, d.Sha256Filename)
							c <- d
						}
					}
				}
			}
			return nil
		},
	)
	manifest := map[string]build.ManifestData{}
	for d := range c {
		if _, ok := manifest[d.Filename]; ok {
			return nil, fmt.Errorf("unexpected duplicate filename: '%s'", d.Filename)
		}
		manifest[d.Filename] = d
	}
	return manifest, group.Wait()
}

func main() {
	data := util.Must(process())
	b := util.Must(json.Marshal(data))
	util.Must0(os.WriteFile(output, b, 0o644))
}
