package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type interpreter struct {
	stack  []int
	ptr    int
	writer io.Writer
}

func (i *interpreter) incVal(val int) { i.stack[i.ptr] += val }
func (i *interpreter) modVal(val int) { i.stack[i.ptr] %= val }
func (i *interpreter) divVal(val int) { i.stack[i.ptr] /= val }
func (i *interpreter) peekVal() int   { return i.stack[i.ptr] }

func (i *interpreter) incPtr(v int) {
	i.ptr += v
	if i.ptr >= len(i.stack) {
		panic("stack overflow")
	}
}

func (i *interpreter) decPtr(v int) {
	i.ptr -= v
	if i.ptr < 0 {
		panic("stack underflow")
	}
}

func newInterpreter() *interpreter {
	return &interpreter{
		stack:  make([]int, 1024),
		writer: os.Stdout,
	}
}

func mustParseNumberRune(str []rune) []rune {
	number := parseNumberRune(str)
	if len(number) == 0 {
		panic("missing expected number")
	}
	return number
}

func parseNumberRune(str []rune) []rune {
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return str[0:i]
		}
	}
	return str[:]
}

func toInt(str []rune) int {
	if val, err := strconv.Atoi(string(str)); err != nil {
		panic(err)
	} else {
		return val
	}
}

func balancedRune(str []rune) []rune {
	depth := 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return str[:i]
			}
		}
	}
	panic("expected balanced expression")
}

func parsePointers(str []rune) (count int, depth int) {
LOOP:
	for {
		switch str[count] {
		case ' ', '\t', '\n':
			// pass
		case '<':
			depth--
		case '>':
			depth++
		default:
			break LOOP
		}
		count++
	}
	return
}

func mustParseConditional(str []rune) (l int, t []rune, f []rune) {
	if len(str) < 1 {
		panic(fmt.Errorf("attempt to parse empty conditional expression"))
	}
	if str[0] != '?' {
		panic(fmt.Errorf("missing expected ? in conditional"))
	}
	l++

TRUE:
	for l < len(str) {
		switch str[l] {
		case ',':
			l++
			break TRUE
		default:
			t = append(t, str[l])
			l++
		}
	}

FALSE:
	for l < len(str) {
		switch str[l] {
		case '.':
			l++
			break FALSE
		default:
			f = append(f, str[l])
			l++
		}
	}

	return
}

func (intrptr *interpreter) exec(str []rune) {
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number := parseNumberRune(str[i:])
			i += len(number) - 1
			intrptr.incVal(toInt(number))
		case '/':
			switch {
			case len(str) == i-1:
				panic(fmt.Errorf("missing expected addition term"))
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				intrptr.incPtr(d)
				v := intrptr.peekVal()
				intrptr.incPtr(-d)
				intrptr.divVal(v)
				intrptr.incPtr(d)
			default:
				panic(fmt.Errorf("unexpected term %s", string(str[i+1])))
			}
		case '%':
			switch {
			case len(str) == i-1:
				panic(fmt.Errorf("missing expected addition term"))
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				intrptr.incPtr(d)
				v := intrptr.peekVal()
				intrptr.incPtr(-d)
				intrptr.modVal(v)
				intrptr.incPtr(d)
			default:
				panic(fmt.Errorf("unexpected term %s", string(str[i+1])))
			}
		case '+':
			switch {
			case len(str) == i-1:
				panic(fmt.Errorf("missing expected addition term"))
			case str[i+1] == '+':
				i += 1
				intrptr.incVal(1)
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				intrptr.incPtr(d)
				v := intrptr.peekVal()
				intrptr.incPtr(-d)
				intrptr.incVal(v)
				intrptr.incPtr(d)
			default:
				number := mustParseNumberRune(str[i+1:])
				i += len(number)
				intrptr.incVal(toInt(number))
			}
		case '-':
			switch {
			case len(str) == i-1:
				panic(fmt.Errorf("missing expected addition term"))
			case str[i+1] == '-':
				i += 1
				intrptr.incVal(-1)
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				intrptr.incPtr(d)
				v := intrptr.peekVal()
				intrptr.incPtr(-d)
				intrptr.incVal(-v)
				intrptr.incPtr(d)
			default:
				number := mustParseNumberRune(str[i+1:])
				i += len(number)
				intrptr.incVal(-toInt(number))
			}
		case '(':
			balanced := balancedRune(str[i:])
			intrptr.exec(balanced[1:])
			if intrptr.peekVal() == 0 {
				i += len(balanced) - 1
			} else {
				i--
			}
		case '?':
			l, t, f := mustParseConditional(str[i:])
			if intrptr.peekVal() == 0 {
				intrptr.exec(f)
			} else {
				intrptr.exec(t)
			}
			i += l - 1
		case ')':
			continue
		case '.':
			continue
		case '>':
			intrptr.incPtr(1)
		case '<':
			intrptr.decPtr(1)
		case 'π':
			fmt.Fprintf(intrptr.writer, "%c", intrptr.peekVal())
		case ' ', '\n', '\t':
			continue
		case '=':
			l := bytes.IndexRune([]byte(string(str[i:])), '\n')
			if l > -1 {
				i += l
			} else {
				i += len(str[i+1:])
			}
		default:
			panic(fmt.Errorf("unexpected char %s", string(str[i])))
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		panic(fmt.Errorf("usage numberscript <filename>"))
	}
	filename := flag.Args()[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read file:", err)
		os.Exit(1)
	}

	newInterpreter().exec([]rune(string(content)))
}
