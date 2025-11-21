package js

import (
	"fmt"
	"io/fs"
	"syscall/js"
)

var _ fs.DirEntry = new(jsGoDirEntry)

func NewDirEntry(v js.Value) fs.DirEntry {
	return &jsGoDirEntry{v}
}

type jsGoDirEntry struct{ Value js.Value }

func (d *jsGoDirEntry) Name() string {
	return d.Value.Call("name").String()
}

func (d *jsGoDirEntry) IsDir() bool {
	return d.Value.Call("isDir").Bool()
}

func (d *jsGoDirEntry) Type() fs.FileMode {
	return 0
}

func (d *jsGoDirEntry) Info() (fs.FileInfo, error) {
	v := d.Value.Call("info")
	if v.IsNull() || v.IsUndefined() {
		return nil, fmt.Errorf("no info")
	}
	if v.InstanceOf(js.Global().Get("Error")) {
		return nil, fmt.Errorf(v.Get("message").String())
	}
	return &jsGoFileInfo{Value: v}, nil
}

var (
	_ fs.File        = new(jsGoReadDirFile)
	_ fs.ReadDirFile = new(jsGoReadDirFile)
)

func NewReadDirFile(v js.Value) fs.ReadDirFile {
	return &jsGoReadDirFile{
		File:  NewFile(v),
		Value: v,
	}
}

type jsGoReadDirFile struct {
	fs.File
	Value js.Value
}

func (f *jsGoReadDirFile) ReadDir(n int) ([]fs.DirEntry, error) {
	rd := f.Value.Get("readDir")
	if rd.IsUndefined() || rd.IsNull() {
		return nil, fs.ErrInvalid
	}
	res := f.Value.Call("readDir", n)
	l := res.Length()
	entries := make([]fs.DirEntry, 0, int(l))
	for i := 0; i < int(l); i++ {
		entries = append(entries, NewDirEntry(res.Index(i)))
	}
	return entries, nil
}

var _ fs.FS = new(jsGoFS)

func NewFS(v js.Value) fs.FS {
	return &jsGoFS{v}
}

type jsGoFS struct {
	Value js.Value
}

func (s *jsGoFS) Open(p string) (fs.File, error) {
	res := s.Value.Call("open", p)
	if res.IsNull() || res.IsUndefined() {
		return nil, fs.ErrNotExist
	}
	if res.InstanceOf(js.Global().Get("Error")) {
		return nil, fmt.Errorf(res.Get("message").String())
	}
	if res.Call("stat").Call("isDir").Bool() {
		return NewReadDirFile(res), nil
	}
	return NewFile(res), nil
}
