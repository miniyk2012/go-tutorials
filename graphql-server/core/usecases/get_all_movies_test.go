package usecases

import (
	"testing"
		"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters/testadapters"
	"github.com/stretchr/testify/require"
)

func TestGetAllMovies(t *testing.T) {
	mockAdapters := testadapters.New()

	movies, err := GetAllMovies(mockAdapters.GetAllMovies)
	require.Nil(t, err)
	require.Nil(t, movies)
}
