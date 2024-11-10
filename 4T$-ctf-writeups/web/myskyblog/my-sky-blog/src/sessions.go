package main

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	// Private fields
	users []*User
	id    string

	// Public fields
	User    *User
	Posts   []*Post
	NbPosts int

	HasError bool
	Error    string

	CoolSentence string
}

var sessions = make(map[string]*Session)

func CreateEmptySession() *Session {
	admin := &User{
		isAdmin:  true,
		Username: "admin",
	}

	// Get a random password
	randomPassword := uuid.New().String()

	admin.ChangePassword(randomPassword)

	id := uuid.New().String()

	return &Session{
		users: []*User{
			admin,
		},
		id: id,

		User: nil,
		Posts: []*Post{
			{
				Author:    admin,
				Title:     "Welcome to my beautiful Sky Blog!",
				Body:      "I welcome you to my blog, where I'll post about my adventures in the sky !",
				UpdatedAt: time.Date(2024, 5, 1, 12, 54, 30, 20, time.UTC),
			},
		},
		NbPosts: 1,
	}
}
