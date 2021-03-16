package lib

import (
	"errors"
	"fmt"

	"github.com/holmes89/burrowdb"
)

type Parser struct {
	lex *Lexer
}

func NewParser(lex *Lexer) *Parser {
	return &Parser{lex}
}

type Expr struct {
	Command string
	Node    burrowdb.Node
}

func (e *Expr) String() string {
	return fmt.Sprintf("Command: %s, Node: {id: %s, label: %s}", e.Command, e.Node.ID, e.Node.Label)
}

var ErrEndOfInputStream = errors.New("end of input stream")

func (p *Parser) Parse() (*Expr, error) {
	token := p.lex.NextToken()
	switch token.typ {
	case tokenEOF:
		return nil, ErrEndOfInputStream
	case tokenConst:
		return p.handleConst(token)
	default:
		return nil, errors.New("unknown token")
	}
}

func (p *Parser) handleConst(token *Token) (*Expr, error) {
	switch token {
	case tokenCreate:
		return p.parseCreate()
	default:
		return nil, errors.New("invalid command")
	}
}

func (p *Parser) parseCreate() (*Expr, error) {

	exp := &Expr{
		Command: "CREATE",
	}

	// TODO multinode creation
	// TODO multilabel
	if t := p.lex.NextToken().typ; t != tokenLpar {
		return nil, errors.New("invalid syntax")
	}

	// First token is the node id
	token := p.lex.NextToken()
	if token.typ != tokenString {
		return nil, errors.New("invalid syntax")
	}

	exp.Node.ID = token.text

	// next can be rparen or label or properties
	for {
		token = p.lex.NextToken()
		switch token.typ {
		case tokenRpar:
			return exp, nil
		case tokenColon:
			l, err := p.parseLabel()
			if err != nil {
				return nil, err
			}
			exp.Node.Label = l
		default:
			return nil, errors.New("invalid syntax")
		}
	}
}

func (p *Parser) parseLabel() (burrowdb.Label, error) {
	token := p.lex.NextToken()
	if token.typ != tokenString {
		return "", errors.New("invalid syntax label must be string")
	}
	return (burrowdb.Label)(token.text), nil
}
