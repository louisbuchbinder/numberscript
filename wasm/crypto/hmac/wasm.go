package wasmhmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"github.com/louisbuchbinder/core/wasm"
)

func fn(h func() hash.Hash, args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	keyBytes, err := hex.DecodeString(args[0].String())
	if err != nil {
		return err
	}
	message := []byte(args[1].String())
	mac := hmac.New(h, keyBytes)
	mac.Write(message)
	return hex.EncodeToString(mac.Sum(nil))
}

func HMAC_MD5(args []wasm.Value) any {
	return fn(md5.New, args)
}

func HMAC_SHA1(args []wasm.Value) any {
	return fn(sha1.New, args)
}

func HMAC_SHA256(args []wasm.Value) any {
	return fn(sha256.New, args)
}

func HMAC_SHA512(args []wasm.Value) any {
	return fn(sha512.New, args)
}
