package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Ptr[A any](a A) *A { return &a }

func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[A any](v A, err error) A {
	if err != nil {
		panic(err)
	} else {
		return v
	}
}

func Map[A, B any](in []A, fn func(int, A) B) []B {
	result := make([]B, len(in))
	for i, v := range in {
		result[i] = fn(i, v)
	}
	return result
}

func MapOrError[A, B any](in []A, fn func(int, A) (B, error)) ([]B, error) {
	result := make([]B, len(in))
	for i, v := range in {
		if v0, err := fn(i, v); err != nil {
			return nil, err
		} else {
			result[i] = v0
		}
	}
	return result, nil
}

func Filter[A any](in []A, fn func(int, A) bool) []A {
	result := []A{}
	for i, v := range in {
		if fn(i, v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[From any, To any](in []From, fn func(int, To, From) To, init To) To {
	if len(in) == 0 {
		return init
	}
	result := init
	for i, v := range in {
		result = fn(i, result, v)
	}
	return result
}

func Concat[A any](ins ...[]A) []A {
	result := make([]A, 0)
	for _, in := range ins {
		result = append(result, in...)
	}
	return result
}

func Flatten[A any](in [][]A) []A {
	return Concat(in...)
}

type Entry[Key comparable, Value any] struct {
	Key   Key
	Value Value
}

func Entries[Key comparable, Value any](object map[Key]Value) []Entry[Key, Value] {
	out := make([]Entry[Key, Value], len(object))
	for key, value := range object {
		out = append(out, Entry[Key, Value]{Key: key, Value: value})
	}
	return out
}

func Keys[Key comparable, Value any](object map[Key]Value) []Key {
	return Map(Entries(object), func(_ int, entry Entry[Key, Value]) Key {
		return entry.Key
	})
}

func Values[Key comparable, Value any](object map[Key]Value) []Value {
	return Map(Entries(object), func(_ int, entry Entry[Key, Value]) Value {
		return entry.Value
	})
}

func Sha256Hex(b []byte) string {
	hsh := sha256.Sum256(b)
	return hex.EncodeToString(hsh[:])
}

func Sha256HexOfFile(f io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

type Template interface {
	Execute(io.Writer, any) error
}

func ExecuteTemplate(t Template, data any) ([]byte, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
