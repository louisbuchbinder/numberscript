package main

import (
	"flag"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/louisbuchbinder/core/lib/util"
)

const JS_DEFINE_EXPORT_TREE_TEMPLATE = `    if (typeof window.{{.ObjPath}} === "undefined") {
        window.{{.ObjPath}} = {};
    }`

type JsDefineExportTreeTemplateInput struct {
	ObjPath string
}

const WASM_JS_TEMPLATE = `
if (typeof window.wasmContentLoaders === "undefined") {
    window.wasmContentLoaders = [];
}
window.wasmContentLoaders.push((function () {
{{.DefineTree}}
    const go = new Go();
    return WebAssembly.instantiateStreaming(
        fetch("/{{.Module}}/pkg/{{.SHA256Prefix}}main.wasm"),
        go.importObject
    ).then((result) => {
        go.run(result.instance);
    });
})());`

type WasmJsTemplateInput struct {
	DefineTree   string
	Module       string
	SHA256Prefix string
}
type stringArray []string

func (a *stringArray) String() string {
	return strings.Join(*a, ",")
}

func (a *stringArray) Set(value string) error {
	*a = append(*a, value)
	return nil
}

var (
	jsDefineExportTreeTemplate *template.Template
	wasmJsTemplate             *template.Template

	wasmExports  stringArray
	module       string
	wasmGo       string
	output       string
	outputSHA256 string
)

func init() {
	jsDefineExportTreeTemplate = template.Must(template.New("JS_DEFINE_EXPORT_TREE_TEMPLATE").Parse(JS_DEFINE_EXPORT_TREE_TEMPLATE))
	wasmJsTemplate = template.Must(template.New("WASM_JS_TEMPLATE").Parse(WASM_JS_TEMPLATE))

	flag.StringVar(&module, "module", "", "[required] the module path for the wasm js script load")
	flag.StringVar(&wasmGo, "wasm-go", "", "[required] the path to the wasm-go file")
	flag.Var(&wasmExports, "wasm-export", "append multiple wasm exports")
	flag.StringVar(&output, "output", "", "[required] the output file path")
	flag.StringVar(&outputSHA256, "output-sha256", "", "[required] the output sha256 file path")
	flag.Parse()
	if module == "" || wasmGo == "" || output == "" || outputSHA256 == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func buildWasmExportParts(wasmExportsList []string) []string {
	if len(wasmExportsList) == 0 {
		return nil
	}

	seen := make(map[string]bool, 0)
	for _, exp := range wasmExportsList {
		parts := strings.Split(exp, ".")
		for i := 1; i < len(parts); i++ {
			prefix := strings.Join(parts[:i], ".")
			seen[prefix] = true
		}
	}

	wasmExportParts := make([]string, 0, len(seen))
	for k := range seen {
		wasmExportParts = append(wasmExportParts, k)
	}

	sort.Strings(wasmExportParts)
	return wasmExportParts
}

func main() {
	wasmExportParts := buildWasmExportParts(wasmExports)
	defineTree := strings.Join(util.Must(util.MapOrError(wasmExportParts, func(_ int, objPath string) (string, error) {
		b, err := util.ExecuteTemplate(jsDefineExportTreeTemplate, JsDefineExportTreeTemplateInput{
			ObjPath: objPath,
		})
		return string(b), err
	})), "\n")
	wasmJs := util.Must(util.ExecuteTemplate(wasmJsTemplate, WasmJsTemplateInput{
		DefineTree:   defineTree,
		Module:       module,
		SHA256Prefix: "",
	}))

	wasmGoFile := util.Must(os.Open(wasmGo))
	defer wasmGoFile.Close()

	sha256 := util.Must(util.Sha256HexOfFile(wasmGoFile))

	wasmJsSHA256 := util.Must(util.ExecuteTemplate(wasmJsTemplate, WasmJsTemplateInput{
		DefineTree:   defineTree,
		Module:       module,
		SHA256Prefix: sha256 + ".",
	}))

	util.Must0(os.WriteFile(output, wasmJs, 0o644))
	util.Must0(os.WriteFile(outputSHA256, wasmJsSHA256, 0o644))
}
