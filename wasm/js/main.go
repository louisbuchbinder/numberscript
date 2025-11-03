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
