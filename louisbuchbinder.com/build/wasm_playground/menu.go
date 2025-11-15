package wasm_playground

import (
	"html/template"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var WasmPlaygroundMenuItemContainers = []templates.WasmPlaygroundMenuItemContainer{
	{
		Title: "Archive",
		Items: []templates.WasmPlaygroundMenuItem{
			{
				Url:   "/archive/checksum",
				Title: "Checksum",
			},
		},
	},
	{
		Title: "Crypto",
		Items: []templates.WasmPlaygroundMenuItem{
			{
				Url:   "/crypto/aes",
				Title: "AES",
			},
			{
				Url:   "/crypto/ecdh",
				Title: "ECDH",
			},
			{
				Url:   "/crypto/ecdsa",
				Title: "ECDSA",
			},
			{
				Url:   "/crypto/ed25519",
				Title: "ED25519",
			},
			{
				Url:   "/crypto/hmac",
				Title: "HMAC",
			},
			{
				Url:   "/crypto/md5",
				Title: "MD5",
			},
			{
				Url:   "/crypto/pbkdf2",
				Title: "PBKDF2",
			},
			{
				Url:   "/crypto/rand",
				Title: "Rand",
			},
			{
				Url:   "/crypto/sha1",
				Title: "SHA1",
			},
			{
				Url:   "/crypto/sha3",
				Title: "SHA3",
			},
			{
				Url:   "/crypto/sha256",
				Title: "SHA256",
			},
			{
				Url:   "/crypto/sha512",
				Title: "SHA512",
			},
		},
	},
	{
		Title: "Encoding",
		Items: []templates.WasmPlaygroundMenuItem{
			{
				Url:   "/encoding/base32",
				Title: "Base32",
			},
			{
				Url:   "/encoding/base64",
				Title: "Base64",
			},
			{
				Url:   "/encoding/hex",
				Title: "Hex",
			},
			{
				Url:   "/encoding/html",
				Title: "HTML",
			},
			{
				Url:   "/encoding/uri",
				Title: "URI",
			},
		},
	},
	{
		Title: "Hash",
		Items: []templates.WasmPlaygroundMenuItem{
			{
				Url:   "/hash/adler32",
				Title: "Adler32",
			},
			{
				Url:   "/hash/crc32",
				Title: "CRC32",
			},
			{
				Url:   "/hash/crc64",
				Title: "CRC64",
			},
			{
				Url:   "/hash/fnv",
				Title: "FNV",
			},
		},
	},
	{
		Title: "Math",
		Items: []templates.WasmPlaygroundMenuItem{
			{
				Url:   "/math",
				Title: "General",
			},
		},
	},
}

var WasmPlaygroundMainMenuItem = util.Flatten(util.Map(WasmPlaygroundMenuItemContainers, func(_ int, container templates.WasmPlaygroundMenuItemContainer) []templates.WasmPlaygroundMainMenuItem {
	return util.Map(container.Items, func(_ int, item templates.WasmPlaygroundMenuItem) templates.WasmPlaygroundMainMenuItem {
		return templates.WasmPlaygroundMainMenuItem{
			Key:         container.Title + "/" + item.Title,
			Url:         item.Url,
			Description: "",
		}
	})
}))

func Menu(ActiveContainer, ActiveItem string) template.HTML {
	return template.HTML(templates.MustRenderWasmPlaygroundMenuTemplate(templates.WasmPlaygroundMenuTemplateInput{
		ActiveContainer: ActiveContainer,
		ActiveItem:      ActiveItem,
		Containers:      WasmPlaygroundMenuItemContainers,
	}))
}
