package resolver

import (
	"goask/core/adapter"
	"goask/core/entity"
	"goask/log"
)

type Mutation struct {
	Data adapter.Data
}

func (m *Mutation) Question(args struct{ UserID int32 }) (QuestionMutation, error) {
	_, err := m.Data.UserByID(int64(args.UserID))
	if err != nil {
		return QuestionMutation{}, err
	}

	return QuestionMutation{
		stdResolver: stdResolver{
			data: m.Data,
			log:  &log.Logger{},
		},
		userSession: UserSession{
			UserID: int64(args.UserID),
		},
	}, nil
}

func (m *Mutation) Answer() (AnswerMutation, error) {
	return AnswerMutation{stdResolver: stdResolver{
		data: m.Data,
		log:  &log.Logger{},
	}}, nil
}

func (m *Mutation) User() (UserMutation, error) {
	return UserMutation{stdResolver: stdResolver{
		data: m.Data,
		log:  &log.Logger{},
	}}, nil
}

// QuestionMutation resolves all mutations of questions.
type QuestionMutation struct {
	stdResolver
	userSession UserSession
}

// Create creates a question.
func (m QuestionMutation) Create(args struct{ Title, Content string }) (Question, error) {
	if err := m.check(); err != nil {
		return Question{}, err
	}

	q, err := m.data.CreateQuestion(
		entity.Question{
			Title:    args.Title,
			Content:  args.Content,
			AuthorID: m.userSession.UserID,
		},
	)

	return QuestionOne(q, m.data), err
}

// Update updates a question
func (m QuestionMutation) Update(input QuestionInput) (Question, error) {
	if err := m.check(); err != nil {
		return Question{}, err
	}

	input.QuestionUpdate.ID = int64(input.ID)
	q, err := m.data.UpdateQuestion(input.QuestionUpdate)
	if err != nil {
		m.log.Error(err)
	}
	return QuestionOne(q, m.data), err
}

type AnswerMutation struct {
	stdResolver
}

func (m AnswerMutation) Create(args AnswerCreationInput) (Answer, error) {
	if err := m.check(); err != nil {
		return Answer{}, err
	}

	answerCreation := entity.AnswerCreation{}
	answerCreation.Content = args.Content
	answerCreation.QuestionID = int64(args.QuestionID)
	answer, err := m.data.CreateAnswer(answerCreation)
	if err != nil {
		m.log.Error(err)
	}
	return Answer{entity: answer, data: m.data}, err
}

type UserMutation struct {
	stdResolver
}

func (m UserMutation) Create(args struct{ Name string }) (User, error) {
	if err := m.check(); err != nil {
		return User{}, err
	}

	user, err := m.data.CreateUser(args.Name)
	return UserOne(user, m.data), err
}

type logger interface {
	Error(err error)
}
