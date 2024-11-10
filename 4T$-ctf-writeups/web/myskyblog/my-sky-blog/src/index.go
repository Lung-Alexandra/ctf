package main

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"

	"github.com/labstack/echo/v4"
)

var randomSentences = []string{
	"Hey, %s; there are %d posts right now !",
	"What's up %s ? We're at %d ! Aiming for the stars !",
	"Want a coffee %s ? Ofc there's %d posts in here !",
	"Hello there, General %s, have you seen ? There's %d posts !",
}

func index(c echo.Context) error {
	s := c.Get("session").(*Session)
	templatePath := "templates/index.html"

	indexTemplate := template.Must(template.ParseFiles(templatePath))

	if s.User == nil {
		return c.Redirect(302, "/login")
	}

	// We do a cool sentence for our users :D
	chosenSentence := fmt.Sprintf(randomSentences[rand.Intn(len(randomSentences))], s.User.Username, s.NbPosts)
	coolSentence := template.Must(indexTemplate.New("cool").Parse(chosenSentence))

	// Execute our cool template 😎
	var buf bytes.Buffer
	if err := coolSentence.Execute(&buf, s); err != nil {
		fmt.Println("Error executing template:", err)
		return indexTemplate.Execute(c.Response().Writer, s)
	}

	s.CoolSentence = buf.String()

	return indexTemplate.Execute(c.Response().Writer, s)
}
