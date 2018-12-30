package resolver

import (
	"goask/core/adapter"
)

type Query struct {
	Data adapter.Data
}

func (q *Query) Questions(args struct{ Search *string }) ([]Question, error) {
	questions, err := q.Data.Questions(args.Search)
	return QuestionAll(questions, q.Data), err
}

func (q *Query) GetUser(args struct{ ID int32 }) (*User, error) {
	user, err := q.Data.UserByID(int64(args.ID))
	if err != nil {
		return nil, err
	}
	userResolver := UserOne(user, q.Data)
	return &userResolver, nil
}

func (q *Query) Users() ([]User, error) {
	users, err := q.Data.Users()
	if err != nil {
		return nil, err
	}
	return UserAll(users, q.Data), nil
}
