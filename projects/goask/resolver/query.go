package resolver

import (
	"github.com/pkg/errors"
	"goask/core/adapter"
)

type Query struct {
	Data adapter.Data
}

func (q *Query) check() error {
	if q.Data == nil {
		return errors.New("resolver.Query.Data is not initialized")
	}
	return nil
}

func (q *Query) Questions(args struct{Search *string}) ([]Question, error) {
	if err := q.check(); err != nil {
		return nil, err
	}
	questions, err := q.Data.Questions(args.Search)
	return (*Question).all(nil, questions), err
}
