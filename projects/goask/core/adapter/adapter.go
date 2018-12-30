package adapter

import (
	"fmt"
	"goask/core/entity"
)

type Data interface {
	QuestionDAO
	AnswerDAO
	UserDAO
}

type QuestionDAO interface {
	Questions(search *string) ([]entity.Question, error)
	QuestionByID(ID int64) (entity.Question, error)
	CreateQuestion(post entity.Question) (entity.Question, error)
	UpdateQuestion(post entity.QuestionUpdate) (entity.Question, error)
}

type ErrQuestionNotFound struct {
	ID int64
}

func (e *ErrQuestionNotFound) Error() string {
	return fmt.Sprintf("question:%d not found", e.ID)
}

type AnswerDAO interface {
	AnswersOfQuestion(QuestionID int64) []entity.Answer
	CreateAnswer(answer entity.AnswerCreation) (entity.Answer, error)
}

type UserDAO interface {
	UserByID(ID int64) (entity.User, error)
	Users() ([]entity.User, error)
	CreateUser(name string) (entity.User, error)
}

type ErrUserNotFound struct {
	ID int64
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user:%d not found", e.ID)
}
