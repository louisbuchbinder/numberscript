package main

import (
	"embed"
	"testing"

	"github.com/louisbuchbinder/numberscript/lib"

	"github.com/stretchr/testify/assert"
)

var _ = new(embed.FS)

//go:embed test/basicStackExample.ns
var basicStackExample string

//go:embed test/basicStackExample.expect
var basicStackExampleExpect string

//go:embed test/hello.ns
var hello string

//go:embed test/hello.expect
var helloExpect string

//go:embed test/alphabet.ns
var alphabet string

//go:embed test/alphabet.expect
var alphabetExpect string

//go:embed test/printNumber.ns
var printNumber string

//go:embed test/printNumber.expect
var printNumberExpect string

//go:embed test/fibonacci.ns
var fibonacci string

//go:embed test/fibonacci.expect
var fibonacciExpect string

type testCase struct {
	name   string
	script string
	expect string
}

var testCases []testCase

func init() {
	testCases = []testCase{
		{
			name:   "print 0",
			script: "48π",
			expect: "0",
		},
		{
			name:   "print a",
			script: "97π",
			expect: "a",
		},
		{
			name:   "print A",
			script: "65π",
			expect: "A",
		},
		{
			name:   "basicStackExample",
			script: basicStackExample,
			expect: basicStackExampleExpect,
		},
		{
			name:   "alphabet",
			script: alphabet,
			expect: alphabetExpect,
		},
		{
			name:   "hello",
			script: hello,
			expect: helloExpect,
		},
		{
			name:   "printNumber",
			script: printNumber,
			expect: printNumberExpect,
		},
		{
			name:   "fibonacci",
			script: fibonacci,
			expect: fibonacciExpect,
		},
	}
}

func TestInterpreter(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			intrptr := lib.NewInterpreter()
			b, err := intrptr.Exec([]rune(tc.script))
			assert.Nil(t, err)
			assert.Equal(tt, tc.expect, string(b))
		})
	}
}
