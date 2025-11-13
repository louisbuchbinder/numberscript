package wasmadler32

import (
	"encoding/hex"
	"hash/adler32"

	"github.com/louisbuchbinder/core/wasm"
)

func Checksum(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	data, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	return int(adler32.Checksum(data))
}
