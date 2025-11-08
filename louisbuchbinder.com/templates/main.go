package templates

import (
	_ "github.com/louisbuchbinder/core/lib/util"
)

type WasmPlaygroundTab struct {
	Name   string
	Title  string
	Args   []WasmPlaygroundTabArg
	Result WasmPlaygroundTabResult
}

type WasmPlaygroundTabValType string

const (
	WasmPlaygroundTabValType_Text WasmPlaygroundTabValType = "text"
	WasmPlaygroundTabValType_Int  WasmPlaygroundTabValType = "int"
)

type WasmPlaygroundTabArg struct {
	Name      string
	Title     string
	Type      WasmPlaygroundTabValType
	Operators []WasmPlaygroundTabOperator
	Options   WasmPlaygroundTabArgOptions
}

type WasmPlaygroundTabArgOptions struct {
	TextOptions WasmPlaygroundTabArgOptions_Text
	IntOptions  WasmPlaygroundTabArgOptions_Int
}

type WasmPlaygroundTabArgOptions_Text struct{}

type WasmPlaygroundTabArgOptions_Int struct {
	Min int
	Max int
}

type WasmPlaygroundTabResult struct {
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
	Key   string
}
