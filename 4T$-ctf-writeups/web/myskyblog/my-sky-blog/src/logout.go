package main

import "github.com/labstack/echo/v4"

func logout(c echo.Context) error {
	// Nooooo, don't leave me 😭

	s := c.Get("session").(*Session)
	s.User = nil
	return c.Redirect(302, "/login")
}
