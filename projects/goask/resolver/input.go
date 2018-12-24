package resolver

import (
	"goask/core/entity"
)

type QuestionInput struct {
	entity.QuestionUpdate
	ID int32
}
