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
	qMutation, err := mutation.Question()
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
		answerMutation, err := mutation.Answer()
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
