package main

import (
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var MainPage = templates.MustRenderDocumentTemplate(templates.DocumentTemplateInput{
	Title:   "Home",
	Scripts: "",
	Main:    template.HTML(templates.MustRenderHomeTemplate(templates.HomeTemplateInput{})),
})

var EncodingHexPage = templates.MustRenderDocumentTemplate(templates.DocumentTemplateInput{
	Title: "Hex Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "Hex Encoding",
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

func write(f string, content []byte) error {
	if err := os.MkdirAll(path.Dir(f), 0o700); err != nil {
		return err
	}
	if err := os.WriteFile(f, content, 0o644); err != nil {
		return err
	}
	return nil
}

func main() {
	util.Must0(write("index.html", MainPage))
	util.Must0(write("encoding/hex/index.html", EncodingHexPage))
}
