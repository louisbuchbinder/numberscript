package js

import (
	"fmt"
	"io"
	"io/fs"
	"syscall/js"

	"github.com/louisbuchbinder/core/wasm"
)

type wrapper struct {
	js.Value
}

func (w wrapper) IsArray() bool {
	return js.Global().Get("Array").Call("isArray", w.Value).Bool()
}

func (w wrapper) IsUint8Array() bool {
	return w.InstanceOf(js.Global().Get("Uint8Array"))
}

func (w wrapper) Uint8Array() []uint8 {
	if !w.IsUint8Array() {
		return nil
	}
	b := make([]uint8, w.Value.Length())
	n := js.CopyBytesToGo(b, w.Value)
	return b[:n]
}

func (w wrapper) Array() []any {
	if !w.IsArray() {
		return nil
	}
	n := w.Value.Length()
	arr := make([]any, n)
	for i := range n {
		arr[i] = w.Value.Index(i)
	}
	return arr
}

func (w wrapper) IsFile() bool {
	return w.InstanceOf(js.Global().Get("GoFile")) || w.InstanceOf(js.Global().Get("OpfsFile"))
}

func (w wrapper) File() (fs.File, error) {
	if !w.IsFile() {
		return nil, fmt.Errorf("expected js.Value wrapper to be GoFile or OpfsFile")
	}
	return NewFile(w.Value), nil
}

func (w wrapper) IsFileWriter() bool {
	return w.InstanceOf(js.Global().Get("OpfsFile"))
}

type OpfsFile interface {
	io.WriteCloser
	fs.File
}

func (w wrapper) FileWriter() (wasm.OpfsFile, error) {
	if !w.IsFileWriter() {
		return nil, fmt.Errorf("expected js.Value wrapper to be OpfsFile")
	}
	return NewFileWriter(w.Value), nil
}

func (w wrapper) IsFS() bool {
	return w.InstanceOf(js.Global().Get("GoFS"))
}

func (w wrapper) FS() (fs.FS, error) {
	if !w.IsFS() {
		return nil, fmt.Errorf("expected js.Value wrapper to be GoFS")
	}
	return NewFS(w.Value), nil
}

func (w wrapper) IsPromise() bool {
	return w.InstanceOf(js.Global().Get("Promise"))
}

func (w wrapper) Invoke(args ...any) wasm.Value {
	return wrapper{w.Value.Invoke(args...)}
}

func (w wrapper) Reject(err error) {
	_ = w.Value.Invoke(Error(err))
}

func ToValues(args []js.Value) []wasm.Value {
	values := make([]wasm.Value, len(args))
	for i, arg := range args {
		values[i] = wrapper{arg}
	}
	return values
}

func OrError(val any) any {
	if err, ok := val.(error); ok {
		return Error(err)
	}
	// TODO: other arrays here
	// TODO: actually can we return []any to avoid this
	if l, ok := val.([]int); ok {
		return Array(l)
	}
	// TODO support dicts here
	return val
}

func Array[T any](items []T) js.Value {
	arr := js.Global().Get("Array").New()
	for _, val := range items {
		arr.Call("push", val)
	}
	return arr
}

func Error(err error) js.Value {
	return js.Global().Get("Error").New(err.Error())
}
