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
		Docstring: `
			Adler-32 is composed of two sums accumulated per byte: s1 is
			the sum of all bytes, s2 is the sum of all s1 values. Both sums
			are done modulo 65521. s1 is initialized to 1, s2 to zero.  The
			Adler-32 checksum is stored as s2*65536 + s1 in most-
			significant-byte first (network) order.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "checksum",
				Title:     "Checksum",
				Docstring: "Checksum returns the Adler-32 checksum of data.",
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
		Docstring: `
			32-bit cyclic redundancy check, or CRC-32,
			checksum. See <a target="_blank" href="https://en.wikipedia.org/wiki/Cyclic_redundancy_check">wiki:Cyclic_redundancy_check</a> for
			information.

			Polynomials are represented in LSB-first form also known as reversed
			representation.

			See
			<a target="_blank" href="https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials">wiki:Mathematics_of_cyclic_redundancy_checks</a>
			for information.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "ieee",
				Title: "IEEE",
				Docstring: `
					IEEE is by far and away the most common CRC-32 polynomial.
					Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...
				`,
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
				Docstring: `
					Castagnoli's polynomial, used in iSCSI.
					Has <a target="_blank" href="https://dx.doi.org/10.1109/26.231911">better</a> error detection characteristics than IEEE.
				`,
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
				Docstring: `
				Koopman's polynomial.
				Also has <a target="_blank" href="https://dx.doi.org/10.1109/DSN.2002.1028931">better</a> error detection characteristics than IEEE.
				`,
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
		Docstring: `
		64-bit cyclic redundancy check, or CRC-64,
		checksum. See <a target="_blank" href="https://en.wikipedia.org/wiki/Cyclic_redundancy_check">wiki:Cyclic_redundancy_check<a> for
		information.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:      "iso",
				Title:     "ISO",
				Docstring: "The ISO polynomial, defined in ISO 3309 and used in HDLC.",
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
				Name:      "ecma",
				Title:     "ECMA",
				Docstring: "The ECMA polynomial, defined in ECMA 182.",
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

var HashFNVDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "FNV Hash",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/encoding/hex/pkg/wasm.js"})), // TODO: use the hash-named file
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/hash/fnv/pkg/wasm.js"})),     // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "FNV Hash",
		Menu:  Menu("Hash", "FNV"),
		Docstring: `
			FNV-1 and FNV-1a, non-cryptographic hash
			functions created by Glenn Fowler, Landon Curt Noll, and Phong Vo. See
			<a target="_blank" href="https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function">wiki:Fowler-Noll-Vo_hash_function</a>.
		`,
		Tabs: []templates.WasmPlaygroundTab{
			{
				Name:  "fnv-128",
				Title: "FNV 128",
				Docstring: `
					128-bit FNV-1 hash.Hash. Its Sum method will lay the
    				value out in big-endian byte order.
				`,
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
								Operator: "wasm.hash.fnv.Hash128",
							},
						},
					},
				},
			},
			{
				Name:  "fnv-128a",
				Title: "FNV 128a",
				Docstring: `
					128-bit FNV-1a hash.Hash. Its Sum method will lay the
					value out in big-endian byte order.
				`,
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
								Operator: "wasm.hash.fnv.Hash128a",
							},
						},
					},
				},
			},
			{
				Name:  "fnv-32",
				Title: "FNV 32",
				Docstring: `
					32-bit FNV-1 hash.Hash. Its Sum method will lay the
					value out in big-endian byte order.
				`,
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
								Operator: "wasm.hash.fnv.Hash32",
							},
						},
					},
				},
			},
			{
				Name:  "fnv-32a",
				Title: "FNV 32a",
				Docstring: `
					32-bit FNV-1a hash.Hash. Its Sum method will lay the
					value out in big-endian byte order.
				`,
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
								Operator: "wasm.hash.fnv.Hash32a",
							},
						},
					},
				},
			},
			{
				Name:  "fnv-64",
				Title: "FNV 64",
				Docstring: `
					64-bit FNV-1 hash.Hash. Its Sum method will lay the
					value out in big-endian byte order.
				`,
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
								Operator: "wasm.hash.fnv.Hash64",
							},
						},
					},
				},
			},
			{
				Name:  "fnv-64a",
				Title: "FNV 64a",
				Docstring: `
					64-bit FNV-1a hash.Hash. Its Sum method will lay the
					value out in big-endian byte order.
				`,
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
								Operator: "wasm.hash.fnv.Hash64a",
							},
						},
					},
				},
			},
		},
	})),
}
