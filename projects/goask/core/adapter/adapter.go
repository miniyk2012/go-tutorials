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
	CreateAnswer(QuestionID int64, Content string, AuthorID int64) (entity.Answer, error)
	AcceptAnswer(AnswerID int64, UserID int64) (entity.Answer, error)
}

// ErrQuestionOfAnswerNotFound is a data integrity error.
type ErrQuestionOfAnswerNotFound struct {
	QuestionID int64
	AnswerID int64
}

func (e *ErrQuestionOfAnswerNotFound) Error() string {
	return fmt.Sprintf("question:%d of answer:%d not found", e.QuestionID, e.AnswerID)
}

type UserDAO interface {
	UserByID(ID int64) (entity.User, error)
	Users() ([]entity.User, error)
	CreateUser(name string) (entity.User, error)
	QuestionsByUserID(ID int64) ([]entity.Question, error)
}

type ErrAnswerNotFound struct {
	ID int64
}

func (e *ErrAnswerNotFound) Error() string {
	return fmt.Sprintf("question:%d not found", e.ID)
}

type ErrUserNotFound struct {
	ID int64
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user:%d not found", e.ID)
}

type ErrUserIsNotAuthorOfQuestion struct {
	UserID int64
	QuestionID int64
}

func (e *ErrUserIsNotAuthorOfQuestion) Error() string {
	return fmt.Sprintf("user:%d is no the author of question:%d", e.UserID, e.QuestionID)
}
