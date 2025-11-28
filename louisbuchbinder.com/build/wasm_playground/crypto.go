package wasm_playground

import (
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var CryptoAESDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "AES Encryption",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/aes/pkg/wasm.js"})),   // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "AES Encryption",
		Menu:  Menu("Crypto", "AES"),
		Docstring: `
			AES encryption (formerly Rijndael), as defined in U.S.
			Federal Information Processing Standards Publication 197.
			The AES operations in this package are not implemented using constant-time
			algorithms.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "encrypt",
				Title: "Encrypt",
				Docstring: `
					Encrypt using Galois Counter
					Mode, with randomly-generated nonces.

					It generates a random 96-bit nonce, which is prepended to the ciphertext
					by Seal, and is extracted from the ciphertext by Open. The NonceSize of the
					AEAD is zero, while the Overhead is 28 bytes (the combination of nonce size
					and tag size).

					A given key MUST NOT be used to encrypt more than 2^32 messages, to limit
					the risk of a random nonce collision to negligible levels.

					The key argument must be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "AES Ciphertext Result:",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.aes.Encrypt",
							},
						},
					},
				},
			},
			{
				Name:  "encrypt-consistent",
				Title: "Encrypt Consistent",
				Docstring: `
					Encrypt using Galois Counter Mode, with the specified nonce.
					Repeated calls to this function will produce the same results.
					The key argument must be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "nonce",
						Title: "Nonce",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "AES Ciphertext Result:",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.aes.EncryptConsistent",
							},
						},
					},
				},
			},
			{
				Name:  "decrypt",
				Title: "Decrypt",
				Docstring: `
					Decrypt the given ciphertext given the encryption key.
					The key argument must be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "ciphertext",
						Title: "Ciphertext",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "AES Plaintext Result:",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.aes.Decrypt",
							},
						},
					},
				},
			},
			{
				Name:  "decrypt-consistent",
				Title: "Decrypt Consistent",
				Docstring: `
					Decrypt the given ciphertext given the encryption key and nonce.
					The key argument must be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "nonce",
						Title: "Nonce",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "ciphertext",
						Title: "Ciphertext",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "AES Plaintext Result:",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.aes.DecryptConsistent",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoECDHDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Elliptic Curve Diffie-Hellman",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/ecdh/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Elliptic Curve Diffie-Hellman",
		Menu:  Menu("Crypto", "ECDH"),
		Docstring: `
			Elliptic Curve Diffie-Hellman over NIST curves and
			Curve25519.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "p256",
				Title: "P256",
				Docstring: `
					Generate a public/private key pair using
					NIST P-256 (FIPS 186-3, section D.2.3), also known as secp256r1 or prime256v1.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.P256",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.P256",
							},
						},
					},
				},
			},
			{
				Name:  "p384",
				Title: "P384",
				Docstring: `
					Generate a public/private key pair using
					NIST P-384 (FIPS 186-3, section D.2.4), also known as secp384r1.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.P384",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.P384",
							},
						},
					},
				},
			},
			{
				Name:  "p521",
				Title: "P521",
				Docstring: `
					Generate a public/private key pair using
					NIST P-521 (FIPS 186-3, section D.2.5), also known as secp521r1.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.P521",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.P521",
							},
						},
					},
				},
			},
			{
				Name:  "x25519",
				Title: "X25519",
				Docstring: `
					Generate a public/private key pair using
					the X25519 function over Curve25519 (RFC 7748, Section 5).
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.X25519",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdh.X25519",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoECDSADocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Elliptic Curve Diffie-Hellman",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/ecdsa/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Elliptic Curve Digital Signature Algorithm",
		Menu:  Menu("Crypto", "ECDSA"),
		Docstring: `
			Elliptic Curve Digital Signature Algorithm,
			as defined in <a target="_blank" href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS 186-5</a>.

			Signatures generated by this package are not deterministic, but entropy is mixed
			with the private key and the message, achieving the same level of security in
			case of randomness source failure.

			Operations involving private keys are implemented using constant-time
			algorithms, as long as an elliptic.Curve returned by elliptic.P224,
			elliptic.P256, elliptic.P384, or elliptic.P521 is used.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "p224",
				Title: "P224",
				Docstring: `
					Generate public/private key pairs using
					NIST P-224 (FIPS 186-3, section D.2.2), also known as secp224r1.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P224",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P224",
							},
						},
					},
				},
			},
			{
				Name:  "p256",
				Title: "P256",
				Docstring: `
					Generate public/private key pairs using
					NIST P-256 (FIPS 186-3, section D.2.3), also known as secp256r1 or prime256v1.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P256",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P256",
							},
						},
					},
				},
			},
			{
				Name:  "p384",
				Title: "P384",
				Docstring: `
					Generate public/private key pairs using
					NIST P-384 (FIPS 186-3, section D.2.4), also known as secp384r1.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P384",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P384",
							},
						},
					},
				},
			},
			{
				Name:  "p521",
				Title: "P521",
				Docstring: `
					Generate public/private key pairs using
					NIST P-521 (FIPS 186-3, section D.2.5), also known as secp521r1.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P521",
							},
						},
					},
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.P521",
							},
						},
					},
				},
			},
			{
				Name:  "signasn1",
				Title: "SignASN1",
				Docstring: `
					Hash and sign the content given the private key and type.
					The type must match the private key type and be one of
					"P224", "P256", "P384", or "P521". A corresponding checksum of
					"SHA224", "SHA256", "SHA384", or "SHA512" will be used respectively.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "type",
						Title: "Type",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "private-key",
						Title: "PrivateKey",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Signature",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ecdsa.SignASN1",
							},
						},
					},
				},
			},
			{
				Name:  "verifyasn1",
				Title: "VerifyASN1",
				Docstring: `
					Verify the signature matches the content hash given the public key and type.
					The type must match the public key type and be one of
					"P224", "P256", "P384", or "P521". A corresponding checksum of
					"SHA224", "SHA256", "SHA384", or "SHA512" will be used respectively.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "type",
						Title: "Type",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "public-key",
						Title: "PublicKey",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "signature",
						Title: "Signature",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Is Valid",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.ecdsa.VerifyASN1",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoED25519DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "ED25519",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/ed25519/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "ED25519",
		Menu:  Menu("Crypto", "ED25519"),
		Docstring: `
			<a target="_blank" href="https://ed25519.cr.yp.to/">Ed25519</a> signature algorithm.
			These functions are also compatible with the “Ed25519” function defined in
			RFC 8032. However, unlike RFC 8032's formulation, this package's private key
			representation includes a public key suffix to make multiple signing operations
			with the same key more efficient.
			Operations involving private keys are implemented using constant-time
			algorithms.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "generate-key",
				Title:     "GenerateKey",
				Docstring: "Generate a public/private key pair.",
				Args:      nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Public Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ed25519.GenerateKey",
							},
						},
					},
					{
						Title: "Private Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ed25519.GenerateKey",
							},
						},
					},
				},
			},
			{
				Name:      "sign",
				Title:     "Sign",
				Docstring: "Sign the message with privateKey and returns a signature.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "private-key",
						Title: "PrivateKey",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Signature",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.ed25519.Sign",
							},
						},
					},
				},
			},
			{
				Name:  "verify",
				Title: "Verify",
				Docstring: `
					Verify reports whether sig is a valid signature of message by publicKey.
					The inputs are not considered confidential, and may leak through timing side
    				channels, or if an attacker has control of part of the inputs.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "public-key",
						Title: "PublicKey",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "signature",
						Title: "Signature",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Title: "Is Valid",
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.ed25519.Verify",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoHMACDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "HMAC",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/hmac/pkg/wasm.js"})),  // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Hash Based Message Authentication Code (HMAC)",
		Menu:  Menu("Crypto", "HMAC"),
		Docstring: `
			Keyed-Hash Message Authentication Code (HMAC) as
			defined in U.S. Federal Information Processing Standards Publication 198.
			An HMAC is a cryptographic hash that uses a key to sign a message. The receiver
			verifies the hash by recomputing it using the same key.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "hmac-md5",
				Title: "HMAC_MD5",
				Docstring: `
					Generate a MD5 HMAC for the content given the key.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.hmac.HMAC_MD5",
							},
						},
					},
				},
			},
			{
				Name:  "hmac-sha1",
				Title: "HMAC_SHA1",
				Docstring: `
					Generate a SHA1 HMAC for the content given the key.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.hmac.HMAC_SHA1",
							},
						},
					},
				},
			},
			{
				Name:  "hmac-sha256",
				Title: "HMAC_SHA256",
				Docstring: `
					Generate a SHA256 HMAC for the content given the key.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.hmac.HMAC_SHA256",
							},
						},
					},
				},
			},
			{
				Name:  "hmac-sha512",
				Title: "HMAC_SHA512",
				Docstring: `
					Generate a SHA512 HMAC for the content given the key.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "key",
						Title: "Key",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "content",
						Title: "Content",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.hmac.HMAC_SHA512",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoMD5DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "MD5 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/md5/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "MD5 Hash",
		Menu:  Menu("Crypto", "MD5"),
		Docstring: `
			MD5 hash algorithm as defined in RFC 1321.
			MD5 is cryptographically broken and should not be used for secure applications.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "hash",
				Title:     "Hash",
				Docstring: "MD5 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.md5.Sum",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoPBKDF2DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "PBKDF2",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/pbkdf2/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Password Based Key Derrivation Function (PBKDF2)",
		Menu:  Menu("Crypto", "PBKDF2"),
		Docstring: `
			Password based key derivation function PBKDF2 as defined in RFC
			8018 (PKCS #5 v2.1).
			A key derivation function is useful when encrypting data based on a password
			or any other not-fully-random data. It uses a pseudorandom function to derive a
			secure encryption key based on the password.
			<br>
			The key is
			derived based on the method described as PBKDF2 with the HMAC variant using
			the supplied hash function.
			Remember to get a good random salt. At least 8 bytes is recommended by the
			RFC.

			Using a higher iteration count will increase the cost of an exhaustive
			search but will also make derivation proportionally slower.

			keyLength must be a positive integer
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "pbkdf2-sha1",
				Title:     "PBKDF2_SHA1",
				Docstring: "Generate a key using the SHA1 hashing function.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "password",
						Title: "Password",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "salt",
						Title: "Salt",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "iterations",
						Title: "Iterations",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 1,
								Max: 2000000,
							},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "key-length",
						Title: "Key Length",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 20,
								Max: 128,
							},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.pbkdf2.PBKDF2_SHA1",
							},
						},
					},
				},
			},
			{
				Name:      "pbkdf2-sha256",
				Title:     "PBKDF2_SHA256",
				Docstring: "Generate a key using the SHA256 hashing function.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "password",
						Title: "Password",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "salt",
						Title: "Salt",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "iterations",
						Title: "Iterations",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 1,
								Max: 2000000,
							},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "key-length",
						Title: "Key Length",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 32,
								Max: 128,
							},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.pbkdf2.PBKDF2_SHA256",
							},
						},
					},
				},
			},
			{
				Name:      "pbkdf2-sha512",
				Title:     "PBKDF2_SHA512",
				Docstring: "Generate a key using the SHA512 hashing function.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "password",
						Title: "Password",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "salt",
						Title: "Salt",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "iterations",
						Title: "Iterations",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 64,
								Max: 128,
							},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "key-length",
						Title: "Key Length",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 1,
								Max: 2000000,
							},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.pbkdf2.PBKDF2_SHA512",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoRandDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Rand",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/rand/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title:     "Rand",
		Menu:      Menu("Crypto", "Rand"),
		Docstring: "Cryptographically secure random number generator.",
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "int",
				Title: "Int",
				Docstring: `
					Int returns a uniform random value in [0, max).
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "max",
						Title: "Max",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.rand.Int",
							},
						},
					},
				},
			},
			{
				Name:      "prime",
				Title:     "Prime",
				Docstring: "Prime returns a number of the given bit length that is prime with high probability.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "bits",
						Title: "Bits",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 2,
								Max: 4096,
							},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.rand.Prime",
							},
						},
					},
				},
			},
			{
				Name:  "bytes",
				Title: "Bytes",
				Docstring: `
					Generates cryptographically secure random bytes.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "count",
						Title: "Count",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 0,
								Max: 1000000,
							},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-hex",
								Title:    "As Hex",
								Operator: "wasm.crypto.rand.Bytes",
							},
						},
					},
				},
			},
			{
				Name:  "text",
				Title: "Text",
				Docstring: `
					Text returns a cryptographically random string using the standard RFC
					4648 base32 alphabet for use when a secret string, token, password, or
					other text is needed. The result contains at least 128 bits of randomness,
					enough to prevent brute force guessing attacks and to make the likelihood
					of collisions vanishingly small. A future version may return longer texts as
					needed to maintain those properties.
				`,
				Args: nil,
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.rand.Text",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoSHA1DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA1 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha1/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "SHA1 Hash",
		Menu:  Menu("Crypto", "SHA1"),
		Docstring: `
			SHA-1 hash algorithm as defined in RFC 3174.
			SHA-1 is cryptographically broken and should not be used for secure
			applications.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "hash",
				Title:     "Hash",
				Docstring: "SHA-1 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha1.Sum",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoSHA3DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA3 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha3/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "SHA3 Hash",
		Menu:  Menu("Crypto", "SHA3"),
		Docstring: `
			SHA-3 hash algorithms and the SHAKE extendable
			output functions defined in FIPS 202.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "sum224-hash",
				Title:     "Sum224 Hash",
				Docstring: "SHA3-224 hash of data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha3.Sum224",
							},
						},
					},
				},
			},
			{
				Name:      "sum256-hash",
				Title:     "Sum256 Hash",
				Docstring: "SHA3-256 hash of data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha3.Sum256",
							},
						},
					},
				},
			},
			{
				Name:      "sum384-hash",
				Title:     "Sum384 Hash",
				Docstring: "SHA3-384 hash of data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha3.Sum384",
							},
						},
					},
				},
			},
			{
				Name:      "sum512-hash",
				Title:     "Sum512 Hash",
				Docstring: "SHA3-512 hash of data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha3.Sum512",
							},
						},
					},
				},
			},
			{
				Name:  "sum-shake-128-hash",
				Title: "SumSHAKE128 Hash",
				Docstring: `
					Applies the SHAKE128 extendable output function to data and
					returns an output of the given length in bytes.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "length",
						Title: "Length",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 0,
								Max: 1024,
							},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha3.SumSHAKE128",
							},
						},
					},
				},
			},
			{
				Name:  "sum-shake-256-hash",
				Title: "SumSHAKE256 Hash",
				Docstring: `
					SumSHAKE256 applies the SHAKE256 extendable output function to data and
					returns an output of the given length in bytes.
				`,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Number,
						Name:  "length",
						Title: "Length",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							NumberOptions: &templates.WasmPlaygroundTabArgOptions_Number{
								Min: 0,
								Max: 1024,
							},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha3.SumSHAKE256",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoSHA256DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA256 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha256/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title:     "SHA256 Hash",
		Menu:      Menu("Crypto", "SHA256"),
		Docstring: "SHA224 and SHA256 hash algorithms as defined in FIPS 180-4.",
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "sum224-hash",
				Title:     "Sum224 Hash",
				Docstring: "SHA224 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha256.Sum224",
							},
						},
					},
				},
			},
			{
				Name:      "sum256-hash",
				Title:     "Sum256 Hash",
				Docstring: "SHA256 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha256.Sum256",
							},
						},
					},
				},
			},
		},
	})),
}

var CryptoSHA512DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA512 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha512/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "SHA512 Hash",
		Menu:  Menu("Crypto", "SHA512"),
		Docstring: `
			SHA-384, SHA-512, SHA-512/224, and SHA-512/256 hash algorithms as defined in FIPS 180-4.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "sum512_224-hash",
				Title:     "Sum512_224 Hash",
				Docstring: "Sum512/224 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha512.Sum512_224",
							},
						},
					},
				},
			},
			{
				Name:      "sum512_256-hash",
				Title:     "Sum512_256 Hash",
				Docstring: "Sum512/256 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha512.Sum512_256",
							},
						},
					},
				},
			},
			{
				Name:      "sum384-hash",
				Title:     "Sum384 Hash",
				Docstring: "SHA384 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha512.Sum384",
							},
						},
					},
				},
			},
			{
				Name:      "sum512-hash",
				Title:     "Sum512 Hash",
				Docstring: "SHA512 checksum of the data.",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha512.Sum512",
							},
						},
					},
				},
			},
		},
	})),
}
