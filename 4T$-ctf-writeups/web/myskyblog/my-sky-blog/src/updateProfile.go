package main

import (
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func profile(c echo.Context) error {
	s := c.Get("session").(*Session)
	templatePath := "templates/profile.html"
	tpl := template.Must(template.ParseFiles(templatePath))
	if s.User == nil {
		return c.Redirect(302, "/login")
	}

	if c.Request().Method == "GET" {
		return tpl.Execute(c.Response().Writer, s)
	}

	if c.Request().Method == "POST" {
		username := c.FormValue("username")
		password := c.FormValue("password")

		logrus.Infof("Updating profile for user %s to %s on session %s", s.User.Username, username, s.id)
		s.User.Username = username
		if password != "" {
			s.User.ChangePassword(password)
		}
	}

	// Redirect to /
	return c.Redirect(302, "/")
}
