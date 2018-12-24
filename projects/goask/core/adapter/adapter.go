package adapter

import (
	"fmt"
	"goask/core/entity"
)

type Data interface {
	Questions() ([]entity.Question, error)
	CreateQuestion(post entity.Post) (entity.Question, error)
	UpdateQuestion(post entity.PostUpdate) (entity.Question, error)
}

type ErrQuestionNotFound struct {
	ID int64
}

func (e *ErrQuestionNotFound) Error() string {
	return fmt.Sprintf("question:%d not found", e.ID)
}
