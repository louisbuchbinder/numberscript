package wasmaes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

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

func Encrypt(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	keyHex := args[0].String()
	plaintext := args[1].String()
	block, err := newCipherBlock(keyHex)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCMWithRandomNonce(block)
	if err != nil {
		return err
	}
	ciphertext := aead.Seal(nil, nil, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	keyHex := args[0].String()
	ciphertextHex := args[1].String()
	block, err := newCipherBlock(keyHex)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCMWithRandomNonce(block)
	if err != nil {
		return err
	}
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return err
	}
	plaintext, err := aead.Open(nil, nil, ciphertext, nil)
	if err != nil {
		return err
	}
	return string(plaintext)
}

func EncryptConsistent(args []wasm.Value) any {
	if len(args) < 3 {
		return nil
	}
	keyHex := args[0].String()
	nonceHex := args[1].String()
	plaintext := args[2].String()
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return err
	}
	block, err := newCipherBlock(keyHex)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	if len(nonce) != aead.NonceSize() {
		return fmt.Errorf("Incorrect nonce size. Expected %d bytes, but instead got %d", aead.NonceSize(), len(nonce))
	}
	ciphertext := aead.Seal(nil, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext)
}

func DecryptConsistent(args []wasm.Value) any {
	if len(args) < 3 {
		return nil
	}
	keyHex := args[0].String()
	nonceHex := args[1].String()
	ciphertextHex := args[2].String()
	block, err := newCipherBlock(keyHex)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return err
	}
	if len(nonce) != aead.NonceSize() {
		return fmt.Errorf("Incorrect nonce size. Expected %d bytes, but instead got %d", aead.NonceSize(), len(nonce))
	}
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return err
	}
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}
	return string(plaintext)
}
