package wasmrand

import (
	"crypto/rand"

	"github.com/louisbuchbinder/core/wasm"
)

func Text(args []wasm.Value) any {
	return rand.Text()
}
