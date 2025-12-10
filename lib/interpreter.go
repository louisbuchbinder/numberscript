package lib

import (
	"bytes"
	"fmt"
	"strconv"
)

type Interpreter interface {
	Exec(str []rune) ([]byte, error)
}

func (intrptr *interpreter) Exec(str []rune) ([]byte, error) {
	if err := intrptr.exec(str); err != nil {
		return nil, err
	}
	return intrptr.buffer.Bytes(), nil
}

type interpreter struct {
	stack  []int
	ptr    int
	buffer *bytes.Buffer
}

func (i *interpreter) incVal(val int) { i.stack[i.ptr] += val }
func (i *interpreter) modVal(val int) { i.stack[i.ptr] %= val }
func (i *interpreter) divVal(val int) { i.stack[i.ptr] /= val }
func (i *interpreter) peekVal() int   { return i.stack[i.ptr] }

func (i *interpreter) incPtr(v int) error {
	i.ptr += v
	if i.ptr >= len(i.stack) {
		return fmt.Errorf("stack overflow")
	}
	return nil
}

func (i *interpreter) decPtr(v int) error {
	i.ptr -= v
	if i.ptr < 0 {
		return fmt.Errorf("stack underflow")
	}
	return nil
}

func NewInterpreter() *interpreter {
	return &interpreter{
		stack:  make([]int, 1024),
		buffer: new(bytes.Buffer),
	}
}

func mustParseNumberRune(str []rune) ([]rune, error) {
	number := parseNumberRune(str)
	if len(number) == 0 {
		return nil, fmt.Errorf("missing expected number")
	}
	return number, nil
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

func toInt(str []rune) (int, error) {
	if val, err := strconv.Atoi(string(str)); err != nil {
		return 0, err
	} else {
		return val, nil
	}
}

func balancedRune(str []rune) ([]rune, error) {
	depth := 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return str[:i], nil
			}
		}
	}
	return nil, fmt.Errorf("expected balanced expression")
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

func mustParseConditional(str []rune) (l int, t []rune, f []rune, err error) {
	if len(str) < 1 {
		return 0, nil, nil, fmt.Errorf("attempt to parse empty conditional expression")
	}
	if str[0] != '?' {
		return 0, nil, nil, fmt.Errorf("missing expected ? in conditional")
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

func (intrptr *interpreter) exec(str []rune) error {
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number := parseNumberRune(str[i:])
			i += len(number) - 1
			v, err := toInt(number)
			if err != nil {
				return err
			}
			intrptr.incVal(v)
		case '/':
			switch {
			case len(str) == i-1:
				return fmt.Errorf("missing expected addition term")
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
				v := intrptr.peekVal()
				if err := intrptr.incPtr(-d); err != nil {
					return err
				}
				intrptr.divVal(v)
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unexpected term %s", string(str[i+1]))
			}
		case '%':
			switch {
			case len(str) == i-1:
				return fmt.Errorf("missing expected addition term")
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
				v := intrptr.peekVal()
				if err := intrptr.incPtr(-d); err != nil {
					return err
				}
				intrptr.modVal(v)
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unexpected term %s", string(str[i+1]))
			}
		case '+':
			switch {
			case len(str) == i-1:
				return fmt.Errorf("missing expected addition term")
			case str[i+1] == '+':
				i += 1
				intrptr.incVal(1)
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
				v := intrptr.peekVal()
				if err := intrptr.incPtr(-d); err != nil {
					return err
				}
				intrptr.incVal(v)
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
			default:
				number, err := mustParseNumberRune(str[i+1:])
				if err != nil {
					return err
				}
				i += len(number)
				v, err := toInt(number)
				if err != nil {
					return err
				}
				intrptr.incVal(v)
			}
		case '-':
			switch {
			case len(str) == i-1:
				return fmt.Errorf("missing expected addition term")
			case str[i+1] == '-':
				i += 1
				intrptr.incVal(-1)
			case str[i+1] == '<' || str[i+1] == '>':
				j, d := parsePointers(str[i+1:])
				i += j
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
				v := intrptr.peekVal()
				if err := intrptr.incPtr(-d); err != nil {
					return err
				}
				intrptr.incVal(-v)
				if err := intrptr.incPtr(d); err != nil {
					return err
				}
			default:
				number, err := mustParseNumberRune(str[i+1:])
				if err != nil {
					return err
				}
				i += len(number)
				v, err := toInt(number)
				if err != nil {
					return err
				}
				intrptr.incVal(-1 * v)
			}
		case '(':
			balanced, err := balancedRune(str[i:])
			if err != nil {
				return err
			}
			intrptr.Exec(balanced[1:])
			if intrptr.peekVal() == 0 {
				i += len(balanced) - 1
			} else {
				i--
			}
		case '?':
			l, t, f, err := mustParseConditional(str[i:])
			if err != nil {
				return err
			}
			if intrptr.peekVal() == 0 {
				intrptr.Exec(f)
			} else {
				intrptr.Exec(t)
			}
			i += l - 1
		case ')':
			continue
		case '.':
			continue
		case '>':
			if err := intrptr.incPtr(1); err != nil {
				return err
			}
		case '<':
			if err := intrptr.decPtr(1); err != nil {
				return err
			}
		case 'π':
			fmt.Fprintf(intrptr.buffer, "%c", intrptr.peekVal())
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
			return fmt.Errorf("unexpected char %s", string(str[i]))
		}
	}
	return nil
}
