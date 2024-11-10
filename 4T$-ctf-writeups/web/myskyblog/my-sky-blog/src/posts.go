package main

import "time"

type Post struct {
	Author *User
	Title  string
	Body   string

	UpdatedAt time.Time
}
