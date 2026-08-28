package _15_interpreter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Parser 把表达式字符串解析成一棵 AST。
type Parser struct {
	tokens []string
	pos    int
}

func NewParser(input string) *Parser {
	return &Parser{tokens: tokenize(input)}
}

// tokenize 把输入拆成数字与运算符。
func tokenize(input string) []string {
	var tokens []string
	var num strings.Builder
	flush := func() {
		if num.Len() > 0 {
			tokens = append(tokens, num.String())
			num.Reset()
		}
	}
	for _, r := range input {
		switch {
		case unicode.IsDigit(r):
			num.WriteRune(r)
		case r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')':
			flush()
			tokens = append(tokens, string(r))
		}
	}
	flush()
	return tokens
}

func (p *Parser) Parse() (Expression, error) {
	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	if p.pos != len(p.tokens) {
		return nil, fmt.Errorf("unexpected token %q", p.tokens[p.pos])
	}
	return expr, nil
}

// parseExpr 处理加减（优先级最低）：expr = term { ('+'|'-') term }。
func (p *Parser) parseExpr() (Expression, error) {
	left, err := p.parseTerm()
	if err != nil {
		return nil, err
	}
	for p.peek() == "+" || p.peek() == "-" {
		op := p.next()
		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}
		if op == "+" {
			left = &Add{left: left, right: right}
		} else {
			left = &Sub{left: left, right: right}
		}
	}
	return left, nil
}

// parseTerm 处理乘除：term = factor { ('*'|'/') factor }。
func (p *Parser) parseTerm() (Expression, error) {
	left, err := p.parseFactor()
	if err != nil {
		return nil, err
	}
	for p.peek() == "*" || p.peek() == "/" {
		op := p.next()
		right, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		if op == "*" {
			left = &Mul{left: left, right: right}
		} else {
			left = &Div{left: left, right: right}
		}
	}
	return left, nil
}

// parseFactor 处理数字或括号：factor = number | '(' expr ')'。
func (p *Parser) parseFactor() (Expression, error) {
	tok := p.peek()
	if tok == "" {
		return nil, fmt.Errorf("unexpected end of input")
	}
	if tok == "(" {
		p.next()
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		if p.next() != ")" {
			return nil, fmt.Errorf("missing )")
		}
		return expr, nil
	}
	p.next()
	n, err := strconv.Atoi(tok)
	if err != nil {
		return nil, fmt.Errorf("invalid number %q", tok)
	}
	return &Number{value: n}, nil
}

func (p *Parser) peek() string {
	if p.pos >= len(p.tokens) {
		return ""
	}
	return p.tokens[p.pos]
}

func (p *Parser) next() string {
	tok := p.peek()
	if tok != "" {
		p.pos++
	}
	return tok
}

// Eval 便捷入口：解析并求值。
func Eval(input string) (int, error) {
	expr, err := NewParser(input).Parse()
	if err != nil {
		return 0, err
	}
	return expr.Interpret(), nil
}
