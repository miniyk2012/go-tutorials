package resolver

import (
	"goask/core/adapter"
	"goask/core/entity"

	"github.com/pkg/errors"
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

func (q *Query) Questions() ([]Question, error) {
	if err := q.check(); err != nil {
		return nil, err
	}
	questions, err := q.Data.Questions()
	return (*Question).all(nil, questions), err
}

// Question is the GraphQL resolver for Question type.
type Question struct {
	entity entity.Question
}

func (q *Question) all(questions []entity.Question) []Question {
	ret := make([]Question, len(questions))
	for i, question := range questions {
		ret[i] = q.one(question)
	}
	return ret
}

func (q *Question) one(question entity.Question) Question {
	return Question{
		entity: question,
	}
}

func (q Question) ID() int32 {
	return int32(q.entity.ID)
}

func (q Question) Title() string {
	return string(q.entity.Title)
}

func (q Question) Content() string {
	return string(q.entity.Content)
}

//type Comment struct {
//
//}
//
//type Answer struct {
//
//}
