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
	users     []entity.User
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

func (d *Data) QuestionByID(ID int64) (entity.Question, error) {
	for _, q := range d.questions {
		if q.ID == ID {
			return q, nil
		}
	}
	return entity.Question{}, errors.WithStack(&adapter.ErrQuestionNotFound{ID: ID})
}

func (d *Data) QuestionsByUserID(ID int64) ([]entity.Question, error) {
	var ret []entity.Question
	for _, q := range d.questions {
		if q.AuthorID == ID {
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

func (d *Data) AnswersOfQuestion(QuestionID int64) (ret []entity.Answer) {
	for _, answer := range d.answers {
		if answer.QuestionID == QuestionID {
			ret = append(ret, answer)
		}
	}
	return
}

func (d *Data) CreateAnswer(QuestionID int64, Content string, AuthorID int64) (entity.Answer, error) {
	for _, q := range d.questions {
		if q.ID == QuestionID {
			answer := d.answers.Add(QuestionID, Content, AuthorID)
			return answer, nil
		}
	}
	return entity.Answer{}, errors.WithStack(&adapter.ErrQuestionNotFound{ID: QuestionID})
}

func (d *Data) UserByID(ID int64) (entity.User, error) {
	for _, user := range d.users {
		if user.ID == ID {
			return user, nil
		}
	}
	return entity.User{}, errors.WithStack(&adapter.ErrUserNotFound{ID: ID})
}

func (d *Data) Users() ([]entity.User, error) {
	return d.users, nil
}

func (d *Data) CreateUser(name string) (entity.User, error) {
	user := entity.User{ID: int64(len(d.users) + 1), Name: name}
	d.users = append(d.users, user)
	return user, nil
}

func match(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}

type Answers []entity.Answer

func (a *Answers) Add(QuestionID int64, Content string, AuthorID int64) entity.Answer {
	*a = append(*a, entity.Answer{
		ID:         int64(len(*a) + 1),
		Content:    Content,
		QuestionID: QuestionID,
		AuthorID: AuthorID,
	})
	return (*a)[len(*a)-1]
}

func (a *Answers) OfQuestion(questionID int64) Answers {
	var ans Answers
	for _, answer := range *a {
		if answer.QuestionID == questionID {
			ans = append(ans, answer)
		}
	}
	return ans
}
