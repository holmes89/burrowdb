package lib

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"strings"
	"text/scanner"
	"unicode"
)

type Lexer struct {
	*scanner.Scanner
	buf bytes.Buffer
}

// New returns new lexer
func NewLexer(r io.Reader) *Lexer {
	var s scanner.Scanner
	s.Init(r)
	s.Mode &^= scanner.ScanChars | scanner.ScanRawStrings
	return &Lexer{
		Scanner: &s,
	}
}

func (l *Lexer) IsNextLParen() bool {

	tok := l.Peek()
	if isSpace(tok) {
		l.Next()
		return l.IsNextLParen()
	}
	return tok == '('
}

func (l *Lexer) IsNextRParen() bool {

	tok := l.Peek()
	if isSpace(tok) {
		l.Next()
		return l.IsNextRParen()
	}
	return tok == ')'
}

func (l *Lexer) IsNextColon() bool {
	tok := l.Peek()
	if isSpace(tok) {
		l.Next()
		return l.IsNextColon()
	}
	return tok == ':'
}

func (l *Lexer) IsNextRBrace() bool {
	tok := l.Peek()
	if isSpace(tok) {
		l.Next()
		return l.IsNextRBrace()
	}
	return tok == '}'
}

func (l *Lexer) IsNextLBrace() bool {
	tok := l.Peek()
	if isSpace(tok) {
		l.Next()
		return l.IsNextRBrace()
	}
	return tok == '{'
}

func (l *Lexer) ReadUntilRParen() []byte {
	var rs []rune
	for {
		tok := l.Peek()
		if isSpace(tok) {
			l.Next()
			continue
		}
		if tok == ')' {
			break
		}
		tok = l.Next()
		rs = append(rs, tok)
	}
	return []byte(string(rs))
}

func (l *Lexer) NextToken() *Token {
	for {
		r := l.Next()
		switch {
		case isSpace(r):
		case r == ';':
			l.skipToNewline()
		case r == scanner.EOF:
			return mkToken(tokenEOF, "EOF")
		case r == '\n':
			return mkToken(tokenNewline, "\n")
		case r == '(':
			return mkToken(tokenLpar, "(")
		case r == ')':
			return mkToken(tokenRpar, ")")
		case r == '{':
			return mkToken(tokenLbrace, "{")
		case r == '}':
			return mkToken(tokenRbrace, "}")
		case r == ':':
			return mkToken(tokenColon, ":")
		case isNumber(r):
			return l.number(r)
		case r == '"':
			return mkToken(tokenQuote, "\"")
		case r == '_' || unicode.IsLetter(r):
			return l.alphanum(tokenString, r)
		default:
			return mkToken(tokenChar, string(r))
		}
	}
}

var atoms = make(map[string]*Token)

var zero big.Int

func mkToken(typ TokenType, text string) *Token {
	if typ == tokenNumber {
		var z big.Int
		num, ok := z.SetString(text, 0)
		if !ok {
			errorf("bad number syntax: %s", text)
		}
		return number(num)
	}

	tok := atoms[strings.ToLower(text)]
	if tok == nil {
		tok = &Token{typ, text, &zero}
		atoms[text] = tok
	}
	return tok
}

func number(num *big.Int) *Token {
	return &Token{tokenNumber, "", num}
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

func isAlphanum(r rune) bool {
	return isNumber(r) || unicode.IsLetter(r)
}

func isNotDoubleQuote(r rune) bool { //doesn't support escaping in string
	return r != '"'
}

func (l *Lexer) accum(r rune, valid func(rune) bool) {
	l.buf.Reset()
	for {
		l.buf.WriteRune(r)
		r = l.Peek()
		if r == EofRune {
			return
		}
		if !valid(r) {
			return
		}
		r = l.Next()
	}
}

func (l *Lexer) number(r rune) *Token {
	// Integer only for now.
	l.accum(r, isNumber)
	l.endToken()
	return mkToken(tokenNumber, l.buf.String())
}

func (l *Lexer) alphanum(typ TokenType, r rune) *Token {
	// TODO: ASCII only for now.
	l.accum(r, isAlphanum)
	l.endToken()
	return mkToken(typ, l.buf.String())
}

func (l *Lexer) str(r rune) *Token {
	l.accum(r, isNotDoubleQuote)
	l.endToken()
	l.buf.WriteRune(l.Next()) //add end quote
	return mkToken(tokenString, l.buf.String())
}

// endToken guarantees that the following rune separates this token from the next.
func (l *Lexer) endToken() {
	if r := l.Peek(); isAlphanum(r) || !isSpace(r) && r != '(' && r != ')' && r != '{' && r != '}' && r != ':' && r != EofRune {
		errorf("invalid token after %s", l.String())
	}
}

func (l *Lexer) skipToNewline() {
	for {
		if r := l.Next(); r == 10 || r == 13 || r == scanner.EOF { //new lines and carrage returns
			break
		}
	}
}

func errorf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

var (
	tokenCreate = mkToken(tokenConst, "create")
	tokenMatch  = mkToken(tokenConst, "match")
)
