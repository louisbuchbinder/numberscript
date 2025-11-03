package wasmhex

import (
	"encoding/hex"

	"github.com/louisbuchbinder/core/wasm"
)

func EncodeToString(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := hex.EncodeToString([]byte(args[0].String()))
	return string(v)
}

func DecodeString(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v, err := hex.DecodeString(string(args[0].String()))
	if err != nil {
		return err
	}
	return string(v)
}
