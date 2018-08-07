package usecases

import (
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
)

// GetAllMovies get all movies in the db and return it to the user.
func GetAllMovies(getter adapters.GetAllMovies ) ([]entities.Movie, error) {
	movies, err := getter()
	return movies, err
}

