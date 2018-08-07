package adapters

import (
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
)

// Movie
type GetAllMovies func() ([]entities.Movie, error)
type AddMovieToDB func(entities.Movie) error


// User
type User interface {
	UserGetter
	UserCreator
}

type UserGetter interface {
	GetUser(uid string) (entities.User, error)
}

type UserCreator interface {
	CreateUser(user entities.User) (entities.User, error)
}
