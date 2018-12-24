package fakeadapter

import (
	"goask/core/adapter"
	"goask/core/entity"

	"github.com/pkg/errors"
)

type Data struct {
	questions []entity.Question
}

var _ adapter.Data = &Data{}

func (d *Data) Questions() ([]entity.Question, error) {
	return d.questions, nil
}

func (d *Data) CreateQuestion(p entity.Post) (entity.Question, error) {
	p.ID = int64(len(d.questions) + 1)
	d.questions = append(d.questions, entity.Question{Post: p})
	return d.questions[len(d.questions)-1], nil
}

func (d *Data) UpdateQuestion(p entity.PostUpdate) (entity.Question, error) {
	if p.ID == 0 {
		return entity.Question{}, errors.New("ID can not be 0 nor absent")
	}
	for i, q := range d.questions {
		if q.ID == p.ID {
			if p.Content != nil {
				q.Content = *p.Content
			}
			if p.Title != nil {
				q.Title = *p.Title
			}
			d.questions[i] = q
			return q, nil
		}
	}
	return entity.Question{}, errors.WithStack(&adapter.ErrQuestionNotFound{ID: p.ID})
}
