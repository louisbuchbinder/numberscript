package wasmmd5

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/louisbuchbinder/core/wasm"
)

func Sum(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := md5.Sum([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}
