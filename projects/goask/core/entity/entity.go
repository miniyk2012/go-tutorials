package entity

// Author is the individual who created a Post.
type Author struct {
	ID   int64
	Name string
}

type Question struct {
	ID      int64
	Content string
	Author  Author
	Title   string
}

type QuestionUpdate struct {
	ID      int64
	Title   *string
	Content *string
}

type Answer struct {
	ID         int64
	QuestionID int64
	Content    string
	Author     Author
}

type AnswerCreation struct {
	QuestionID int64
	Content    string
}

type User struct {
	ID   int64
	Name string
}
