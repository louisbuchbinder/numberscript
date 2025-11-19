package wasmsha512

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/louisbuchbinder/core/lib/util"
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

func AsyncChecksum(args []wasm.Value) any {
	resolve, reject := args[0], args[1]
	if len(args) < 3 {
		return resolve.Invoke()
	}
	file, err := args[2].File()
	if err != nil {
		reject.Reject(err)
		return nil
	}
	info, err := file.Stat()
	if err != nil {
		reject.Reject(err)
		return nil
	}
	go func() {
		h := sha512.New()
		n, err := util.CopyFlexBuffer(h, file, info.Size())
		if err != nil {
			reject.Reject(err)
			return
		}
		if n != info.Size() {
			reject.Reject(fmt.Errorf("expected to read %d bytes, but instead got %d", info.Size(), n))
			return
		}
		resolve.Invoke(hex.EncodeToString(h.Sum(nil)))
	}()
	return nil
}
