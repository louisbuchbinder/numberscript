package wasm_playground

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

func MathPlaygroundTabSingleArgFloat64(title string) templates.WasmPlaygroundTab {
	return templates.WasmPlaygroundTab{
		Name:  strings.ToLower(title),
		Title: title,
		Args: []templates.WasmPlaygroundTabArg{
			{
				Type:  templates.WasmPlaygroundTabValType_Number,
				Name:  "data",
				Title: "Data",
				Operators: []templates.WasmPlaygroundTabOperator{
					{Name: "from-number", Title: "From Number", Operator: "safeFloat"},
				},
			},
		},
		Result: templates.WasmPlaygroundTabResult{
			Operators: []templates.WasmPlaygroundTabOperator{
				{
					Name:     "as-number",
					Title:    "As Number",
					Operator: fmt.Sprintf("wasm.math.%s", title),
				},
			},
		},
	}
}

var MathDocumentTemplateInput = templates.DocumentTemplateInput{
	Title: "Math",
	Scripts: template.HTML(strings.Join([]string{
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/external/go1.24.5_wasm_exec.js"})),
		string(templates.MustRenderScriptTemplate(templates.ScriptTemplateInput{Src: "/wasm/math/pkg/wasm.js"})), // TODO: use the hash-named file
	}, "\n")),
	Main: template.HTML(templates.MustRenderWasmPlaygroundTemplate(templates.WasmPlaygroundTemplateInput{
		Title: "General Math",
		Menu:  Menu("Math", "General"),
		Tabs: []templates.WasmPlaygroundTab{
			MathPlaygroundTabSingleArgFloat64("Abs"),
			MathPlaygroundTabSingleArgFloat64("Acos"),
			MathPlaygroundTabSingleArgFloat64("Acosh"),
			MathPlaygroundTabSingleArgFloat64("Asin"),
			MathPlaygroundTabSingleArgFloat64("Asinh"),
			MathPlaygroundTabSingleArgFloat64("Atan"),
			// func Atan2(y, x float64) float64
			MathPlaygroundTabSingleArgFloat64("Atanh"),
			MathPlaygroundTabSingleArgFloat64("Cbrt"),
			MathPlaygroundTabSingleArgFloat64("Ceil"),
			// func Copysign(f, sign float64) float64
			MathPlaygroundTabSingleArgFloat64("Cos"),
			MathPlaygroundTabSingleArgFloat64("Cosh"),
			// func Dim(x, y float64) float64
			MathPlaygroundTabSingleArgFloat64("Erf"),
			MathPlaygroundTabSingleArgFloat64("Erfc"),
			MathPlaygroundTabSingleArgFloat64("Erfcinv"),
			MathPlaygroundTabSingleArgFloat64("Erfinv"),
			MathPlaygroundTabSingleArgFloat64("Exp"),
			MathPlaygroundTabSingleArgFloat64("Exp2"),
			MathPlaygroundTabSingleArgFloat64("Expm1"),
			// func FMA(x, y, z float64) float64
			// func Float32bits(f float32) uint32
			// func Float32frombits(b uint32) float32
			// func Float64bits(f float64) uint64
			// func Float64frombits(b uint64) float64
			MathPlaygroundTabSingleArgFloat64("Floor"),
			// func Frexp(f float64) (frac float64, exp int)
			MathPlaygroundTabSingleArgFloat64("Gamma"),
			// func Hypot(p, q float64) float64
			// func Ilogb(x float64) int
			// func Inf(sign int) float64
			// func IsInf(f float64, sign int) bool
			// func IsNaN(f float64) (is bool)
			MathPlaygroundTabSingleArgFloat64("J0"),
			MathPlaygroundTabSingleArgFloat64("J1"),
			// func Jn(n int, x float64) float64
			// func Ldexp(frac float64, exp int) float64
			// func Lgamma(x float64) (lgamma float64, sign int)
			MathPlaygroundTabSingleArgFloat64("Log"),
			MathPlaygroundTabSingleArgFloat64("Log10"),
			MathPlaygroundTabSingleArgFloat64("Log1p"),
			MathPlaygroundTabSingleArgFloat64("Log2"),
			MathPlaygroundTabSingleArgFloat64("Logb"),
			// func Max(x, y float64) float64
			// func Min(x, y float64) float64
			// func Mod(x, y float64) float64
			// func Modf(f float64) (int float64, frac float64)
			// func NaN() float64
			// func Nextafter(x, y float64) (r float64)
			// func Nextafter32(x, y float32) (r float32)
			// func Pow(x, y float64) float64
			// func Pow10(n int) float64
			// func Remainder(x, y float64) float64
			MathPlaygroundTabSingleArgFloat64("Round"),
			MathPlaygroundTabSingleArgFloat64("RoundToEven"),
			// func Signbit(x float64) bool
			MathPlaygroundTabSingleArgFloat64("Sin"),
			// func Sincos(x float64) (sin, cos float64)
			MathPlaygroundTabSingleArgFloat64("Sinh"),
			MathPlaygroundTabSingleArgFloat64("Sqrt"),
			MathPlaygroundTabSingleArgFloat64("Tan"),
			MathPlaygroundTabSingleArgFloat64("Tanh"),
			MathPlaygroundTabSingleArgFloat64("Trunc"),
			MathPlaygroundTabSingleArgFloat64("Y0"),
			MathPlaygroundTabSingleArgFloat64("Y1"),
			// func Yn(n int, x float64) float64
		},
	})),
}
