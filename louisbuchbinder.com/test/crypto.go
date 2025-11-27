package test

import "github.com/google/uuid"

var cryptoTestcasesMd5 = []testcase{
	{
		id:     uuid.MustParse("7B6142D2-EDCE-49D5-9A1E-1112F5B3C529"),
		module: "crypto",
		pkg:    "md5",
		tab:    "hash",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"ed076287532e86365e841e92bfc50d8c"},
	},
	{
		id:     uuid.MustParse("BB1BE48B-C6CE-4212-A304-90F1C4015269"),
		module: "crypto",
		pkg:    "md5",
		tab:    "hash",
		in:     []testcaseInput{{content: "The quick brown fox jumps over the lazy dog", operator: "from-text"}},
		out:    []string{"9e107d9d372bb6826bd81d3542a419d6"},
	},
}

var cryptoTestcasesSha1 = []testcase{
	{
		id:     uuid.MustParse("A39F8969-8AA8-4D87-83AD-4A582670D3F7"),
		module: "crypto",
		pkg:    "sha1",
		tab:    "hash",
		in:     []testcaseInput{{content: "Hello World!", operator: "from-text"}},
		out:    []string{"2ef7bde608ce5404e97d5f042f95f89f1c232871"},
	},
	{
		id:     uuid.MustParse("F3BE1707-5014-49D2-8263-AAF3F0BE2C1E"),
		module: "crypto",
		pkg:    "sha1",
		tab:    "hash",
		in:     []testcaseInput{{content: "The quick brown fox jumps over the lazy dog", operator: "from-text"}},
		out:    []string{"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"},
	},
}

var cryptoTestcasesSha256 = []testcase{
	{
		id:     uuid.MustParse("569CB8E9-0884-4497-9AF7-7C73723CDD60"),
		module: "crypto",
		pkg:    "sha256",
		tab:    "sum224-hash",
		in:     []testcaseInput{{content: "", operator: "from-text"}},
		out:    []string{"d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f"},
	},
	{
		id:     uuid.MustParse("56C2990F-85F9-4182-8582-C2558597ACC8"),
		module: "crypto",
		pkg:    "sha256",
		tab:    "sum224-hash",
		in:     []testcaseInput{{content: "pack my box with five dozen liquor jugs", operator: "from-text"}},
		out:    []string{"c78caba92774efee0a7fbff50315ab3fae393edc92dd57ca6015fa26"},
	},
	{
		id:     uuid.MustParse("C7E954A6-505C-41BA-B2F0-9D50DCCAADFD"),
		module: "crypto",
		pkg:    "sha256",
		tab:    "sum256-hash",
		in:     []testcaseInput{{content: "", operator: "from-text"}},
		out:    []string{"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	},
	{
		id:     uuid.MustParse("8F10CFD4-6F9B-4EB6-ACD7-2665AC45819D"),
		module: "crypto",
		pkg:    "sha256",
		tab:    "sum256-hash",
		in:     []testcaseInput{{content: "pack my box with five dozen liquor jugs", operator: "from-text"}},
		out:    []string{"c826b48f8387e1436da6af000625071279ad55628460ee16fab5697afa7fac94"},
	},
}

var cryptoTestcasesSha512 = []testcase{
	{
		id:     uuid.MustParse("D3E71779-99AA-4FFA-855E-C6A7AF631C14"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum512_224-hash",
		in:     []testcaseInput{{content: "", operator: "from-text"}},
		out:    []string{"6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4"},
	},
	{
		id:     uuid.MustParse("065C2BDA-3CA0-480D-8203-B381CC674493"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum512_224-hash",
		in:     []testcaseInput{{content: "pack my box with five dozen liquor jugs", operator: "from-text"}},
		out:    []string{"127a506396d7247dc6186ab998f873cfa7b41c0f162b549201bda3c4"},
	},
	{
		id:     uuid.MustParse("9B8F832A-2BAC-418D-9CFB-59E3D0221952"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum512_256-hash",
		in:     []testcaseInput{{content: "", operator: "from-text"}},
		out:    []string{"c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a"},
	},
	{
		id:     uuid.MustParse("D08205AF-4380-49C1-B9C9-DB4113634984"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum512_256-hash",
		in:     []testcaseInput{{content: "pack my box with five dozen liquor jugs", operator: "from-text"}},
		out:    []string{"eb36fb032081478da18b17877f90f58800518a96328fb30f6a06529bd1a39e0d"},
	},
	{
		id:     uuid.MustParse("DA194F25-B03D-4E24-A05B-F6490A580E97"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum384-hash",
		in:     []testcaseInput{{content: "", operator: "from-text"}},
		out:    []string{"38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b"},
	},
	{
		id:     uuid.MustParse("39D838C2-8263-45DC-B740-A208DD8EF336"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum384-hash",
		in:     []testcaseInput{{content: "pack my box with five dozen liquor jugs", operator: "from-text"}},
		out:    []string{"959bd183db4d4c77551c49c166df2412be1e0da4ac6fe807dab4da6ef6f4037ad358f1a425f0d0aac1c0da286696e12a"},
	},
	{
		id:     uuid.MustParse("442BF9A6-1102-46BF-A0C2-9E76066E6397"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum512-hash",
		in:     []testcaseInput{{content: "", operator: "from-text"}},
		out:    []string{"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
	},
	{
		id:     uuid.MustParse("B977089E-428C-45E2-B70A-51ACA22DDA91"),
		module: "crypto",
		pkg:    "sha512",
		tab:    "sum512-hash",
		in:     []testcaseInput{{content: "pack my box with five dozen liquor jugs", operator: "from-text"}},
		out:    []string{"db3eaf174b6264e6a37f9fe26738876a1efecc7c8556ee9da8257bea9b3cb0caa66da1c072c66b8088d35de14cd0528ce15e4a80f61610e5978e39fdd726c529"},
	},
}
