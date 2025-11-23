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

	var handler, errHandler js.Func
	fn := func(this js.Value, args []js.Value) any {
		defer handler.Release()
		defer errHandler.Release()
		if len(args) < 1 {
			w.Data <- js.Undefined()
		} else {
			w.Data <- args[0]
		}
		return nil
	}
	handler = js.FuncOf(fn)
	errHandler = js.FuncOf(fn)

	wr := wrapper{w.Value.Call("write", ta)}
	if !wr.IsPromise() {
		return 0, fmt.Errorf("expected a Promise in goFileWriter.Write")
	}
	_ = wr.Call("then", handler, errHandler)

	d := <-w.Data
	wv := wrapper{d}
	if wv.Value.IsNull() {
		return 0, nil
	}
	if wv.Value.InstanceOf(js.Global().Get("Error")) {
		return 0, fmt.Errorf(wv.Get("message").String())
	}
	if wv.Value.Type() == js.TypeNumber {
		return wv.Value.Int(), nil
	}
	return 0, fmt.Errorf("unexpected write result: %v", wv.Value)
}

func (w *goFileWriter) Close() error {
	if w == nil || w.Value.IsUndefined() || w.Value.IsNull() {
		return fmt.Errorf("invalid writer")
	}
	var handler, errHandler js.Func
	fn := func(this js.Value, args []js.Value) any {
		defer handler.Release()
		defer errHandler.Release()
		if len(args) < 1 {
			w.Data <- js.Undefined()
		} else {
			w.Data <- args[0]
		}
		return nil
	}
	handler = js.FuncOf(fn)
	errHandler = js.FuncOf(fn)

	wr := wrapper{w.Value.Call("close")}
	_ = wr.Call("then", handler, errHandler)

	d := <-w.Data
	wv := wrapper{d}
	if wv.Value.IsNull() || wv.Value.IsUndefined() {
		return nil
	}
	if wv.Value.InstanceOf(js.Global().Get("Error")) {
		return fmt.Errorf(wv.Get("message").String())
	}
	return nil
}
