package resolver

import (
	"goask/core/adapter/fakeadapter"
	"goask/value"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResolver(t *testing.T) {

	data := &fakeadapter.Data{}

	// Query
	query := Query{Data: data}

	// Mutation
	mutation := Mutation{Data: data}
	qMutation, err := mutation.QuestionMutation(struct{ UserID int32 }{UserID: int32(1)})
	require.EqualError(t, err, "user:1 not found")

	userMutation, err := mutation.User()
	require.NoError(t, err)
	userMutation.Create(struct{ Name string }{Name: "Test User"})

	qMutation, err = mutation.QuestionMutation(struct{ UserID int32 }{UserID: int32(1)})
	require.NoError(t, err)

	// Get all Questions
	questions, err := query.Questions(struct{ Search *string }{nil})
	require.NoError(t, err)
	require.Equal(t, len(questions), 0)

	// Create Question
	qResolver, err := qMutation.Create(struct{ Title, Content string }{Title: "t", Content: "c"})
	require.NoError(t, err)
	require.Equal(t, qResolver.ID(), int32(1))
	require.Equal(t, qResolver.Content(), "c")
	require.Equal(t, qResolver.Title(), "t")

	userResolver, err := qResolver.Author()
	require.NoError(t, err)
	require.Equal(t, "Test User", userResolver.Name())

	// Update Question
	update := QuestionInput{}
	update.Content = value.String("content")
	update.ID = 1
	qResolver, err = qMutation.Update(update)
	require.NoError(t, err)
	require.Equal(t, qResolver.Title(), "t") // unchanged
	require.Equal(t, qResolver.Content(), "content")

	// Get all Questions
	questions, err = query.Questions(struct{ Search *string }{nil})
	require.NoError(t, err)
	require.Equal(t, len(questions), 1)

	t.Run("interact with answers", func(t *testing.T) {
		answerMutation, err := mutation.Answer(struct{ UserID int32 }{UserID: int32(1)})
		require.NoError(t, err)

		answer, err := answerMutation.Create(AnswerCreationInput{QuestionID: 1, Content: "This is an answer"})
		require.NoError(t, err)

		require.Equal(t, int32(1), answer.ID())
		require.Equal(t, "This is an answer", answer.Content())

		question, err := answer.Question()
		require.NoError(t, err)
		require.Equal(t, int32(1), question.ID())

		answers := question.Answers()
		require.Equal(t, 1, len(answers))
		require.Equal(t, "This is an answer", answers[0].Content())
	})
}

func TestUser(t *testing.T) {

	data := &fakeadapter.Data{}
	query := Query{Data: data}
	mutation := Mutation{Data: data}

	user, err := query.GetUser(struct{ ID int32 }{ID: 1})
	require.Error(t, err, "user:1 not found")
	require.Nil(t, user)

	userMutation, err := mutation.User()
	require.NoError(t, err)

	userResolver, err := userMutation.Create(struct{ Name string }{Name: "A Person"})
	require.NoError(t, err)
	require.Equal(t, int32(1), userResolver.ID())
	require.Equal(t, "A Person", userResolver.Name())

	userResolvers, err := query.Users()
	require.NoError(t, err)
	require.Equal(t, 1, len(userResolvers))
	require.Equal(t, "A Person", userResolvers[0].Name())

	// Questions
	questionResolvers, err := userResolver.Questions()
	require.NoError(t, err)
	require.Empty(t, questionResolvers)
}
