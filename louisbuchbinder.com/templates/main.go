package templates

import (
	"encoding/json"
	"html/template"

	"github.com/louisbuchbinder/core/lib/util"
	_ "github.com/louisbuchbinder/core/lib/util"
)

var FuncMap = template.FuncMap{
	"jsonMarshal": jsonMarshal,
}

func jsonMarshal(a any) string {
	return string(util.Must(json.Marshal(a)))
}

type WasmPlaygroundTab struct {
	Name              string
	Title             string
	Docstring         template.HTML
	HasGenerateButton bool
	Args              []WasmPlaygroundTabArg
	Results           []WasmPlaygroundTabResult
}

type WasmPlaygroundTabValType string

const (
	WasmPlaygroundTabValType_Text     WasmPlaygroundTabValType = "text"
	WasmPlaygroundTabValType_Number   WasmPlaygroundTabValType = "number"
	WasmPlaygroundTabValType_File     WasmPlaygroundTabValType = "file"
	WasmPlaygroundTabValType_Files    WasmPlaygroundTabValType = "files"
	WasmPlaygroundTabValType_Download WasmPlaygroundTabValType = "download"
)

type WasmPlaygroundTabArg struct {
	Name      string
	Title     string
	Type      WasmPlaygroundTabValType
	Operators []WasmPlaygroundTabOperator
	Options   WasmPlaygroundTabArgOptions
}

type WasmPlaygroundTabArgOptions struct {
	TextOptions   *WasmPlaygroundTabArgOptions_Text
	NumberOptions *WasmPlaygroundTabArgOptions_Number
}

type WasmPlaygroundTabArgOptions_Text struct{}

type WasmPlaygroundTabArgOptions_Number struct {
	Min float64
	Max float64
}

type WasmPlaygroundTabResult struct {
	Title     string
	Type      WasmPlaygroundTabValType
	Operators []WasmPlaygroundTabOperator
}

type WasmPlaygroundTabOperator struct {
	Name     string
	Title    string
	Operator string
}

type WasmPlaygroundMenuItemContainer struct {
	Title string
	Items []WasmPlaygroundMenuItem
}

type WasmPlaygroundMenuItem struct {
	Url   string
	Title string
}

type WasmPlaygroundMainMenuItem struct {
	Url         string
	Key         string
	Description string
}

type ServiceWorkerRoute struct {
	Path    string `json:"path"`
	Handler string `json:"handler"`
}
