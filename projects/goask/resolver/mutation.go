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

func (m *Mutation) Answer() (AnswerMutation, error) {
	return AnswerMutation{Data: m.Data}, nil
}

// QuestionMutation resolves all mutations of questions.
type QuestionMutation struct {
	Data adapter.Data
}

// Create creates a question.
func (m QuestionMutation) Create(args struct{ Title, Content string }) (Question, error) {

	q, err := m.Data.CreateQuestion(
		entity.Question{
			Title:   args.Title,
			Content: args.Content,
		},
	)

	return QuestionOne(q, m.Data), err
}

// Update updates a question
func (m QuestionMutation) Update(input QuestionInput) (Question, error) {
	input.QuestionUpdate.ID = int64(input.ID)
	q, err := m.Data.UpdateQuestion(input.QuestionUpdate)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	return QuestionOne(q, m.Data), err
}

type AnswerMutation struct {
	Data adapter.Data
}

func (m AnswerMutation) Create(args AnswerCreationInput) (Answer, error) {
	answerCreation := entity.AnswerCreation{}
	answerCreation.Content = args.Content
	answerCreation.QuestionID = int64(args.QuestionID)
	answer, err := m.Data.CreateAnswer(answerCreation)
	if err != nil {
		log.Println(err) // todo: inject a logger
	}
	return Answer{entity: answer, data: m.Data}, err
}
