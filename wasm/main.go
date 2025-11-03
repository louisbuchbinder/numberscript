package wasm

type Value interface {
	Bool() bool
	Float() float64
	Int() int
	String() string
}
