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
)

type WasmPlaygroundTabArg struct {
	Name      string
	Title     string
	Type      WasmPlaygroundTabValType
	Operators []WasmPlaygroundTabOperator
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
