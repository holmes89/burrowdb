package lib

import (
	"errors"

	"github.com/holmes89/burrowdb"
)

type Evaluator struct {
	creator burrowdb.NodeCreator
}

func NewEvaluator(creator burrowdb.NodeCreator) *Evaluator {
	return &Evaluator{
		creator: creator,
	}
}

func (e *Evaluator) Eval(ex *Expr) error {
	switch ex.Command {
	case "CREATE":
		return e.creator.Create(&ex.Node)
	default:
		return errors.New("invalid command")
	}
}
