package resolver

import (
	"goask/core/adapter/fakeadapter"
	"goask/value"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuery_Questions(t *testing.T) {

	data := &fakeadapter.Data{}

	// Query
	query := Query{
		Data: data,
	}

	// Get all Questions
	questions, err := query.Questions(struct{Search *string}{nil})
	require.NoError(t, err)
	require.Equal(t, len(questions), 0)

	// Mutation
	mutation := Mutation{Data: data}

	qMutation, err := mutation.Question()
	require.NoError(t, err)

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
	questions, err = query.Questions(struct{Search *string}{nil})
	require.NoError(t, err)
	require.Equal(t, len(questions), 1)
}
