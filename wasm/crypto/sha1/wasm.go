package wasmsha1

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/louisbuchbinder/core/wasm"
)

func Sum(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := sha1.Sum([]byte(args[0].String()))
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
		h := sha1.New()
		n, err := io.Copy(h, file)
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
