package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (post *Post) Prepare() error {
	post.format()

	if err := post.validate(); err != nil {
		return err
	}

	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("Title is required")
	}

	if post.Content == "" {
		return errors.New("Content is required")
	}

	if post.AuthorID == 0 {
		return errors.New("Author is required")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
