package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/louisbuchbinder/core/lib/build"
	"github.com/louisbuchbinder/core/lib/util"
)

var dir string

func init() {
	flag.StringVar(&dir, "dir", "", "[required] the dir path to generate manifest entries")
	flag.Parse()
	if dir == "" {
		flag.Usage()
		os.Exit(1)
	}
}

type File struct {
	Path string
	File fs.File
}

func getFiles() ([]*File, error) {
	r, err := os.OpenRoot(dir)
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Close() }()
	stat, err := r.Stat(".")
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("unexpected input dir is not a directory")
	}
	var files []*File
	filesystem := r.FS()
	if err := fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			file, err := filesystem.Open(path)
			if err != nil {
				return err
			}
			files = append(files, &File{
				Path: filepath.Join(dir, path),
				File: file,
			})
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return files, nil
}

func processFile(f *File) error {
	defer func() { _ = f.File.Close() }()
	sha256sum, err := util.Sha256HexOfFile(f.File)
	if err != nil {
		return err
	}
	src, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer func() { _ = src.Close() }()
	filename := filepath.Base(f.Path)
	sha256Filename := fmt.Sprintf("%s.%s", sha256sum, filename)
	filedir := filepath.Dir(f.Path)
	dest, err := os.Create(filepath.Join(filedir, sha256Filename))
	if err != nil {
		return err
	}
	defer func() { _ = dest.Close() }()
	n, err := io.Copy(dest, src)
	if err != nil {
		return err
	}
	stat, err := f.File.Stat()
	if err != nil {
		return err
	}
	if n != stat.Size() {
		return fmt.Errorf("failed to write file copy, wrote '%d', expected '%d'", n, stat.Size())
	}
	manifestPath := filepath.Join(filedir, "manifest.jsonl")
	manifest, err := os.OpenFile(manifestPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	b, err := json.Marshal(&build.ManifestData{
		Filename:       filename,
		Sha256Filename: sha256Filename,
	})
	if err != nil {
		return err
	}
	n1, err := manifest.Write(append(b, '\n'))
	if err != nil {
		return err
	}
	if n1 != len(b)+1 {
		return fmt.Errorf("failed to write manifest entry, wrote '%d', expected '%d'", n1, len(b)+1)
	}
	return nil
}

func main() {
	files := util.Must(getFiles())
	for _, file := range files {
		util.Must0(processFile(file))
	}
}
