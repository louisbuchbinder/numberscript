package wasmcrc64

import (
	"encoding/hex"
	"hash/crc64"

	"github.com/louisbuchbinder/core/wasm"
)

func ChecksumISO(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	data, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	return int(crc64.Checksum(data, crc64.MakeTable(crc64.ISO)))
}

func ChecksumECMA(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	data, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	return int(crc64.Checksum(data, crc64.MakeTable(crc64.ECMA)))
}
