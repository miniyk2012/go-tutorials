package usecases

import (
			"fmt"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
)

const (
	ActionAddMovie = "AddMovie"
)

func AuthorizeAddMovie(user entities.User) error {

	for _, action := range user.AuthorizedActions {
		if action == ActionAddMovie {
			return nil
		}
	}
	return fmt.Errorf("user:%s is not granted with action:%s", user.UID, ActionAddMovie)
}


