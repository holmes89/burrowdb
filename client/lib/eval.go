package lib

import (
	"errors"

	"github.com/holmes89/burrowdb"
)

type Evaluator struct {
	repo *burrowdb.Repository
}

func NewEvaluator(repo *burrowdb.Repository) *Evaluator {
	return &Evaluator{
		repo: repo,
	}
}

func (e *Evaluator) Eval(ex *Expr) error {
	switch ex.Command {
	case "CREATE":
		return e.repo.Create(&ex.Node)
	default:
		return errors.New("invalid command")
	}
}
