package main

import (
	"html/template"
	"os"
	"path"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/build/wasm_playground"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

func Page(Main template.HTML) []byte {
	return templates.MustRenderDocumentTemplate(templates.DocumentTemplateInput{
		Title:   "Home",
		Scripts: "",
		Navbar:  template.HTML(templates.MustRenderNavbarTemplate(templates.NavbarTemplateInput{})),
		Main:    Main,
	})
}

var MainPage = Page(template.HTML(templates.MustRenderHomeTemplate(templates.HomeTemplateInput{})))

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

	util.Must0(write("crypto/md5/index.html", Page(template.HTML(wasm_playground.CryptoMD5Page))))
	util.Must0(write("crypto/rand/index.html", Page(template.HTML(wasm_playground.CryptoRandPage))))
	util.Must0(write("crypto/sha1/index.html", Page(template.HTML(wasm_playground.CryptoSHA1Page))))
	util.Must0(write("crypto/sha3/index.html", Page(template.HTML(wasm_playground.CryptoSHA3Page))))
	util.Must0(write("crypto/sha256/index.html", Page(template.HTML(wasm_playground.CryptoSHA256Page))))
	util.Must0(write("crypto/sha512/index.html", Page(template.HTML(wasm_playground.CryptoSHA512Page))))

	util.Must0(write("encoding/base32/index.html", Page(template.HTML(wasm_playground.EncodingBase32Page))))
	util.Must0(write("encoding/base64/index.html", Page(template.HTML(wasm_playground.EncodingBase64Page))))
	util.Must0(write("encoding/hex/index.html", Page(template.HTML(wasm_playground.EncodingHexPage))))
	util.Must0(write("encoding/html/index.html", Page(template.HTML(wasm_playground.EncodingHTMLPage))))
	util.Must0(write("encoding/uri/index.html", Page(template.HTML(wasm_playground.EncodingURIPage))))
}
