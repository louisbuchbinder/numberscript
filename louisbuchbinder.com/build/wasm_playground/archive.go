package wasm_playground

import (
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var ArchiveChecksumDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Archive Checksum",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/static/js/GoFile.js"})),            // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/md5/pkg/wasm.js"})),    // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha1/pkg/wasm.js"})),   // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha256/pkg/wasm.js"})), // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/crypto/sha512/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Archive Checksum",
		Menu:  Menu("Archive", "Checksum"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "md5",
				Title: "MD5",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_File,
						Name:  "file",
						Title: "File",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-file", Title: "From File", Operator: "firstGoFile"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.md5.AsyncChecksum",
							},
						},
					},
				},
			},
			{
				Name:  "sha1",
				Title: "SHA1",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_File,
						Name:  "file",
						Title: "File",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-file", Title: "From File", Operator: "firstGoFile"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha1.AsyncChecksum",
							},
						},
					},
				},
			},
			{
				Name:  "sha256",
				Title: "SHA256",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_File,
						Name:  "file",
						Title: "File",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-file", Title: "From File", Operator: "firstGoFile"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha256.AsyncChecksum",
							},
						},
					},
				},
			},
			{
				Name:  "sha512",
				Title: "SHA512",
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_File,
						Name:  "file",
						Title: "File",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-file", Title: "From File", Operator: "firstGoFile"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.crypto.sha512.AsyncChecksum",
							},
						},
					},
				},
			},
		},
	})),
}
