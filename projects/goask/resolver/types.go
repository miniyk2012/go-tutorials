package resolver

import (
	"goask/core/entity"
)

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

// Answer is the GraphQL resolver for Answer type.
type Answer struct {
	entity entity.Answer
}

func (a Answer) ID() int32 {
	return int32(a.entity.ID)
}

func (a Answer) Content() string {
	return a.entity.Content
}

func (a Answer) Question() Question {
	return Question{}
}

