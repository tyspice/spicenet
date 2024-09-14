package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/tyspice/spicenet/view/page"
)

func readFileContents(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	e := echo.New()
	e.Static("/assets", "./public/assets")

	heading, err := readFileContents("./assets/heading.txt")
	if err != nil {
		e.Logger.Error(err)
	}

	e.GET("/", func(c echo.Context) error {
		return page.Home(heading).Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/about", func(c echo.Context) error {
		return page.About(heading).Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
