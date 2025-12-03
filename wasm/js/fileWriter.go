package js

import (
	"fmt"
	"io/fs"
	"syscall/js"

	"github.com/louisbuchbinder/core/wasm"
)

type goFileWriter struct {
	File  fs.File
	Value js.Value
	Data  chan js.Value
}

func NewFileWriter(v js.Value) wasm.OpfsFile {
	return &goFileWriter{File: NewFile(v), Value: v, Data: make(chan js.Value, 1)}
}

func (w *goFileWriter) Stat() (fs.FileInfo, error) {
	return w.File.Stat()
}

func (w *goFileWriter) Read(b []byte) (int, error) {
	return w.File.Read(b)
}

func (w *goFileWriter) Write(p []byte) (int, error) {
	if w == nil || w.Value.IsUndefined() || w.Value.IsNull() {
		return 0, fmt.Errorf("invalid writer")
	}

	ta := js.Global().Get("Uint8Array").New(len(p))
	js.CopyBytesToJS(ta, p)
	v, err := PromiseResolveOrReject(w.Value.Call("write", ta))
	if err != nil {
		return 0, err
	}
	if v.Type() == js.TypeNumber {
		return v.Int(), nil
	}
	return 0, fmt.Errorf("unexpected write result: %v", v)
}

func (w *goFileWriter) Close() error {
	if w == nil || w.Value.IsUndefined() || w.Value.IsNull() {
		return fmt.Errorf("invalid writer")
	}
	_, err := PromiseResolveOrReject(w.Value.Call("close"))
	return err
}
