package wasm

type Value interface {
	Bool() bool
	Float() float64
	Int() int
	String() string
	Array() []any
	Uint8Array() []uint8
	IsArray() bool
	IsUint8Array() bool
}
