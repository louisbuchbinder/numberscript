package test

type encodingTestcase struct {
	pkg string
	tab string
	in  string
	out string
}

var encodeTestcasesBase32 = []encodingTestcase{
	{
		pkg: "base32",
		tab: "encode",
		in:  "",
		out: "",
	},
	{
		pkg: "base32",
		tab: "encode",
		in:  "abcdefghijklmnopqrstuvwxyz0123456789",
		out: "MFRGGZDFMZTWQ2LKNNWG23TPOBYXE43UOV3HO6DZPIYDCMRTGQ2TMNZYHE======",
	},
	{
		pkg: "base32",
		tab: "encode",
		in:  "The quick fox jumps over the lazy dog.",
		out: "KRUGKIDROVUWG2ZAMZXXQIDKOVWXA4ZAN53GK4RAORUGKIDMMF5HSIDEN5TS4===",
	},
	{
		pkg: "base32",
		tab: "decode",
		in:  "MFRGGZDFMZTWQ2LKNNWG23TPOBYXE43UOV3HO6DZPIYDCMRTGQ2TMNZYHE======",
		out: "abcdefghijklmnopqrstuvwxyz0123456789",
	},
	{
		pkg: "base32",
		tab: "decode",
		in:  "KRUGKIDROVUWG2ZAMZXXQIDKOVWXA4ZAN53GK4RAORUGKIDMMF5HSIDEN5TS4===",
		out: "The quick fox jumps over the lazy dog.",
	},
}

var encodeTestcasesBase64 = []encodingTestcase{
	{
		pkg: "base64",
		tab: "encode",
		in:  "",
		out: "",
	},
	{
		pkg: "base64",
		tab: "encode",
		in:  "abcdefghijklmnopqrstuvwxyz0123456789",
		out: "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXowMTIzNDU2Nzg5",
	},
	{
		pkg: "base64",
		tab: "encode",
		in:  "The quick fox jumps over the lazy dog.",
		out: "VGhlIHF1aWNrIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=",
	},
	{
		pkg: "base64",
		tab: "decode",
		in:  "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXowMTIzNDU2Nzg5",
		out: "abcdefghijklmnopqrstuvwxyz0123456789",
	},
	{
		pkg: "base64",
		tab: "decode",
		in:  "VGhlIHF1aWNrIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=",
		out: "The quick fox jumps over the lazy dog.",
	},
}

var encodeTestcasesHex = []encodingTestcase{
	{
		pkg: "hex",
		tab: "encode",
		in:  "",
		out: "",
	},
	{
		pkg: "hex",
		tab: "encode",
		in:  "abcdefghijklmnopqrstuvwxyz0123456789",
		out: "6162636465666768696a6b6c6d6e6f707172737475767778797a30313233343536373839",
	},
	{
		pkg: "hex",
		tab: "encode",
		in:  "The quick fox jumps over the lazy dog.",
		out: "54686520717569636b20666f78206a756d7073206f76657220746865206c617a7920646f672e",
	},
	{
		pkg: "hex",
		tab: "decode",
		in:  "6162636465666768696a6b6c6d6e6f707172737475767778797a30313233343536373839",
		out: "abcdefghijklmnopqrstuvwxyz0123456789",
	},
	{
		pkg: "hex",
		tab: "decode",
		in:  "54686520717569636b20666f78206a756d7073206f76657220746865206c617a7920646f672e",
		out: "The quick fox jumps over the lazy dog.",
	},
}

var encodeTestcasesHtml = []encodingTestcase{
	{
		pkg: "html",
		tab: "escape",
		in:  "",
		out: "",
	},
	{
		pkg: "html",
		tab: "escape",
		in:  "<h1>Hello World!</h1>",
		out: "&lt;h1&gt;Hello World!&lt;/h1&gt;",
	},
	{
		pkg: "html",
		tab: "escape",
		in:  "<div><p>The <i>quick</i> fox <b>jumps</b> over the lazy dog.</p></div>",
		out: "&lt;div&gt;&lt;p&gt;The &lt;i&gt;quick&lt;/i&gt; fox &lt;b&gt;jumps&lt;/b&gt; over the lazy dog.&lt;/p&gt;&lt;/div&gt;",
	},
	{
		pkg: "html",
		tab: "unescape",
		in:  "&lt;h1&gt;Hello World!&lt;/h1&gt;",
		out: "<h1>Hello World!</h1>",
	},
	{
		pkg: "html",
		tab: "unescape",
		in:  "&lt;div&gt;&lt;p&gt;The &lt;i&gt;quick&lt;/i&gt; fox &lt;b&gt;jumps&lt;/b&gt; over the lazy dog.&lt;/p&gt;&lt;/div&gt;",
		out: "<div><p>The <i>quick</i> fox <b>jumps</b> over the lazy dog.</p></div>",
	},
}

var encodeTestcasesUri = []encodingTestcase{
	{
		pkg: "uri",
		tab: "encode-uri",
		in:  "",
		out: "",
	},
	{
		pkg: "uri",
		tab: "encode-uri",
		in:  "https://www.example.com/?pie=π&emoji=😁",
		out: "https://www.example.com/?pie=%CF%80&emoji=%F0%9F%98%81",
	},
	{
		pkg: "uri",
		tab: "decode-uri",
		in:  "",
		out: "",
	},
	{
		pkg: "uri",
		tab: "decode-uri",
		in:  "https://www.example.com/?pie=%CF%80&emoji=%F0%9F%98%81",
		out: "https://www.example.com/?pie=π&emoji=😁",
	},
	{
		pkg: "uri",
		tab: "encode-uri-component",
		in:  "",
		out: "",
	},
	{
		pkg: "uri",
		tab: "encode-uri-component",
		in:  "hello world!",
		out: "hello%20world!",
	},
	{
		pkg: "uri",
		tab: "encode-uri-component",
		in:  "`1234567890-=+_)(*&^%$#@!~",
		out: "%601234567890-%3D%2B_)(*%26%5E%25%24%23%40!~",
	},
	{
		pkg: "uri",
		tab: "decode-uri-component",
		in:  "",
		out: "",
	},
	{
		pkg: "uri",
		tab: "decode-uri-component",
		in:  "hello%20world!",
		out: "hello world!",
	},
	{
		pkg: "uri",
		tab: "decode-uri-component",
		in:  "%601234567890-%3D%2B_)(*%26%5E%25%24%23%40!~",
		out: "`1234567890-=+_)(*&^%$#@!~",
	},
}
