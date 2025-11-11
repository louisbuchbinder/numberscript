package wasmed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/louisbuchbinder/core/wasm"
)

func GenerateKey(args []wasm.Value) any {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(publicKey), hex.EncodeToString(privateKey)}
}

func Sign(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	privateKeyBytes, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	if len(privateKeyBytes) != ed25519.PrivateKeySize {
		return fmt.Errorf("expected private key length to be %d, but instead got: %d", ed25519.PrivateKeySize, len(privateKeyBytes))
	}
	message := []byte(args[1].String())
	sig := ed25519.Sign(ed25519.PrivateKey(privateKeyBytes), message)
	return hex.EncodeToString(sig)
}

func Verify(args []wasm.Value) any {
	if len(args) < 3 {
		return nil
	}
	publicKeyBytes, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	if len(publicKeyBytes) != ed25519.PublicKeySize {
		return fmt.Errorf("expected public key length to be %d, but instead got: %d", ed25519.PublicKeySize, len(publicKeyBytes))
	}
	message := []byte(args[1].String())
	sigBytes, err := hex.DecodeString(args[2].String())
	if err != nil {
		return err
	}
	return ed25519.Verify(ed25519.PublicKey(publicKeyBytes), message, sigBytes)
}
