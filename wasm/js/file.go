package js

import (
	"fmt"
	"io"
	"io/fs"
	"syscall/js"
	"time"
)

var _ fs.FileInfo = new(jsGoFileInfo)

type jsGoFileInfo struct{ Value js.Value }

func (info *jsGoFileInfo) Name() string      { return info.Value.Call("name").String() }
func (info *jsGoFileInfo) Size() int64       { return int64(info.Value.Call("size").Int()) }
func (info *jsGoFileInfo) Mode() fs.FileMode { return 0o644 }
func (info *jsGoFileInfo) ModTime() time.Time {
	return time.UnixMilli(int64(info.Value.Call("modTime").Int()))
}
func (info *jsGoFileInfo) IsDir() bool { return info.Value.Call("isDir").Bool() }
func (info *jsGoFileInfo) Sys() any    { return nil }

func NewFile(v js.Value) fs.File {
	return &jsGoFile{
		Value:      v,
		Data:       make(chan js.Value, 1),
		InProgress: false,
	}
}

var _ fs.File = new(jsGoFile)

type jsGoFile struct {
	Value      js.Value
	Data       chan js.Value
	InProgress bool
}

func (f *jsGoFile) Stat() (fs.FileInfo, error) {
	return &jsGoFileInfo{f.Value.Call("stat")}, nil
}

func (f *jsGoFile) Read(b []byte) (int, error) {
	var handler, errHandler js.Func
	fn := func(this js.Value, args []js.Value) any {
		defer handler.Release()
		defer errHandler.Release()
		defer func() { f.InProgress = false }()
		if len(args) < 1 {
			f.Data <- js.Undefined()
		} else {
			f.Data <- args[0]
		}
		return nil
	}
	handler = js.FuncOf(fn)
	errHandler = js.FuncOf(fn)

	if !f.InProgress {
		f.InProgress = true
		w := wrapper{f.Value.Call("read")}
		if !w.IsPromise() {
			return 0, fmt.Errorf("expected a Promise in jsGoFile.Read")
		}
		_ = w.Call("then", handler, errHandler)
		return 0, nil
	}

	select {
	case <-time.NewTicker(time.Nanosecond).C:
		return 0, nil
	case d := <-f.Data:
		w := wrapper{d}
		if w.Value.IsNull() {
			return 0, io.EOF
		}
		if w.Value.InstanceOf(js.Global().Get("Error")) {
			return 0, fmt.Errorf(w.Get("message").String())
		}
		if !w.IsUint8Array() {
			return 0, fmt.Errorf("unexpected data returned from GoFile.read, expected Uint8Array")
		}
		return copy(b, w.Uint8Array()), nil
	}
}

func (f *jsGoFile) Close() error {
	v := f.Value.Call("close")
	if !v.IsNull() {
		return fmt.Errorf("failed to close jsGoFile")
	}
	return nil
}
