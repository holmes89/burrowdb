package lib

import (
	"fmt"
	"math/big"
	"strings"
)

//go:generate stringer -type TokenType -trimprefix Token
type TokenType int

type Token struct {
	typ  TokenType
	text string
	num  *big.Int
}

func (t *Token) IsEOF() bool {
	return t.typ == tokenEOF
}

const (
	tokenError TokenType = iota
	tokenEOF
	tokenConst
	tokenNumber
	tokenLpar
	tokenRpar
	tokenLbrace
	tokenRbrace
	tokenColon
	tokenString
	tokenChar
	tokenQuote
	tokenNewline
)

const EofRune rune = -1

func (t Token) String() string {
	if t.typ == tokenNumber {
		return fmt.Sprint(t.typ, t.num)
	}
	return fmt.Sprint(t.typ, " ", t.text)
}

func (t Token) buildString(b *strings.Builder) {
	if t.typ == tokenNumber {
		b.WriteString(fmt.Sprint(t.num))
	} else {
		b.WriteString(t.text)
	}
}
