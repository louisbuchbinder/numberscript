package wasmzip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/wasm"
)

func AsyncZip(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	resolve, reject := args[0], args[1]
	if len(args) < 3 {
		return resolve.Invoke()
	}
	fs, err := args[2].FS()
	if err != nil {
		reject.Reject(err)
		return nil
	}
	_ = fs
	go func() {
		buf := new(bytes.Buffer) // TODO
		writer := zip.NewWriter(buf)
		if err := writer.AddFS(fs); err != nil {
			reject.Reject(err)
			return
		}
		if err := writer.Close(); err != nil {
			reject.Reject(err)
			return
		}
		resolve.Invoke(strings.Join(util.Map(buf.Bytes(), func(_ int, b uint8) string { return fmt.Sprintf("%d", b) }), " "))
	}()
	return nil
}
