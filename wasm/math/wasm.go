package wasmmath

import (
	"math"

	"github.com/louisbuchbinder/core/wasm"
)

func Abs(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Abs(args[0].Float())
	return float64(v)
}

func Acos(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Acos(args[0].Float())
	return float64(v)
}

func Acosh(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Acosh(args[0].Float())
	return float64(v)
}

func Asin(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Asin(args[0].Float())
	return float64(v)
}

func Asinh(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Asinh(args[0].Float())
	return float64(v)
}

func Atan(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Atan(args[0].Float())
	return float64(v)
}

func Atan2(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Atan2(args[0].Float(), args[1].Float())
	return float64(v)
}

func Atanh(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Atanh(args[0].Float())
	return float64(v)
}

func Cbrt(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Cbrt(args[0].Float())
	return float64(v)
}

func Ceil(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Ceil(args[0].Float())
	return float64(v)
}

func Copysign(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Copysign(args[0].Float(), args[1].Float())
	return float64(v)
}

func Cos(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Cos(args[0].Float())
	return float64(v)
}

func Cosh(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Cosh(args[0].Float())
	return float64(v)
}

func Dim(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Dim(args[0].Float(), args[1].Float())
	return float64(v)
}

func Erf(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Erf(args[0].Float())
	return float64(v)
}

func Erfc(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Erfc(args[0].Float())
	return float64(v)
}

func Erfcinv(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Erfcinv(args[0].Float())
	return float64(v)
}

func Erfinv(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Erfinv(args[0].Float())
	return float64(v)
}

func Exp(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Exp(args[0].Float())
	return float64(v)
}

func Exp2(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Exp2(args[0].Float())
	return float64(v)
}

func Expm1(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Expm1(args[0].Float())
	return float64(v)
}

func FMA(args []wasm.Value) any {
	if (len(args)) < 3 {
		return nil
	}
	v := math.FMA(args[0].Float(), args[1].Float(), args[2].Float())
	return float64(v)
}

func Float32bits(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Float32bits(float32(args[0].Float()))
	return uint32(v)
}

func Float32frombits(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Float32frombits(uint32(args[0].Int()))
	return float32(v)
}

func Float64bits(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Float64bits(args[0].Float())
	return uint64(v)
}

func Float64frombits(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Float64frombits(uint64(args[0].Int()))
	return float64(v)
}

func Floor(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Floor(args[0].Float())
	return float64(v)
}

// func Frexp(f float64) (frac float64, exp int)

func Gamma(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Gamma(args[0].Float())
	return float64(v)
}

func Hypot(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Hypot(args[0].Float(), args[1].Float())
	return float64(v)
}

func Ilogb(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Ilogb(args[0].Float())
	return int(v)
}

func J0(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.J0(args[0].Float())
	return float64(v)
}

func J1(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.J1(args[0].Float())
	return float64(v)
}

func Jn(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Jn(args[0].Int(), args[1].Float())
	return float64(v)
}

func Ldexp(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Ldexp(args[0].Float(), args[1].Int())
	return float64(v)
}

// func Lgamma(x float64) (lgamma float64, sign int)

func Log(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Log(args[0].Float())
	return float64(v)
}

func Log10(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Log10(args[0].Float())
	return float64(v)
}

func Log1p(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Log1p(args[0].Float())
	return float64(v)
}

func Log2(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Log2(args[0].Float())
	return float64(v)
}

func Logb(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Logb(args[0].Float())
	return float64(v)
}

func Max(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Max(args[0].Float(), args[1].Float())
	return float64(v)
}

func Min(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Min(args[0].Float(), args[1].Float())
	return float64(v)
}

func Mod(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Mod(args[0].Float(), args[1].Float())
	return float64(v)
}

// func Modf(f float64) (int float64, frac float64)

func Nextafter(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Nextafter(args[0].Float(), args[1].Float())
	return float64(v)
}

func Nextafter32(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Nextafter32(float32(args[0].Float()), float32(args[1].Float()))
	return float32(v)
}

func Pow(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Pow(args[0].Float(), args[1].Float())
	return float64(v)
}

func Pow10(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Pow10(args[0].Int())
	return float64(v)
}

func Remainder(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Remainder(args[0].Float(), args[1].Float())
	return float64(v)
}

func Round(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Round(args[0].Float())
	return float64(v)
}

func RoundToEven(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.RoundToEven(args[0].Float())
	return float64(v)
}

func Signbit(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Signbit(args[0].Float())
	return bool(v)
}

func Sin(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Sin(args[0].Float())
	return float64(v)
}

// func Sincos(x float64) (sin, cos float64)

func Sinh(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Sinh(args[0].Float())
	return float64(v)
}

func Sqrt(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Sqrt(args[0].Float())
	return float64(v)
}

func Tan(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Tan(args[0].Float())
	return float64(v)
}

func Tanh(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Tanh(args[0].Float())
	return float64(v)
}

func Trunc(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Trunc(args[0].Float())
	return float64(v)
}

func Y0(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Y0(args[0].Float())
	return float64(v)
}

func Y1(args []wasm.Value) any {
	if len(args) < 1 {
		return nil
	}
	v := math.Y1(args[0].Float())
	return float64(v)
}

func Yn(args []wasm.Value) any {
	if (len(args)) < 2 {
		return nil
	}
	v := math.Yn(args[0].Int(), args[1].Float())
	return float64(v)
}
