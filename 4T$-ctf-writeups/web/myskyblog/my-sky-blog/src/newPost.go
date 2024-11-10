package main

import (
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

func post(c echo.Context) error {
	s := c.Get("session").(*Session)
	templatePath := "templates/post.html"
	tpl := template.Must(template.ParseFiles(templatePath))
	if s.User == nil {
		return c.Redirect(302, "/login")
	}

	if c.Request().Method == "GET" {
		return tpl.Execute(c.Response().Writer, s)
	}

	if c.Request().Method == "POST" {
		title := c.FormValue("title")
		body := c.FormValue("body")

		p := &Post{
			Author:    s.User,
			Title:     title,
			Body:      body,
			UpdatedAt: time.Now(),
		}

		s.Posts = append(s.Posts, p)
		s.NbPosts++
	}

	// Redirect to /
	return c.Redirect(302, "/")
}
