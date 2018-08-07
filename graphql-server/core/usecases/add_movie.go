package usecases

import (
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters"
)

// AddMovie add a movie to the db.
// Authorize the user.
func AddMovie(byUser string, movie entities.Movie, adder adapters.AddMovieToDB, userGetter adapters.UserGetter) error {
	user, err := userGetter.GetUser(byUser)
	if err != nil {
		return err
	}

	err = AuthorizeAddMovie(user)
	if err != nil {
		return err
	}
	return adder(movie)
}
