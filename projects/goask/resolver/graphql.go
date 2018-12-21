package resolver

import (
	"goask/core/entity"
)

type Query struct {
	//questions interactor.Questions
}

func (q *Query) Questions() ([]Question, error) {
	//questions, err := q.questions.GetAll()
	//if err != nil {
	//	return nil, err
	//}

	return (*Question).All(nil, nil), nil
}

type Question struct {
	entity entity.Question
}

func (q *Question) All(questions []entity.Question) []Question {
	ret := make([]Question, len(questions))
	for i, question := range questions {
		ret[i] = Question{
			entity: question,
		}
	}
	return ret
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
