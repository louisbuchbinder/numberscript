package wasm

import (
	"io"
	"io/fs"
)

type OpfsFile interface {
	io.WriteCloser
	fs.File
}

type Value interface {
	Bool() bool
	Float() float64
	Int() int
	String() string
	Array() []any
	Uint8Array() []uint8
	IsArray() bool
	IsUint8Array() bool
	Invoke(...any) Value
	Reject(error)
	File() (fs.File, error)
	FileWriter() (OpfsFile, error)
	FS() (fs.FS, error)
	Request() (JsRequest, error)
}

type JsRequest interface {
	Method() string
	Url() JsUrl
}

type JsUrl interface {
	Hash() string
	Host() string
	Hostname() string
	Href() string
	Origin() string
	Password() string
	Pathname() string
	Port() string
	Protocol() string
	Search() string
	GetSearchParam(string) *string
	Username() string
}
