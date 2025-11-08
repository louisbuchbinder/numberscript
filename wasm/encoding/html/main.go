package html

import (
	"html"

	"github.com/louisbuchbinder/core/wasm"
)

func EscapeString(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	return html.EscapeString(args[0].String())
}

func UnescapeString(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	return html.UnescapeString(args[0].String())
}
