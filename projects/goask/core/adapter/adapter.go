package adapter

import (
	"fmt"
	"goask/core/entity"
)

type Data interface {
	Questions(search *string) ([]entity.Question, error)
	CreateQuestion(post entity.Question) (entity.Question, error)
	UpdateQuestion(post entity.QuestionUpdate) (entity.Question, error)
	CreateAnswer(answer entity.AnswerCreation) (entity.Answer, error)
}

type ErrQuestionNotFound struct {
	ID int64
}

func (e *ErrQuestionNotFound) Error() string {
	return fmt.Sprintf("question:%d not found", e.ID)
}
