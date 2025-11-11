package wasmecdh

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/hex"

	"github.com/louisbuchbinder/core/wasm"
)

func key(c ecdh.Curve) (*ecdh.PrivateKey, error) {
	privateKey, err := c.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func P256(args []wasm.Value) any {
	privateKey, err := key(ecdh.P256())
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKey.Bytes()), hex.EncodeToString(privateKey.PublicKey().Bytes())}
}

func P384(args []wasm.Value) any {
	privateKey, err := key(ecdh.P384())
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKey.Bytes()), hex.EncodeToString(privateKey.PublicKey().Bytes())}
}

func P521(args []wasm.Value) any {
	privateKey, err := key(ecdh.P521())
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKey.Bytes()), hex.EncodeToString(privateKey.PublicKey().Bytes())}
}

func X25519(args []wasm.Value) any {
	privateKey, err := key(ecdh.X25519())
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKey.Bytes()), hex.EncodeToString(privateKey.PublicKey().Bytes())}
}
