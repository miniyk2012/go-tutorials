package resolver

import (
	"goask/core/adapter"
)

type Query struct {
	Data adapter.Data
}


func (q *Query) Questions(args struct{ Search *string }) ([]Question, error) {
	questions, err := q.Data.Questions(args.Search)
	return QuestionAll(questions, q.Data), err
}
