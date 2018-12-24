package entity

import (
	"time"
)

// Model contains basic information of any entities. Think of it as the base class of all entities if you wish.
type Model struct {
	ID        int64
	CreatedAt time.Time
}

// Author is the individual who created a Post.
type Author struct {
	Model
	Name string
	ID   string
}

// Post contains common information of Question, Comment and Answer
type Post struct {
	Model
	Title   string
	Content string
	Author  Author
}

type PostUpdate struct {
	ID      int64
	Title   *string
	Content *string
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
