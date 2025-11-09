package main

import (
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/build/wasm_playground"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

func Page(in templates.DocumentTemplateInput) []byte {
	return templates.MustRenderDocumentTemplate(templates.DocumentTemplateInput{
		Title: in.Title,
		Links: template.HTML(strings.Join([]string{
			string(templates.MustRenderLinkTemplate(templates.LinkTemplateInput{Rel: "stylesheet", Href: "/external/bulma-1.0.4/css/bulma.css"})),
			string(templates.MustRenderLinkTemplate(templates.LinkTemplateInput{Rel: "stylesheet", Href: "/external/font-awesome-7.1.0/css/all.css"})),
			string(templates.MustRenderLinkTemplate(templates.LinkTemplateInput{Rel: "stylesheet", Href: "/css/site.css"})),
			string(in.Links),
		}, "\n")),
		Scripts: template.HTML(strings.Join([]string{
			string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/js/site.js"})),
			string(in.Scripts),
		}, "\n")),
		Navbar: template.HTML(templates.MustRenderNavbarTemplate(templates.NavbarTemplateInput{})),
		Main:   in.Main,
	})
}

var MainPage = Page(templates.DocumentTemplateInput{
	Title: "Main",
	Main:  template.HTML(templates.MustRenderHomeTemplate(templates.HomeTemplateInput{})),
})

var UtilityPlaygroundPage = Page(templates.DocumentTemplateInput{
	Title: "Utility Playground",
	Main: template.HTML(templates.MustRenderWasmPlaygroundMainTemplate(templates.WasmPlaygroundMainTemplateInput{
		Title: "Utility Playground",
		Menu:  wasm_playground.Menu("", ""),
		Items: wasm_playground.WasmPlaygroundMainMenuItem,
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
	util.Must0(write("utility-playground.html", UtilityPlaygroundPage))

	util.Must0(write("crypto/md5/index.html", Page(wasm_playground.CryptoMD5DocumentTemplateInput)))
	util.Must0(write("crypto/rand/index.html", Page(wasm_playground.CryptoRandDocumentTemplateInput)))
	util.Must0(write("crypto/sha1/index.html", Page(wasm_playground.CryptoSHA1DocumentTemplateInput)))
	util.Must0(write("crypto/sha3/index.html", Page(wasm_playground.CryptoSHA3DocumentTemplateInput)))
	util.Must0(write("crypto/sha256/index.html", Page(wasm_playground.CryptoSHA256DocumentTemplateInput)))
	util.Must0(write("crypto/sha512/index.html", Page(wasm_playground.CryptoSHA512DocumentTemplateInput)))

	util.Must0(write("encoding/base32/index.html", Page(wasm_playground.EncodingBase32DocumentTemplateInput)))
	util.Must0(write("encoding/base64/index.html", Page(wasm_playground.EncodingBase64DocumentTemplateInput)))
	util.Must0(write("encoding/hex/index.html", Page(wasm_playground.EncodingHexDocumentTemplateInput)))
	util.Must0(write("encoding/html/index.html", Page(wasm_playground.EncodingHTMLDocumentTemplateInput)))
	util.Must0(write("encoding/uri/index.html", Page(wasm_playground.EncodingURIDocumentTemplateInput)))
}
