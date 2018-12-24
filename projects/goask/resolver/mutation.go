package resolver

import (
	"goask/core/adapter"
	"goask/core/entity"
	"log"
)

type Mutation struct {
	Data adapter.Data
}

func (m *Mutation) Question() (QuestionMutation, error) {
	return QuestionMutation{Data: m.Data}, nil
}

type QuestionMutation struct {
	Data adapter.Data
}

// Create creates a question.
func (m QuestionMutation) Create(args struct{ Title, Content string }) (Question, error) {

	q, err := m.Data.CreateQuestion(
		entity.Post{
			Title:   args.Title,
			Content: args.Content,
		},
	)

	return (*Question).one(nil, q), err
}

// Update updates a question
func (m QuestionMutation) Update(input QuestionInput) (Question, error) {
	input.PostUpdate.ID = int64(input.ID)
	q, err := m.Data.UpdateQuestion(input.PostUpdate)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	return (*Question).one(nil, q), err
}
