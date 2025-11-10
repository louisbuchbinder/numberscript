package wasmaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/louisbuchbinder/core/wasm"
)

func newCipherBlock(keyHex string) (cipher.Block, error) {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, err
	}
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("The key argument must be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256. Instead got: %d", len(key))
	}
	return aes.NewCipher(key)
}

func encrypt(keyHex, plaintext string) ([]byte, []byte, error) {
	block, err := newCipherBlock(keyHex)
	if err != nil {
		return nil, nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}
	return aead.Seal(nil, nonce, []byte(plaintext), nil), nonce, nil
}

func Encrypt(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	key := args[0].String()
	plaintext := args[1].String()
	ciphertext, nonce, err := encrypt(key, plaintext)
	if err != nil {
		return err
	} else {
		return []any{hex.EncodeToString(ciphertext), hex.EncodeToString(nonce)}
	}
}
