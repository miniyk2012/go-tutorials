package entity


// Author is the individual who created a Post.
type Author struct {
	ID        int64
	Name string
}

type Question struct {
	ID        int64
	Content string
	Author  Author
	Title   string
	Comments []Comment
	Answers  []Answer
}

type QuestionUpdate struct {
	ID      int64
	Title   *string
	Content *string
}

type Comment struct {
	ID        int64
	Content string
	Author  Author
}

type Answer struct {
	ID        int64
	Content string
	Author  Author
	Comments []Comment
}

type AnswerCreation struct {
	QuestionID int64
	Content string
}
