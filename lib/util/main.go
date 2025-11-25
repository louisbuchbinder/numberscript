package util

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net"
	"os"
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

const (
	ThirtyTwoKB   = 1024 * 32
	OneMB         = 1048576
	ThirtyTwoMB   = OneMB * 32
	FiveHundredMB = OneMB * 500
	FiveGB        = 1073741824 * 5
)

func CopyFlexBuffer(dst io.Writer, src io.Reader, n int64) (int64, error) {
	var buf []byte
	switch {
	case n > FiveGB:
		buf = make([]byte, ThirtyTwoMB)
	case n > FiveHundredMB:
		buf = make([]byte, OneMB)
	default:
		buf = make([]byte, ThirtyTwoKB)
	}
	return io.CopyBuffer(dst, src, buf)
}

type CleanupFunctionWrapper struct {
	fns []func() error
}

func NewCleanupFunctionWrapper(fns ...func() error) *CleanupFunctionWrapper {
	return &CleanupFunctionWrapper{fns: fns}
}

func (w *CleanupFunctionWrapper) Cleanup() func() error {
	return func() error {
		var firstErr error
		for _, fn := range w.fns {
			if err := fn(); err != nil && firstErr == nil {
				firstErr = err
			}
		}
		return firstErr
	}
}

func (w *CleanupFunctionWrapper) Add(fn func() error) {
	w.fns = append(w.fns, fn)
}

func (w *CleanupFunctionWrapper) AddCancel(fn context.CancelFunc) {
	w.fns = append(w.fns, func() error {
		fn()
		return nil
	})
}

func FreePort() (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

func MkdirTemp(dir string, pattern string) (string, func() error, error) {
	tempdir, err := os.MkdirTemp(dir, pattern)
	cleanup := func() error {
		return os.RemoveAll(tempdir)
	}
	return tempdir, cleanup, err
}
