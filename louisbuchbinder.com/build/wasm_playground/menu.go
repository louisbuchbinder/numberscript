package wasm_playground

import (
	"html/template"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var WasmPlaygroundMenuItemContainers = []templates.WasmPlaygroundMenuItemContainer{
	{
		Title: "Crypto",
		Items: []templates.WasmPlaygroundMenuItem{
			{
				Url:   "/crypto/md5",
				Title: "MD5",
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
}

func Menu(ActiveContainer, ActiveItem string) template.HTML {
	return template.HTML(templates.MustRenderWasmPlaygroundMenuTemplate(templates.WasmPlaygroundMenuTemplateInput{
		ActiveContainer: ActiveContainer,
		ActiveItem:      ActiveItem,
		Containers:      WasmPlaygroundMenuItemContainers,
	}))
}
