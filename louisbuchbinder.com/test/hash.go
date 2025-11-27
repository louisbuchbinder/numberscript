package test

import "github.com/google/uuid"

var hashTestcasesAdler32 = []testcase{
	{
		id:     uuid.MustParse("6791FB63-05BA-4143-BC63-6561A360FAF3"),
		module: "hash",
		pkg:    "adler32",
		tab:    "checksum",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"474547262"},
	},
	{
		id:     uuid.MustParse("FC54FF19-415F-4A7B-A0E0-AA199600A49C"),
		module: "hash",
		pkg:    "adler32",
		tab:    "checksum",
		in:     []testcaseInput{{content: "Wikipedia", operator: "from-text"}},
		out:    []string{"300286872"},
	},
}

var hashTestcasesCrc32 = []testcase{
	{
		id:     uuid.MustParse("2CF267EE-6A34-492C-B0B6-8FFB06A871C0"),
		module: "hash",
		pkg:    "crc32",
		tab:    "ieee",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"472456355"},
	},
	{
		id:     uuid.MustParse("4C5F57F8-C564-4DAE-87DE-A14FF56A1A29"),
		module: "hash",
		pkg:    "crc32",
		tab:    "castangnoli",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"4268552668"},
	},
	{
		id:     uuid.MustParse("6533C1E1-6F29-4FFA-A5F6-8AE307AF5E0D"),
		module: "hash",
		pkg:    "crc32",
		tab:    "koopman",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"1746547384"},
	},
}

var hashTestcasesCrc64 = []testcase{
	{
		id:     uuid.MustParse("E610FE11-F409-4982-B26E-2AE39CA7E653"),
		module: "hash",
		pkg:    "crc64",
		tab:    "iso",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"9059499827269523000"},
	},
	{
		id:     uuid.MustParse("67621F15-F81B-4360-A621-3F9206D02E7E"),
		module: "hash",
		pkg:    "crc64",
		tab:    "ecma",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"8431954862038217000"},
	},
}

var hashTestcasesFnv = []testcase{
	{
		id:     uuid.MustParse("EAEB4CF5-B39D-49CB-B2E4-D22BFA4DCFBD"),
		module: "hash",
		pkg:    "fnv",
		tab:    "fnv-128",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"3e09e68d3967a18bcfaa7ac886326144"},
	},
	{
		id:     uuid.MustParse("770C8A30-F40D-4A03-8D49-0521D5364318"),
		module: "hash",
		pkg:    "fnv",
		tab:    "fnv-128a",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"d2d42892ede872031d2593366229c2d2"},
	},
	{
		id:     uuid.MustParse("4CD265E4-3CE0-4BD5-AAF9-8BA725E2F571"),
		module: "hash",
		pkg:    "fnv",
		tab:    "fnv-32",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"12a9a41c"},
	},
	{
		id:     uuid.MustParse("73ADB512-CDE9-45E9-AD44-95EB1C919AB4"),
		module: "hash",
		pkg:    "fnv",
		tab:    "fnv-32a",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"b1ea4872"},
	},
	{
		id:     uuid.MustParse("5BCD1126-A0EF-4AA6-9FE5-A0F391CC9341"),
		module: "hash",
		pkg:    "fnv",
		tab:    "fnv-64",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"8e59dd02f68c387c"},
	},
	{
		id:     uuid.MustParse("80CF93F0-5F21-418A-BC2C-979371A70CF9"),
		module: "hash",
		pkg:    "fnv",
		tab:    "fnv-64a",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"8c0ec8d1fb9e6e32"},
	},
}
