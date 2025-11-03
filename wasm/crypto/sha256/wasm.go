package wasmsha256

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/louisbuchbinder/core/wasm"
)

func Sum224(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha256.Sum224([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum256(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha256.Sum256([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}
