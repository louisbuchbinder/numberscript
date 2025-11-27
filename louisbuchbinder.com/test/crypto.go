package test

var cryptoTestcasesMd5 = []testcase{
	{
		module:      "crypto",
		pkg:         "md5",
		tab:         "hash",
		argOperator: "from-test",
		in:          []string{"Hello World!"},
		out:         []string{"ed076287532e86365e841e92bfc50d8c"},
	},
	{
		module:      "crypto",
		pkg:         "md5",
		tab:         "hash",
		argOperator: "from-test",
		in:          []string{"The quick brown fox jumps over the lazy dog"},
		out:         []string{"9e107d9d372bb6826bd81d3542a419d6"},
	},
}
