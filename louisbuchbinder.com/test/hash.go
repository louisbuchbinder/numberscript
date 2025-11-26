package test

type hashTestcase struct {
	pkg         string
	tab         string
	argOperator string
	in          string
	out         string
}

var hashTestcasesAdler32 = []hashTestcase{
	{
		pkg:         "adler32",
		tab:         "checksum",
		argOperator: "from-text",
		in:          "Hello World!",
		out:         "474547262",
	},
	{
		pkg:         "adler32",
		tab:         "checksum",
		argOperator: "from-text",
		in:          "Wikipedia",
		out:         "300286872",
	},
	{
		pkg:         "crc32",
		tab:         "ieee",
		argOperator: "from-text",
		in:          "Hello World!",
		out:         "472456355",
	},
	{
		pkg:         "crc32",
		tab:         "castangnoli",
		argOperator: "from-text",
		in:          "Hello World!",
		out:         "4268552668",
	},
	{
		pkg:         "crc32",
		tab:         "koopman",
		argOperator: "from-text",
		in:          "Hello World!",
		out:         "1746547384",
	},
}
