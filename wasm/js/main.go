package js

import (
	"fmt"
	"io"
	"io/fs"
	"syscall/js"
	"time"

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
	return w.InstanceOf(js.Global().Get("GoFS")) || w.InstanceOf(js.Global().Get("OpfsFS"))
}

func (w wrapper) FS() (fs.FS, error) {
	if !w.IsFS() {
		return nil, fmt.Errorf("expected js.Value wrapper to be GoFS or OpfsFS")
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

func (w wrapper) IsRequest() bool {
	return w.InstanceOf(js.Global().Get("Request"))
}

func (w wrapper) Request() (wasm.JsRequest, error) {
	if !w.IsRequest() {
		return nil, fmt.Errorf("js.Value is not an instanceof Request")
	}
	return &jsRequest{
		method: w.Get("method").String(),
		url:    w.Get("url").String(),
	}, nil
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

func PromiseResolveOrReject(promise js.Value) (js.Value, error) {
	var handler, errHandler js.Func
	c := make(chan js.Value, 1)
	fn := func(this js.Value, args []js.Value) any {
		defer handler.Release()
		defer errHandler.Release()
		if len(args) < 1 {
			c <- js.Global().Get("Error").New("got unexpected undefined from promise invocation")
		} else {
			c <- args[0]
		}
		return nil
	}
	handler = js.FuncOf(fn)
	errHandler = js.FuncOf(fn)
	promise.Call("then", handler, errHandler)
	ticker := time.NewTicker(time.Nanosecond)
	for {
		select {
		case <-ticker.C:
		case dat := <-c:
			if IsError(dat) {
				return js.Null(), fmt.Errorf(dat.Get("message").String())
			}
			return dat, nil
		}
	}
}

func IsError(maybeErr js.Value) bool {
	return maybeErr.InstanceOf(js.Global().Get("Error"))
}
