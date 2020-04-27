package main

import (
	"flag"
	"fmt"
	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func main() {
	var message string
	flag.StringVar(&message, "message", "Hello", "message")
	flag.Parse()

	fmt.Println(message)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, say(message))
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func say(something string) string {
	saying, err := cowsay.Say(
		cowsay.Phrase(parsePhrase(something)),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	return saying
}

func parsePhrase(message string) string {
	lines := make([]string, 0, 40)
	for _, line := range strings.Split(message, "\n") {
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
