package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode"
)

type tokenType int

const (
	tEOF tokenType = iota
	tNumber
	tPlus
	tMinus
	tMul
	tDiv
	tLParen
	tRParen
)

type token struct {
	typ tokenType
	val float64 // for numbers
}

type lexer struct {
	s   string
	i   int
	n   int
	err error
}

func newLexer(input string) *lexer {
	return &lexer{s: input, n: len(input)}
}

func (l *lexer) skipSpaces() {
	for l.i < l.n && unicode.IsSpace(rune(l.s[l.i])) {
		l.i++
	}
}

func (l *lexer) nextToken() token {
	l.skipSpaces()
	if l.i >= l.n {
		return token{typ: tEOF}
	}

	ch := l.s[l.i]

	switch ch {
	case '+':
		l.i++
		return token{typ: tPlus}
	case '-':
		l.i++
		return token{typ: tMinus}
	case '*':
		l.i++
		return token{typ: tMul}
	case '/':
		l.i++
		return token{typ: tDiv}
	case '(':
		l.i++
		return token{typ: tLParen}
	case ')':
		l.i++
		return token{typ: tRParen}
	default:
		// number: digits with optional dot, supports ".5" and "2." forms
		if ch == '.' || (ch >= '0' && ch <= '9') {
			start := l.i
			dotSeen := false

			if ch == '.' {
				dotSeen = true
				l.i++
			}

			for l.i < l.n {
				c := l.s[l.i]
				if c >= '0' && c <= '9' {
					l.i++
					continue
				}
				if c == '.' {
					if dotSeen {
						l.err = fmt.Errorf("invalid number: multiple dots near position %d", l.i)
						return token{typ: tEOF}
					}
					dotSeen = true
					l.i++
					continue
				}
				break
			}

			numStr := l.s[start:l.i]
			// Reject "." alone
			if numStr == "." {
				l.err = fmt.Errorf("invalid number: '.' at position %d", start)
				return token{typ: tEOF}
			}

			var v float64
			_, err := fmt.Sscanf(numStr, "%f", &v)
			if err != nil {
				l.err = fmt.Errorf("invalid number '%s' at position %d", numStr, start)
				return token{typ: tEOF}
			}
			return token{typ: tNumber, val: v}
		}

		l.err = fmt.Errorf("unexpected character '%c' at position %d", ch, l.i)
		return token{typ: tEOF}
	}
}

type parser struct {
	l   *lexer
	cur token
}

func newParser(input string) *parser {
	lex := newLexer(input)
	p := &parser{l: lex}
	p.cur = p.l.nextToken()
	return p
}

func (p *parser) eat(tt tokenType) error {
	if p.l.err != nil {
		return p.l.err
	}
	if p.cur.typ != tt {
		return fmt.Errorf("unexpected token: got %v, want %v", p.cur.typ, tt)
	}
	p.cur = p.l.nextToken()
	return p.l.err
}

// Grammar (classic precedence):
// expr   := term (('+'|'-') term)*
// term   := unary (('*'|'/') unary)*
// unary  := ('+'|'-') unary | primary
// primary:= number | '(' expr ')'

func (p *parser) parse() (float64, error) {
	res, err := p.expr()
	if err != nil {
		return 0, err
	}
	if p.l.err != nil {
		return 0, p.l.err
	}
	if p.cur.typ != tEOF {
		return 0, fmt.Errorf("unexpected trailing input")
	}
	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, fmt.Errorf("result is not a finite number")
	}
	return res, nil
}

func (p *parser) expr() (float64, error) {
	left, err := p.term()
	if err != nil {
		return 0, err
	}

	for p.cur.typ == tPlus || p.cur.typ == tMinus {
		op := p.cur.typ
		if err := p.eat(op); err != nil {
			return 0, err
		}
		right, err := p.term()
		if err != nil {
			return 0, err
		}
		if op == tPlus {
			left += right
		} else {
			left -= right
		}
	}
	return left, nil
}

func (p *parser) term() (float64, error) {
	left, err := p.unary()
	if err != nil {
		return 0, err
	}

	for p.cur.typ == tMul || p.cur.typ == tDiv {
		op := p.cur.typ
		if err := p.eat(op); err != nil {
			return 0, err
		}
		right, err := p.unary()
		if err != nil {
			return 0, err
		}
		if op == tMul {
			left *= right
		} else {
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			left /= right
		}
	}
	return left, nil
}

func (p *parser) unary() (float64, error) {
	if p.cur.typ == tPlus {
		if err := p.eat(tPlus); err != nil {
			return 0, err
		}
		return p.unary()
	}
	if p.cur.typ == tMinus {
		if err := p.eat(tMinus); err != nil {
			return 0, err
		}
		v, err := p.unary()
		if err != nil {
			return 0, err
		}
		return -v, nil
	}
	return p.primary()
}

func (p *parser) primary() (float64, error) {
	switch p.cur.typ {
	case tNumber:
		v := p.cur.val
		if err := p.eat(tNumber); err != nil {
			return 0, err
		}
		return v, nil
	case tLParen:
		if err := p.eat(tLParen); err != nil {
			return 0, err
		}
		v, err := p.expr()
		if err != nil {
			return 0, err
		}
		if p.cur.typ != tRParen {
			return 0, fmt.Errorf("missing closing ')'")
		}
		if err := p.eat(tRParen); err != nil {
			return 0, err
		}
		return v, nil
	default:
		return 0, fmt.Errorf("unexpected token in expression")
	}
}

func main() {
	var input string

	if len(os.Args) > 1 {
		// Если передали аргумент — берём его
		input = os.Args[1]
	} else {
		// Иначе читаем из STDIN
		data, err := io.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			fmt.Fprintln(os.Stderr, "read error:", err)
			os.Exit(1)
		}
		input = string(data)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Fprintln(os.Stderr, "empty input")
		os.Exit(1)
	}

	p := newParser(input)
	res, err := p.parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse/eval error:", err)
		os.Exit(1)
	}

	fmt.Printf("%g\n", res)
}
