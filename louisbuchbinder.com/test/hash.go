package test

var hashTestcasesAdler32 = []testcase{
	{
		module:      "hash",
		pkg:         "adler32",
		tab:         "checksum",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"474547262"},
	},
	{
		module:      "hash",
		pkg:         "adler32",
		tab:         "checksum",
		argOperator: "from-text",
		in:          []string{"Wikipedia"},
		out:         []string{"300286872"},
	},
}

var hashTestcasesCrc32 = []testcase{
	{
		module:      "hash",
		pkg:         "crc32",
		tab:         "ieee",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"472456355"},
	},
	{
		module:      "hash",
		pkg:         "crc32",
		tab:         "castangnoli",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"4268552668"},
	},
	{
		module:      "hash",
		pkg:         "crc32",
		tab:         "koopman",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"1746547384"},
	},
}

var hashTestcasesCrc64 = []testcase{
	{
		module:      "hash",
		pkg:         "crc64",
		tab:         "iso",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"9059499827269523000"},
	},
	{
		module:      "hash",
		pkg:         "crc64",
		tab:         "ecma",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"8431954862038217000"},
	},
}

var hashTestcasesFnv = []testcase{
	{
		module:      "hash",
		pkg:         "fnv",
		tab:         "fnv-128",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"3e09e68d3967a18bcfaa7ac886326144"},
	},
	{
		module:      "hash",
		pkg:         "fnv",
		tab:         "fnv-128a",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"d2d42892ede872031d2593366229c2d2"},
	},
	{
		module:      "hash",
		pkg:         "fnv",
		tab:         "fnv-32",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"12a9a41c"},
	},
	{
		module:      "hash",
		pkg:         "fnv",
		tab:         "fnv-32a",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"b1ea4872"},
	},
	{
		module:      "hash",
		pkg:         "fnv",
		tab:         "fnv-64",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"8e59dd02f68c387c"},
	},
	{
		module:      "hash",
		pkg:         "fnv",
		tab:         "fnv-64a",
		argOperator: "from-text",
		in:          []string{"Hello World!"},
		out:         []string{"8c0ec8d1fb9e6e32"},
	},
}
