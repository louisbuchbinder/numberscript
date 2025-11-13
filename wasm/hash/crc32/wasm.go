package wasmcrc32

import (
	"encoding/hex"
	"hash/crc32"

	"github.com/louisbuchbinder/core/wasm"
)

func ChecksumIEEE(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	data, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	return int(crc32.Checksum(data, crc32.MakeTable(crc32.IEEE)))
}

func ChecksumCastagnoli(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	data, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	return int(crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli)))
}

func ChecksumKoopman(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	data, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	return int(crc32.Checksum(data, crc32.MakeTable(crc32.Koopman)))
}
