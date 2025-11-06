package js

import (
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
	// TODO other arrays here
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
