package wasmsha1

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/louisbuchbinder/core/wasm"
)

func Sum(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha1.Sum([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}
