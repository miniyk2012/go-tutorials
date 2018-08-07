package usecases

import (
	"testing"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters/testadapters"
	"github.com/stretchr/testify/require"
)

func TestAddMovie(t *testing.T) {
	mockAdapters := testadapters.New()

	t.Run("success", func(t *testing.T) {
		// get all movies, == 0
		movies, err := mockAdapters.GetAllMovies()
		require.Nil(t, err)
		require.Equal(t, 0, len(movies), "Have 0 movies at the beginning")

		// add user-1
		userAdapter := testadapters.NewUserAdapter()
		userAdapter.CreateUser(entities.User{
			UID: "user-1",
			AuthorizedActions: []string{"AddMovie"},
		})

		// add 1 movie
		err = AddMovie("user-1", entities.Movie{}, mockAdapters.AddMovieToDB, userAdapter)
		require.Nil(t, err)

		// get all movies, == 1, the movie been added
		movies, err = mockAdapters.GetAllMovies()
		require.Nil(t, err)
		require.Equal(t, 1, len(movies), "Have 1 movies at the end")
	})

	t.Run("access denied", func(t *testing.T) {
		userAdapter := testadapters.NewUserAdapter()
		userAdapter.CreateUser(entities.User{
			UID: "user-2",
			AuthorizedActions: []string{},
		})

		err := AddMovie("user-2", entities.Movie{}, mockAdapters.AddMovieToDB, userAdapter)
		require.Equal(t, "user:user-2 is not granted with action:AddMovie", err.Error())
	})

	t.Run("user does not exist", func(t *testing.T) {
		userAdapter := testadapters.NewUserAdapter()

		err := AddMovie("user-2", entities.Movie{}, mockAdapters.AddMovieToDB, userAdapter)
		require.Equal(t, "user:user-2 does not exist", err.Error())
	})
}
