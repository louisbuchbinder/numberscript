package wasmbase64

import (
	"encoding/base64"

	"github.com/louisbuchbinder/core/wasm"
)

func EncodeToString(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := base64.StdEncoding.EncodeToString([]byte(args[0].String()))
	return string(v)
}

func DecodeString(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v, err := base64.StdEncoding.DecodeString(string(args[0].String()))
	if err != nil {
		return err
	}
	return string(v)
}
