package wasm_playground

import (
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var EncodingBase32Page = templates.MustRenderDocumentTemplate(templates.DocumentTemplateInput{
	Title: "Base32 Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/base32/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Base32 Encoding",
		Menu:  Menu("Encoding", "Base32"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "encode",
				Title: "Encode",
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
							Operator: "wasm.encoding.base32.EncodeToString",
						},
					},
				},
			},
			{
				Name:  "decode",
				Title: "Decode",
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
							Operator: "wasm.encoding.base32.DecodeString",
						},
					},
				},
			},
		},
	})),
})

var EncodingBase64Page = templates.MustRenderDocumentTemplate(templates.DocumentTemplateInput{
	Title: "Base64 Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/base64/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Base64 Encoding",
		Menu:  Menu("Encoding", "Base64"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "encode",
				Title: "Encode",
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
							Operator: "wasm.encoding.base64.EncodeToString",
						},
					},
				},
			},
			{
				Name:  "decode",
				Title: "Decode",
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
							Operator: "wasm.encoding.base64.DecodeString",
						},
					},
				},
			},
		},
	})),
})

var EncodingHexPage = templates.MustRenderDocumentTemplate(templates.DocumentTemplateInput{
	Title: "Hex Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Hex Encoding",
		Menu:  Menu("Encoding", "Hex"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "encode",
				Title: "Encode",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "String"},
							{Name: "from-bytes", Title: "From Bytes", Operator: "uint8ArrayFromSpaceSeparatedString"},
						},
					},
				},
				Result: templates.WasmPlaygroundTabResult{
					Operators: []templates.WasmPlaygroundTabOperator{
						{
							Name:     "as-text",
							Title:    "As Text",
							Operator: "wasm.encoding.hex.EncodeToString",
						},
					},
				},
			},
			{
				Name:  "decode",
				Title: "Decode",
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
							Operator: "wasm.encoding.hex.DecodeString",
						},
						{
							Name:     "as-bytes",
							Title:    "As Bytes",
							Operator: "wasm.encoding.hex.DecodeStringAsBytes",
						},
					},
				},
			},
		},
	})),
})
