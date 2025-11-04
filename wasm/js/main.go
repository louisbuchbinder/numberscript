package js

import (
	"syscall/js"

	"github.com/louisbuchbinder/core/wasm"
)

func ToValues(args []js.Value) []wasm.Value {
	values := make([]wasm.Value, len(args))
	for i, arg := range args {
		values[i] = arg
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
