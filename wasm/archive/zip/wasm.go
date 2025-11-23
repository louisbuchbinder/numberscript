package wasmzip

import (
	"archive/zip"
	"encoding/json"

	"github.com/louisbuchbinder/core/wasm"
)

type ZipArchiveData struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

func AsyncZip(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	resolve, reject := args[0], args[1]
	if len(args) < 4 {
		return resolve.Invoke()
	}
	fw, err := args[2].FileWriter()
	if err != nil {
		reject.Reject(err)
		return nil
	}
	fs, err := args[3].FS()
	if err != nil {
		reject.Reject(err)
		return nil
	}
	go func() {
		writer := zip.NewWriter(fw)
		if err := writer.AddFS(fs); err != nil {
			reject.Reject(err)
			return
		}
		if err := writer.Close(); err != nil {
			reject.Reject(err)
			return
		}
		if err := fw.Close(); err != nil {
			reject.Reject(err)
			return
		}
		stat, err := fw.Stat()
		if err != nil {
			reject.Reject(err)
			return
		}
		b, err := json.Marshal(ZipArchiveData{
			Filename: stat.Name(),
			Size:     int(stat.Size()),
		})
		resolve.Invoke(string(b))
	}()
	return nil
}
