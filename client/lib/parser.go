package lib

import (
	"encoding/json"
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
	return fmt.Sprintf("Command: %s, Node: {id: %s, label: %+v, properties: %+v}", e.Command, e.Node.ID, e.Node.Labels, e.Node.Properties)
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
	case tokenMatch:
		return p.parseMatch()
	default:
		return nil, errors.New("invalid command")
	}
}

func (p *Parser) parseMatch() (*Expr, error) {

	exp := &Expr{
		Command: "MATCH",
	}

	if t := p.lex.NextToken().typ; t != tokenLpar {
		return nil, errors.New("invalid syntax")
	}

	// First token is a variable to use
	token := p.lex.NextToken()
	if token.typ != tokenString {
		return nil, errors.New("invalid syntax")
	}

	return exp, nil
}

func (p *Parser) parseCreate() (*Expr, error) {

	exp := &Expr{
		Command: "CREATE",
	}

	if t := p.lex.NextToken().typ; t != tokenLpar {
		return nil, errors.New("invalid syntax")
	}

	// First token is the node id
	token := p.lex.NextToken()
	if token.typ != tokenString {
		return nil, errors.New("invalid syntax")
	}

	exp.Node.ID = token.text

	token = p.lex.NextToken()
	switch token.typ {
	case tokenRpar:
		return exp, nil
	case tokenColon:
		l, err := p.parseLabels()
		if err != nil {
			return nil, err
		}
		exp.Node.Labels = l
	case tokenLbrace:
		o, err := p.parseProperties()
		if err != nil {
			return nil, err
		}
		exp.Node.Properties = o
	default:
		return nil, errors.New("invalid syntax")
	}

	token = p.lex.NextToken()
	switch token.typ {
	case tokenRpar:
		return exp, nil
	case tokenLbrace:
		o, err := p.parseProperties()
		if err != nil {
			return nil, err
		}
		exp.Node.Properties = o
	default:
		return nil, errors.New("invalid syntax")
	}

	token = p.lex.NextToken()
	if token.typ != tokenRpar {
		return nil, errors.New("invalid syntax")
	}

	return exp, nil
}

func (p *Parser) parseLabels() (labels []burrowdb.Label, err error) {
	token := p.lex.NextToken()
	if token.typ != tokenString {
		return nil, errors.New("invalid syntax label must be string")
	}
	labels = append(labels, (burrowdb.Label)(token.text))

	for {
		if p.lex.IsNextRParen() || p.lex.IsNextLBrace() {
			break
		}
		token := p.lex.NextToken()

		if token.typ != tokenColon {
			return nil, errors.New("invalid syntax label must start with a colon")
		}

		token = p.lex.NextToken()
		if token.typ != tokenString {
			return nil, errors.New("invalid syntax label must be string")
		}
		labels = append(labels, (burrowdb.Label)(token.text))

	}

	return labels, nil
}

func (p *Parser) parseProperties() (m map[string]interface{}, err error) {
	b := p.lex.ReadUntilRParen()
	b = append([]byte{'{'}, b...)
	if err := json.Unmarshal(b, &m); err != nil {
		return m, errors.New("invalid syntax for properties")
	}
	return m, nil
}
