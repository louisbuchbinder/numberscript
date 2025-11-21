package wasmsha256

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/wasm"
)

func Sum224(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha256.Sum224([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func Sum256(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha256.Sum256([]byte(args[0].String()))
	return hex.EncodeToString(v[:])
}

func AsyncChecksum(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
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
		h := sha256.New()
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
