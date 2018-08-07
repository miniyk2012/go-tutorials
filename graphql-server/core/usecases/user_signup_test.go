package usecases

import (
	"testing"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters/testadapters"
	"github.com/stretchr/testify/require"
)

func TestUserSignUp(t *testing.T) {

	userAdapter := testadapters.NewUserAdapter()

	user1, err := UserSignUp(entities.User{UID: "xxx"}, userAdapter)
	require.Nil(t, err)

	user2, err := userAdapter.GetUser("xxx")
	require.Nil(t, err)

	require.Equal(t, user1, user2)
	require.Equal(t, user1, entities.User{
		UID: "xxx",
		AuthorizedActions: nil,
	})
}
