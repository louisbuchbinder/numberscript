package wasmrand

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/louisbuchbinder/core/wasm"
)

func Int(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	max := new(big.Int)
	if err := max.UnmarshalText([]byte(args[0].String())); err != nil {
		return err
	}
	if max.Cmp(big.NewInt(0)) < 1 {
		return fmt.Errorf("max must be greater than zero")
	}
	if val, err := rand.Int(rand.Reader, max); err != nil {
		return err
	} else {
		return val.String()
	}
}

func Prime(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	bits := args[0].Int()
	if bits < 2 {
		return fmt.Errorf("bits must be greater than or equal to 2")
	}
	if val, err := rand.Prime(rand.Reader, bits); err != nil {
		return err
	} else {
		return val.String()
	}
}

func Bytes(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	b := make([]byte, args[0].Int())
	if n, err := rand.Read(b); err != nil {
		return err
	} else if len(b) != n {
		return fmt.Errorf("failed to read %d random bytes", len(b))
	} else {
		return hex.EncodeToString(b)
	}
}

func Text(args []wasm.Value) any {
	return rand.Text()
}
