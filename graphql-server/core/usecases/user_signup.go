package usecases

import (
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters"
)

func UserSignUp(user entities.User, creator adapters.UserCreator) (entities.User, error) {
	return creator.CreateUser(user)
}
