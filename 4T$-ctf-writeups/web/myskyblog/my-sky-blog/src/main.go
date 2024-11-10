package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(doSessionStuff)

	e.GET("/models/:file", models)
	e.GET("/models/", models)

	e.GET("/register", register)
	e.POST("/register", register)

	e.GET("/login", login)
	e.POST("/login", login)

	e.GET("/post", post)
	e.POST("/post", post)

	e.GET("/profile", profile)
	e.POST("/profile", profile)

	e.GET("/logout", logout)

	e.GET("/flag", flag)

	e.GET("/", index)
	e.Logger.Fatal(e.Start(":8077"))
}

func flag(c echo.Context) error {
	s := c.Get("session").(*Session)

	if s.User == nil {
		return c.Redirect(http.StatusFound, "/login")
	}

	if !s.User.isAdmin {
		return c.String(http.StatusForbidden, "You are not an admin")
	}

	def := "4T${...}"
	flag := os.Getenv("FLAG")
	if flag != "" {
		def = flag
	}

	return c.String(http.StatusOK, def)
}

func models(c echo.Context) error {
	file := c.Param("file")

	if file == "" {
		// List models as <a> tags
		var out string

		ls, err := os.ReadDir("models")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading models directory")
		}

		for _, f := range ls {
			out += "<a href=\"/models/" + f.Name() + "\">" + f.Name() + "</a><br>"
		}

		return c.HTML(http.StatusOK, out)
	} else {
		// Read the file and return it as text
		f, err := os.ReadFile("models/" + file)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading file")
		}

		return c.String(http.StatusOK, string(f))
	}
}

func doSessionStuff(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If there's no session, create one in the cookie
		// If there is a session, get it from the cookie
		cookie, err := c.Cookie("session")
		if err != nil {
			session := CreateEmptySession()
			sessions[session.id] = session

			cookie = &http.Cookie{
				Name:  "session",
				Value: session.id,
				Path:  "/",
			}

			c.Set("session", session)
		} else {
			session, ok := sessions[cookie.Value]
			if !ok {
				session = CreateEmptySession()
				sessions[session.id] = session
			}

			c.Set("session", session)
		}

		cookie.Value = c.Get("session").(*Session).id
		c.SetCookie(cookie)

		return next(c)
	}
}
