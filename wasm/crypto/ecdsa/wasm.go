package wasmecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/louisbuchbinder/core/wasm"
)

func key(c elliptic.Curve) (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func P224(args []wasm.Value) any {
	privateKey, err := key(elliptic.P224())
	if err != nil {
		return err
	}
	privateKeyBytes, err := privateKey.Bytes()
	if err != nil {
		return err
	}
	publicKeyBytes, err := privateKey.PublicKey.Bytes()
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes)}
}

func P256(args []wasm.Value) any {
	privateKey, err := key(elliptic.P256())
	if err != nil {
		return err
	}
	privateKeyBytes, err := privateKey.Bytes()
	if err != nil {
		return err
	}
	publicKeyBytes, err := privateKey.PublicKey.Bytes()
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes)}
}

func P384(args []wasm.Value) any {
	privateKey, err := key(elliptic.P384())
	if err != nil {
		return err
	}
	privateKeyBytes, err := privateKey.Bytes()
	if err != nil {
		return err
	}
	publicKeyBytes, err := privateKey.PublicKey.Bytes()
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes)}
}

func P521(args []wasm.Value) any {
	privateKey, err := key(elliptic.P521())
	if err != nil {
		return err
	}
	privateKeyBytes, err := privateKey.Bytes()
	if err != nil {
		return err
	}
	publicKeyBytes, err := privateKey.PublicKey.Bytes()
	if err != nil {
		return err
	}
	return []any{hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes)}
}

func SignASN1(args []wasm.Value) any {
	if len(args) < 3 {
		return nil
	}
	typ := args[0].String()
	privateKeyBytes, err := hex.DecodeString(args[1].String())
	if err != nil {
		return err
	}
	content := []byte(args[2].String())
	var (
		privateKey *ecdsa.PrivateKey
		hsh        []byte
	)
	switch typ {
	case "P224":
		k, err := ecdsa.ParseRawPrivateKey(elliptic.P224(), privateKeyBytes)
		if err != nil {
			return err
		}
		privateKey = k
		h := sha256.Sum224(content)
		hsh = h[:]
	case "P256":
		k, err := ecdsa.ParseRawPrivateKey(elliptic.P256(), privateKeyBytes)
		if err != nil {
			return err
		}
		privateKey = k
		h := sha256.Sum256(content)
		hsh = h[:]
	case "P384":
		k, err := ecdsa.ParseRawPrivateKey(elliptic.P384(), privateKeyBytes)
		if err != nil {
			return err
		}
		privateKey = k
		h := sha512.Sum384(content)
		hsh = h[:]
	case "P521":
		k, err := ecdsa.ParseRawPrivateKey(elliptic.P521(), privateKeyBytes)
		if err != nil {
			return err
		}
		privateKey = k
		h := sha512.Sum512(content)
		hsh = h[:]
	default:
		return fmt.Errorf("Curve type must be one of 'P224, P256, P384, P521', but instead got: %s", typ)
	}
	signed, err := ecdsa.SignASN1(rand.Reader, privateKey, hsh)
	if err != nil {
		return err
	}
	return hex.EncodeToString(signed)
}

func VerifyASN1(args []wasm.Value) any {
	if len(args) < 4 {
		return nil
	}
	typ := args[0].String()
	publicKeyBytes, err := hex.DecodeString(args[1].String())
	if err != nil {
		return err
	}
	content := []byte(args[2].String())
	sig, err := hex.DecodeString(args[3].String())
	if err != nil {
		return err
	}
	var (
		publicKey *ecdsa.PublicKey
		hsh       []byte
	)
	switch typ {
	case "P224":
		k, err := ecdsa.ParseUncompressedPublicKey(elliptic.P224(), publicKeyBytes)
		if err != nil {
			return err
		}
		publicKey = k
		h := sha256.Sum224(content)
		hsh = h[:]
	case "P256":
		k, err := ecdsa.ParseUncompressedPublicKey(elliptic.P256(), publicKeyBytes)
		if err != nil {
			return err
		}
		publicKey = k
		h := sha256.Sum256(content)
		hsh = h[:]
	case "P384":
		k, err := ecdsa.ParseUncompressedPublicKey(elliptic.P384(), publicKeyBytes)
		if err != nil {
			return err
		}
		publicKey = k
		h := sha512.Sum384(content)
		hsh = h[:]
	case "P521":
		k, err := ecdsa.ParseUncompressedPublicKey(elliptic.P521(), publicKeyBytes)
		if err != nil {
			return err
		}
		publicKey = k
		h := sha512.Sum512(content)
		hsh = h[:]
	default:
		return fmt.Errorf("Curve type must be one of 'P224, P256, P384, P521', but instead got: %s", typ)
	}
	return ecdsa.VerifyASN1(publicKey, hsh, sig)
}
