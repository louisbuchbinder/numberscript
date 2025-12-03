package wasm_playground

import (
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/build/load"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var _ = load.Register(func() {
	ArchiveZipDocumentTemplateInput.Scripts = template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.WASM_GO_JS})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/js/GoFile.js")})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/js/OpfsFile.js")})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/js/GoFS.js")})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/wasm/archive/zip/pkg/sha256.wasm.js")})),
	}, "\n"))
})

var ArchiveZipDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Zip Archive",
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Zip Archive",
		Menu:  Menu("Archive", "Zip"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:              "zip",
				Title:             "Zip",
				HasGenerateButton: true,
				Args: []templates.WasmPlaygroundTabArg{
					{
						Type:  templates.WasmPlaygroundTabValType_Text,
						Name:  "output-filename",
						Title: "Output Filename",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-text", Title: "From Text", Operator: "createOpfsFile"},
						},
					},
					{
						Type:  templates.WasmPlaygroundTabValType_Files,
						Name:  "files",
						Title: "Files",
						Operators: []templates.WasmPlaygroundTabOperator{
							{Name: "from-files", Title: "From Files", Operator: "newGoFS"},
						},
					},
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Type: templates.WasmPlaygroundTabValType_Download,
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-file",
								Title:    "As File",
								Operator: "wasm.archive.zip.AsyncZip",
							},
						},
					},
				},
			},
		},
	})),
}

var _ = load.Register(func() {
	ArchiveChecksumDocumentTemplateInput.Scripts = template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.WASM_GO_JS})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/js/GoFile.js")})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/wasm/crypto/md5/pkg/sha256.wasm.js")})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/wasm/crypto/sha1/pkg/sha256.wasm.js")})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/wasm/crypto/sha256/pkg/sha256.wasm.js")})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: load.Sha256Version("/wasm/crypto/sha512/pkg/sha256.wasm.js")})),
	}, "\n"))
})

var ArchiveChecksumDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Archive Checksum",
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Archive Checksum",
		Menu:  Menu("Archive", "Checksum"),
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:              "md5",
				Title:             "MD5",
				HasGenerateButton: true,
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
				Name:              "sha1",
				Title:             "SHA1",
				HasGenerateButton: true,
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
				Name:              "sha256",
				Title:             "SHA256",
				HasGenerateButton: true,
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
				Name:              "sha512",
				Title:             "SHA512",
				HasGenerateButton: true,
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
