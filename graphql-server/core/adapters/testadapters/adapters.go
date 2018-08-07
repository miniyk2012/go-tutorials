package testadapters

import (
		"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
	"fmt"
)

func New() Adapters {
	return Adapters{}
}

type Adapters struct {
	movies []entities.Movie

}

func (a *Adapters) GetAllMovies() ([]entities.Movie, error) {
	return a.movies, nil
}

func (a *Adapters) AddMovieToDB(movie entities.Movie) error {
	a.movies = append(a.movies, movie)
	return nil
}


type UserAdapter struct {
	users  map[string]entities.User
}

func NewUserAdapter() *UserAdapter {
	return &UserAdapter{
		users: make(map[string]entities.User),
	}
}


func (a *UserAdapter) GetUser(uid string) (entities.User, error) {
	if user, ok := a.users[uid]; ok {
		return user, nil
	}
	return entities.User{}, fmt.Errorf("user:%s does not exist", uid)
}

func (a *UserAdapter) CreateUser(user entities.User) (entities.User, error) {
	a.users[user.UID] = user
	return a.users[user.UID], nil
}
