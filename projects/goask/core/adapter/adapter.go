package adapter

import (
	"fmt"
	"goask/core/entity"
)

type Data interface {
	// Question
	Questions(search *string) ([]entity.Question, error)
	QuestionByID(ID int64) (entity.Question, error)
	CreateQuestion(post entity.Question) (entity.Question, error)
	UpdateQuestion(post entity.QuestionUpdate) (entity.Question, error)
	// Answer
	AnswersOfQuestion(QuestionID int64) []entity.Answer
	CreateAnswer(answer entity.AnswerCreation) (entity.Answer, error)
}

type ErrQuestionNotFound struct {
	ID int64
}

func (e *ErrQuestionNotFound) Error() string {
	return fmt.Sprintf("question:%d not found", e.ID)
}
