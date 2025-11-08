package wasm_playground

import (
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var CryptoMD5DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "MD5 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/md5/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "MD5 Hash",
		Menu:  Menu("Crypto", "MD5"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "hash",
				Title: "Hash",
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
				Result: templates.WasmPlaygroundTabResult{
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
	})),
}

var CryptoRandDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Rand",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/rand/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Rand",
		Menu:  Menu("Crypto", "Rand"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "text",
				Title: "Text",
				Args:  nil,
				Result: templates.WasmPlaygroundTabResult{
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
	})),
}

var CryptoSHA1DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA1 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha1/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "SHA1 Hash",
		Menu:  Menu("Crypto", "SHA1"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "hash",
				Title: "Hash",
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
				Result: templates.WasmPlaygroundTabResult{
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
	})),
}

var CryptoSHA3DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA3 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha3/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "SHA3 Hash",
		Menu:  Menu("Crypto", "SHA3"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "sum224-hash",
				Title: "Sum224 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha3.Sum224",
						},
					},
				},
			},
			{
				Name:  "sum256-hash",
				Title: "Sum256 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha3.Sum256",
						},
					},
				},
			},
			{
				Name:  "sum384-hash",
				Title: "Sum384 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha3.Sum384",
						},
					},
				},
			},
			{
				Name:  "sum512-hash",
				Title: "Sum512 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha3.Sum512",
						},
					},
				},
			},
			{
				Name:  "sum-shake-128-hash",
				Title: "SumSHAKE128 Hash",
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
						Type:  templates.WasmPlaygroundTabValType_Int,
						Name:  "length",
						Title: "Length",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							IntOptions: templates.WasmPlaygroundTabArgOptions_Int{
								Min: 0,
								Max: 1024,
							},
						},
					},
				},
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha3.SumSHAKE128",
						},
					},
				},
			},
			{
				Name:  "sum-shake-256-hash",
				Title: "SumSHAKE256 Hash",
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
						Type:  templates.WasmPlaygroundTabValType_Int,
						Name:  "length",
						Title: "Length",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-integer", Title: "From Integer", Operator: "safeUInt"},
						},
						Options: templates.WasmPlaygroundTabArgOptions{
							IntOptions: templates.WasmPlaygroundTabArgOptions_Int{
								Min: 0,
								Max: 1024,
							},
						},
					},
				},
				Result: templates.WasmPlaygroundTabResult{
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
	})),
}

var CryptoSHA256DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA256 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha256/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "SHA256 Hash",
		Menu:  Menu("Crypto", "SHA256"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "sum224-hash",
				Title: "Sum224 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha256.Sum224",
						},
					},
				},
			},
			{
				Name:  "sum256-hash",
				Title: "Sum256 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
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
	})),
}

var CryptoSHA512DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "SHA512 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha512/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "SHA512 Hash",
		Menu:  Menu("Crypto", "SHA512"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "sum512_224-hash",
				Title: "Sum512_224 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha512.Sum512_224",
						},
					},
				},
			},
			{
				Name:  "sum512_256-hash",
				Title: "Sum512_256 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha512.Sum512_256",
						},
					},
				},
			},
			{
				Name:  "sum384-hash",
				Title: "Sum384 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.crypto.sha512.Sum384",
						},
					},
				},
			},
			{
				Name:  "sum512-hash",
				Title: "Sum512 Hash",
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
				Result: templates.WasmPlaygroundTabResult{
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
	})),
}
