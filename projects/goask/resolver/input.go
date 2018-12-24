package resolver

import (
	"goask/core/entity"
)

type QuestionInput struct {
	entity.PostUpdate
	ID int32
}
