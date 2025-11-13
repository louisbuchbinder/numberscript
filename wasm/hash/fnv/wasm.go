package wasmfnv

import (
	"encoding/hex"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/louisbuchbinder/core/wasm"
)

func fn(h func() hash.Hash, args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	data, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	hsh := h()
	n, err := hsh.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return fmt.Errorf("expected to write %d bytes, but instead got %d", len(data), n)
	}
	return hex.EncodeToString(hsh.Sum(nil))
}

func Hash128(args []wasm.Value) any {
	return fn(fnv.New128, args)
}

func Hash128a(args []wasm.Value) any {
	return fn(fnv.New128a, args)
}

func Hash32(args []wasm.Value) any {
	return fn(func() hash.Hash { h := fnv.New32(); return h }, args)
}

func Hash32a(args []wasm.Value) any {
	return fn(func() hash.Hash { h := fnv.New32a(); return h }, args)
}

func Hash64(args []wasm.Value) any {
	return fn(func() hash.Hash { h := fnv.New64(); return h }, args)
}

func Hash64a(args []wasm.Value) any {
	return fn(func() hash.Hash { h := fnv.New64a(); return h }, args)
}
