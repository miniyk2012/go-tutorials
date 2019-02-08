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
	QuestionByID(ID entity.ID) (entity.Question, error)
	CreateQuestion(post entity.Question) (entity.Question, error)
	UpdateQuestion(post entity.QuestionUpdate) (entity.Question, error)
	DeleteQuestion(userID entity.ID, questionID entity.ID) (entity.Question, error)
}

type ErrQuestionNotFound struct {
	ID entity.ID
}

func (e *ErrQuestionNotFound) Error() string {
	return fmt.Sprintf("question:%d not found", e.ID)
}

type AnswerDAO interface {
	AnswersOfQuestion(QuestionID entity.ID) []entity.Answer
	CreateAnswer(QuestionID entity.ID, Content string, AuthorID entity.ID) (entity.Answer, error)
	AcceptAnswer(AnswerID entity.ID, UserID entity.ID) (entity.Answer, error)
}

// ErrQuestionOfAnswerNotFound is a data integrity error.
type ErrQuestionOfAnswerNotFound struct {
	QuestionID entity.ID
	AnswerID   entity.ID
}

func (e *ErrQuestionOfAnswerNotFound) Error() string {
	return fmt.Sprintf("question:%d of answer:%d not found", e.QuestionID, e.AnswerID)
}

type UserDAO interface {
	UserByID(ID entity.ID) (entity.User, error)
	Users() ([]entity.User, error)
	CreateUser(name string) (entity.User, error)
	QuestionsByUserID(ID entity.ID) ([]entity.Question, error)
}

type ErrAnswerNotFound struct {
	ID entity.ID
}

func (e *ErrAnswerNotFound) Error() string {
	return fmt.Sprintf("question:%d not found", e.ID)
}

type ErrUserNotFound struct {
	ID entity.ID
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user:%d not found", e.ID)
}

type ErrUserIsNotAuthorOfQuestion struct {
	UserID     entity.ID
	QuestionID entity.ID
}

func (e *ErrUserIsNotAuthorOfQuestion) Error() string {
	return fmt.Sprintf("user:%d is no the author of question:%d", e.UserID, e.QuestionID)
}

type ErrQuestionMutationDenied struct {
	UserID     entity.ID
	QuestionID entity.ID
}

func (e *ErrQuestionMutationDenied) Error() string {
	return fmt.Sprintf("user:%d is not authorized to delete question:%d", e.UserID, e.QuestionID)
}
