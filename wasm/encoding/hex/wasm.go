package wasmhex

import (
	"encoding/hex"
	"fmt"
	"unicode/utf8"

	"github.com/louisbuchbinder/core/lib/util"
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
	if !utf8.Valid(v) {
		return fmt.Errorf("result contains invalid UTF-8; decode as bytes instead")
	}
	return string(v)
}

func DecodeStringAsBytes(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v, err := hex.DecodeString(string(args[0].String()))
	if err != nil {
		return err
	}
	return util.Map(v, func(_ int, c uint8) int { return int(c) })
}
