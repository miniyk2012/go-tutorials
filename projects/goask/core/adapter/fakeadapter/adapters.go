package fakeadapter

import (
	"goask/core/adapter"
	"goask/core/entity"
	"strings"

	"github.com/pkg/errors"
)

type Data struct {
	questions []entity.Question
	answers   Answers
}

var _ adapter.Data = &Data{}

func (d *Data) Questions(search *string) ([]entity.Question, error) {
	if search == nil {
		return d.questions, nil
	}
	ret := make([]entity.Question, 0)
	for _, q := range d.questions {
		if match(q.Content, *search) {
			ret = append(ret, q)
		}
	}
	return ret, nil
}

func (d *Data) CreateQuestion(q entity.Question) (entity.Question, error) {
	q.ID = int64(len(d.questions) + 1)
	d.questions = append(d.questions, q)
	return d.questions[len(d.questions)-1], nil
}

func (d *Data) UpdateQuestion(p entity.QuestionUpdate) (entity.Question, error) {
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

func (d *Data) CreateAnswer(answerCreation entity.AnswerCreation) (entity.Answer, error) {
	for _, q := range d.questions {
		if q.ID == answerCreation.QuestionID {
			answer := d.answers.Add(answerCreation)
			q.Answers = append(q.Answers, answer)
			return answer, nil
		}
	}
	return entity.Answer{}, errors.WithStack(&adapter.ErrQuestionNotFound{ID: answerCreation.QuestionID})
}

func match(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}

type Answers []entity.Answer

func (a *Answers) Add(answer entity.AnswerCreation) entity.Answer {
	*a = append(*a, entity.Answer{
		ID: int64(len(*a) + 1),
		Content: answer.Content,
	})
	return (*a)[len(*a) - 1]
}
