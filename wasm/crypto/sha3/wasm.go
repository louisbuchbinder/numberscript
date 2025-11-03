package wasmsha3

import (
	"crypto/sha3"
	"encoding/hex"

	"github.com/louisbuchbinder/core/wasm"
)

func Sum224(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha3.Sum224([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum256(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha3.Sum256([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum384(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha3.Sum384([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum512(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha3.Sum512([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func SumSHAKE128(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	v := sha3.SumSHAKE128([]byte(args[0].String()), args[1].Int())
	return hex.EncodeToString(v[:])
}

func SumSHAKE256(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	v := sha3.SumSHAKE256([]byte(args[0].String()), args[1].Int())
	return hex.EncodeToString(v[:])
}
