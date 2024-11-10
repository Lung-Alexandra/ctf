package main

import (
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func register(c echo.Context) error {
	s := c.Get("session").(*Session)

	templatePath := "templates/register.html"
	tpl := template.Must(template.ParseFiles(templatePath))

	if c.Request().Method == "POST" {
		username := c.FormValue("username")
		password := c.FormValue("password")

		user := &User{
			Username: username,
		}

		user.ChangePassword(password)

		for _, u := range s.users {
			if u.Username == user.Username {
				s.HasError = true
				s.Error = "Username already taken"
				return tpl.Execute(c.Response().Writer, s)
			}
		}

		s.users = append(s.users, user)

		logrus.Infof("Registered user %s on sessions %s", user.Username, s.id)

		return c.Redirect(302, "/login")
	} else {
		return tpl.Execute(c.Response().Writer, s)
	}
}

func login(c echo.Context) error {
	s := c.Get("session").(*Session)
	templatePath := "templates/login.html"
	tpl := template.Must(template.ParseFiles(templatePath))

	// Super 😎 login method !
	// Only chads can do code like this
	if c.Request().Method == "POST" {
		username := c.FormValue("username")
		password := c.FormValue("password")

		var user *User
		for _, u := range s.users {
			if u.Username == username {
				user = u
				break
			}
		}

		if user == nil || !user.CheckPassword(password) {
			s.HasError = true
			s.Error = "Invalid username or password"
			return tpl.Execute(c.Response().Writer, s)
		}

		s.User = user

		logrus.Infof("Logged in user %s on session %s", user.Username, s.id)

		return c.Redirect(302, "/")
	} else {
		return tpl.Execute(c.Response().Writer, s)
	}
}
