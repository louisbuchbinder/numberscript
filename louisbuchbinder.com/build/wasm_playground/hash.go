package wasm_playground

import (
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var HashAdler32DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Adler32 Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/hash/adler32/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Adler32 Hash",
		Menu:  Menu("Hash", "Adler32"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "checksum",
				Title: "Checksum",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.hash.adler32.Checksum",
							},
						},
					},
				},
			},
		},
	})),
}

var HashCRC32DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "CRC32 Checksum",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/hash/crc32/pkg/wasm.js"})),   // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "CRC32 Checksum",
		Menu:  Menu("Hash", "CRC32"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "ieee",
				Title: "IEEE",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.hash.crc32.ChecksumIEEE",
							},
						},
					},
				},
			},
			{
				Name:  "castangnoli",
				Title: "Castagnoli",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.hash.crc32.ChecksumCastagnoli",
							},
						},
					},
				},
			},
			{
				Name:  "koopman",
				Title: "Koopman",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.hash.crc32.ChecksumKoopman",
							},
						},
					},
				},
			},
		},
	})),
}

var HashCRC64DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "CRC64 Checksum",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/hash/crc64/pkg/wasm.js"})),   // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "CRC64 Checksum",
		Menu:  Menu("Hash", "CRC64"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "iso",
				Title: "ISO",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.hash.crc64.ChecksumISO",
							},
						},
					},
				},
			},
			{
				Name:  "ecma",
				Title: "ECMA",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "data",
						Title: "Data",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "wasm.encoding.hex.EncodeToString"},
							{Name: "from-hex", Title: "From Hex", Operator: "String"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.hash.crc64.ChecksumECMA",
							},
						},
					},
				},
			},
		},
	})),
}
