package entity

type Question struct {
	ID       int64
	Content  string
	AuthorID int64 // The ID of the user who created it.
	Title    string
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
	AuthorID   int64 // The ID of the user who created it.
}

type AnswerCreation struct {
	QuestionID int64
	Content    string
}

type User struct {
	ID   int64
	Name string
}
