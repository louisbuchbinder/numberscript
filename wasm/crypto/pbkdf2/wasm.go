package wasmpbkdf2

import (
	"crypto/pbkdf2"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"

	"github.com/louisbuchbinder/core/wasm"
)

func fn(h func() hash.Hash, args []wasm.Value) any {
	if len(args) < 4 {
		return nil
	}
	password := args[0].String()
	saltBytes, err := hex.DecodeString(args[1].String())
	if err != nil {
		return err
	}
	if len(saltBytes) < 8 {
		return fmt.Errorf("expected salt legth to be at leas 8 bytes, but instead got: %d", len(saltBytes))
	}
	iterations := args[2].Int()
	if iterations < 1 || iterations > 2e6 {
		return fmt.Errorf("expected iterations to be between 1 and 2000000, but instead got: %d", iterations)
	}
	keyLength := args[3].Int()
	if keyLength < h().Size() || keyLength > 128 {
		return fmt.Errorf("expected keyLength between %d and 128, but instead got: %d", h().Size(), keyLength)
	}
	key, err := pbkdf2.Key(h, password, saltBytes, iterations, keyLength)
	if err != nil {
		return err
	}
	return hex.EncodeToString(key)
}

func PBKDF2_SHA1(args []wasm.Value) any {
	return fn(sha1.New, args)
}

func PBKDF2_SHA256(args []wasm.Value) any {
	return fn(sha256.New, args)
}

func PBKDF2_SHA512(args []wasm.Value) any {
	return fn(sha512.New, args)
}
