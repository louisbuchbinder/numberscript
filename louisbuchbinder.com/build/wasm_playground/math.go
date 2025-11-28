package wasm_playground

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/louisbuchbinder.com/templates"
)

type MathPlaygroundTabInput struct {
	Title     string
	Docstring template.HTML
	Args      []MathPlaygroundTabArgInput
}

type MathPlaygroundTabArgInput struct {
	Name     string
	Title    string
	Operator string
}

func MathPlaygroundTab(in MathPlaygroundTabInput) templates.WasmPlaygroundTab {
	return templates.WasmPlaygroundTab{
		Name:      strings.ToLower(in.Title),
		Title:     in.Title,
		Docstring: in.Docstring,
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
		Title:     "General Math",
		Menu:      Menu("Math", "General"),
		Docstring: "Basic mathematical functions.",
		Tabs: []templates.WasmPlaygroundTab{
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Abs", Docstring: "Abs returns the absolute value of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Acos", Docstring: "Acos returns the arccosine, in radians, of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Acosh", Docstring: "Acosh returns the inverse hyperbolic cosine of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Asin", Docstring: "Asin returns the arcsine, in radians, of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Asinh", Docstring: "Asinh returns the inverse hyperbolic sine of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Atan", Docstring: "Atan returns the arctangent, in radians, of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Atan2", Docstring: "Atan2 returns the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Atanh", Docstring: "Atanh returns the inverse hyperbolic tangent of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Cbrt", Docstring: "Cbrt returns the cube root of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Ceil", Docstring: "Ceil returns the least integer value greater than or equal to x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Copysign", Docstring: "Copysign returns a value with the magnitude of f and the sign of sign.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input F", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Sign", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Cos", Docstring: "Cos returns the cosine of the radian argument x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Cosh", Docstring: "Cosh returns the hyperbolic cosine of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Dim", Docstring: "Dim returns the maximum of x-y or 0.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erf", Docstring: "Erf returns the error function of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erfc", Docstring: "Erfc returns the complementary error function of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erfcinv", Docstring: "Erfcinv returns the inverse of Erfc(x).", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Erfinv", Docstring: "Erfinv returns the inverse error function of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Exp", Docstring: "Exp returns e**x, the base-e exponential of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Exp2", Docstring: "Exp2 returns 2**x, the base-2 exponential of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Expm1", Docstring: "Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "FMA", Docstring: "FMA returns x * y + z, computed with only one rounding. (That is, FMA returns the fused multiply-add of x, y, and z.)", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}, {Name: "input-z", Title: "Input Z", Operator: "safeFloat"}}}),
			// func Float32bits(f float32) uint32
			// func Float32frombits(b uint32) float32
			// func Float64bits(f float64) uint64
			// func Float64frombits(b uint64) float64
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Floor", Docstring: "Floor returns the greatest integer value less than or equal to x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			// func Frexp(f float64) (frac float64, exp int)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Gamma", Docstring: "Gamma returns the Gamma function of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Hypot", Docstring: "Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Ilogb", Docstring: "Ilogb returns the binary exponent of x as an integer.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "J0", Docstring: "J0 returns the order-zero Bessel function of the first kind.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "J1", Docstring: "J1 returns the order-one Bessel function of the first kind.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Jn", Docstring: "Jn returns the order-n Bessel function of the first kind.", Args: []MathPlaygroundTabArgInput{{Name: "input-n", Title: "Input N", Operator: "safeInt"}, {Name: "input-x", Title: "Input X", Operator: "safeFloat"}}}),
			// func Ldexp(frac float64, exp int) float64
			// func Lgamma(x float64) (lgamma float64, sign int)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log", Docstring: "Log returns the natural logarithm of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log10", Docstring: "Log10 returns the decimal logarithm of x. The special cases are the same as for Log.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log1p", Docstring: "Log1p returns the natural logarithm of 1 plus its argument x. It is more accurate than Log(1 + x) when x is near zero.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Log2", Docstring: "Log2 returns the binary logarithm of x. The special cases are the same as for Log.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Logb", Docstring: "Logb returns the binary exponent of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Max", Docstring: "Max returns the larger of x or y.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Min", Docstring: "Min returns the smaller of x or y.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Mod", Docstring: "Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			// func Modf(f float64) (int float64, frac float64)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Nextafter", Docstring: "Nextafter returns the next representable float64 value after x towards y.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Nextafter32", Docstring: "Nextafter32 returns the next representable float32 value after x towards y.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat32"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat32"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Pow", Docstring: "Pow returns x**y, the base-x exponential of y.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Pow10", Docstring: "Pow10 returns 10**n, the base-10 exponential of n.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeInt"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Remainder", Docstring: "Remainder returns the IEEE 754 floating-point remainder of x/y.", Args: []MathPlaygroundTabArgInput{{Name: "input-x", Title: "Input X", Operator: "safeFloat"}, {Name: "input-y", Title: "Input Y", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Round", Docstring: "Round returns the nearest integer, rounding half away from zero.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "RoundToEven", Docstring: "RoundToEven returns the nearest integer, rounding ties to even.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			// func Signbit(x float64) bool
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Sin", Docstring: "Sin returns the sine of the radian argument x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			// func Sincos(x float64) (sin, cos float64)
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Sinh", Docstring: "Sinh returns the hyperbolic sine of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Sqrt", Docstring: "Sqrt returns the square root of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Tan", Docstring: "Tan returns the tangent of the radian argument x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Tanh", Docstring: "Tanh returns the hyperbolic tangent of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Trunc", Docstring: "Trunc returns the integer value of x.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Y0", Docstring: "Y0 returns the order-zero Bessel function of the second kind.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Y1", Docstring: "Y1 returns the order-one Bessel function of the second kind.", Args: []MathPlaygroundTabArgInput{{Name: "input", Title: "Input", Operator: "safeFloat"}}}),
			MathPlaygroundTab(MathPlaygroundTabInput{Title: "Yn", Docstring: "Yn returns the order-n Bessel function of the second kind.", Args: []MathPlaygroundTabArgInput{{Name: "input-n", Title: "Input N", Operator: "safeInt"}, {Name: "input-x", Title: "Input X", Operator: "safeFloat"}}}),
		},
	})),
}
