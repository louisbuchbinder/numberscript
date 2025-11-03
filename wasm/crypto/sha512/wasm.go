package wasmsha512

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/louisbuchbinder/core/wasm"
)

func Sum384(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha512.Sum384([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum512(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha512.Sum512([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum512_224(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha512.Sum512_224([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum512_256(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha512.Sum512_256([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}
