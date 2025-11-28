package wasm_playground

import (
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

var EncodingBase32DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Base32 Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/base32/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title:     "Base32 Encoding",
		Menu:      Menu("Encoding", "Base32"),
		Docstring: `Base32 encoding as specified by RFC 4648.`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "encode",
				Title:     "Encode",
				Docstring: "Returns the base32 encoding of src",
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
								Operator: "wasm.encoding.base32.EncodeToString",
							},
						},
					},
				},
			},
			{
				Name:  "decode",
				Title: "Decode",
				Docstring: `
					Returns the bytes represented by the base32 string.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
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
		},
	})),
}

var EncodingBase64DocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Base64 Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/base64/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title:     "Base64 Encoding",
		Menu:      Menu("Encoding", "Base64"),
		Docstring: "Base64 encoding as specified by RFC 4648.",
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "encode",
				Title:     "Encode",
				Docstring: "Returns the base64 encoding of src.",
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
								Operator: "wasm.encoding.base64.EncodeToString",
							},
						},
					},
				},
			},
			{
				Name:      "decode",
				Title:     "Decode",
				Docstring: "Returns the bytes represented by the base64 string.",
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
								Operator: "wasm.encoding.base64.DecodeString",
							},
						},
					},
				},
			},
		},
	})),
}

var EncodingHexDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Hex Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title:     "Hex Encoding",
		Menu:      Menu("Encoding", "Hex"),
		Docstring: "Hexadecimal encoding and decoding.",
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "encode",
				Title:     "Encode",
				Docstring: "Returns the hexadecimal encoding of src.",
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
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.encoding.hex.EncodeToString",
							},
						},
					},
				},
			},
			{
				Name:  "decode",
				Title: "Decode",
				Docstring: `
				Returns the bytes represented by the hexadecimal string s.
				Expects that src contains only hexadecimal characters and that
				src has even length.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
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
		},
	})),
}

var EncodingHTMLDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "HTML Encoding",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/html/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title:     "HTML Escape",
		Menu:      Menu("Encoding", "HTML"),
		Docstring: "Functions for escaping and unescaping HTML text.",
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "escape",
				Title: "Escape",
				Docstring: `
					EscapeString escapes special characters like "<" to become "&lt;".
					It escapes only five such characters: <, >, &, ' and ".
					UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
					always true.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.encoding.html.EscapeString",
							},
						},
					},
				},
			},
			{
				Name:  "unescape",
				Title: "Unescape",
				Docstring: `
					UnescapeString unescapes entities like "&lt;" to become "<".
					It unescapes a larger range of entities than EscapeString escapes.
					For example, "&aacute;" unescapes to "á", as does "&#225;" and "&#xE1;".
					UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
					always true.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "wasm.encoding.html.UnescapeString",
							},
						},
					},
				},
			},
		},
	})),
}

var EncodingURIDocumentTemplateInput = templates.DocumentTemplateInput{
	Title:   "URI Encoding",
	Scripts: template.HTML(strings.Join([]string{}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title:     "URI Encoding",
		Menu:      Menu("Encoding", "URI"),
		Docstring: `Global JavaScript uri encoding functions.`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "encode-uri",
				Title: "Encode URI",
				Docstring: `
					The <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI">encodeURI()</a>
					function encodes a <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Glossary/URI">URI</a> by replacing each instance of
					certain characters by one, two, three, or four escape sequences representing the <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Glossary/UTF-8">UTF-8</a>
					encoding of the character (will only be four escape sequences for characters composed of two surrogate characters). Compared to
					<a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent">encodeURIComponent()</a>,
					this function encodes fewer characters, preserving those that are part of the URI syntax.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "encodeURI",
							},
						},
					},
				},
			},
			{
				Name:  "decode-uri",
				Title: "Decode URI",
				Docstring: `
					The <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/decodeURI">decodeURI()</a> function decodes a Uniform Resource Identifier (URI) previously created by
					<a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI">encodeURI()</a>
					or a similar routine.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "decodeURI",
							},
						},
					},
				},
			},
			{
				Name:  "encode-uri-component",
				Title: "Encode URI Component",
				Docstring: `
					The <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent">encodeURIComponent()</a>
					function encodes a <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Glossary/URI">URI</a>
					by replacing each instance of certain characters by one, two, three, or four escape sequences representing the <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Glossary/UTF-8">UTF-8</a>
					encoding of the character (will only be four escape sequences for characters composed of two surrogate characters). Compared to
					<a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI">encodeURI()</a>,
					this function encodes more characters, including those that are part of the URI syntax.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "encodeURIComponent",
							},
						},
					},
				},
			},
			{
				Name:  "decode-uri-component",
				Title: "Decode URI Component",
				Docstring: `
					The <a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/decodeURIComponent">decodeURIComponent()</a>
					function decodes a Uniform Resource Identifier (URI) component previously created by
					<a target="_blank" href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent">encodeURIComponent()</a>
					or by a similar routine.
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
				},
				Results: []templates.WasmPlaygroundTabResult{
					{
						Operators: []templates.WasmPlaygroundTabOperator{
							{
								Name:     "as-text",
								Title:    "As Text",
								Operator: "decodeURIComponent",
							},
						},
					},
				},
			},
		},
	})),
}
