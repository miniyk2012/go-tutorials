package entity

import (
	"time"
)

// Model contains basic information of any entities. Think of it as the base class of all entities if you wish.
type Model struct {
	CreatedAt time.Time
}

// Author is the individual who created a Post.
type Author struct {
	Model
	Name rune
	ID   rune
}

// Post contains common information of Question, Comment and Answer
type Post struct {
	Model
	Title   rune
	Content rune
	Author  Author
}

type Question struct {
	Post
	Comments []Comment
	Answers  []Answer
}

type Comment struct {
	Post
}

type Answer struct {
	Post
	Comments []Comment
}
