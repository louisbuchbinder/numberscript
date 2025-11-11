package wasm_playground

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

type MathPlaygroundTabInput struct {
	Title string
	Args  []MathPlaygroundTabArgInput
}

type MathPlaygroundTabArgInput struct {
	Name     string
	Title    string
	Operator string
}

func MathPlaygroundTab(in MathPlaygroundTabInput) templates.WasmPlaygroundTab {
	return templates.WasmPlaygroundTab{
		Name:  strings.ToLower(in.Title),
		Title: in.Title,
		Args: util.Map(in.Args, func(_ int, a MathPlaygroundTabArgInput) templates.WasmPlaygroundTabArg {
			return templates.WasmPlaygroundTabArg{
				Type:  templates.WasmPlaygroundTabValType_Number,
				Name:  a.Name,
				Title: a.Title,
				Operators: []templates.WasmPlaygroundTabOperator{
					{Name: "from-number", Title: "From Number", Operator: a.Operator},
				},
			}
		}),
		Results: []templates.WasmPlaygroundTabResult{
			{
				Operators: []templates.WasmPlaygroundTabOperator{
					{
						Name:     "as-number",
						Title:    "As Number",
						Operator: fmt.Sprintf("wasm.math.%s", in.Title),
					},
				},
			},
		},
	}
}

var MathDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Math",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: WASM_GO_SCRIPT_SRC})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/math/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "General Math",
		Menu:  Menu("Math", "General"),
		Tabs: []templates.WasmPlaygroundTab{
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Abs", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Acos", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Acosh", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Asin", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Asinh", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Atan", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Atan2", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Atanh", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Cbrt", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Ceil", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Copysign", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input F", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Sign", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Cos", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Cosh", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Dim", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erf", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erfc", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erfcinv", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erfinv", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Exp", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Exp2", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Expm1", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "FMA", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}, {Name: "input-z", Title: "Input Z", Operator: "safeFloat"}}}),
			// func Float32bits(f float32) uint32
			// func Float32frombits(b uint32) float32
			// func Float64bits(f float64) uint64
			// func Float64frombits(b uint64) float64
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Floor", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			// func Frexp(f float64) (frac float64, exp int)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Gamma", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Hypot", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Ilogb", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "J0", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "J1", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Jn", Args: []MathPlaygroundTabArgInput{{Name: "input-n", Title: "Input N", Operator: "safeInt"}, {Name: "input-x", Title: "Input X", Operator: "safeFloat"}}}),
			// func Ldexp(frac float64, exp int) float64
			// func Lgamma(x float64) (lgamma float64, sign int)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log10", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log1p", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log2", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Logb", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Max", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Min", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Mod", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			// func Modf(f float64) (int float64, frac float64)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Nextafter", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Nextafter32", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat32"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat32"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Pow", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Pow10", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeInt"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Remainder", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Round", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "RoundToEven", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			// func Signbit(x float64) bool
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Sin", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			// func Sincos(x float64) (sin, cos float64)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Sinh", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Sqrt", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Tan", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Tanh", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Trunc", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Y0", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Y1", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Yn", Args: []MathPlaygroundTabArgInput{{Name: "input-n", Title: "Input N", Operator: "safeInt"}, {Name: "input-x", Title: "Input X", Operator: "safeFloat"}}}),
		},
	})),
}
