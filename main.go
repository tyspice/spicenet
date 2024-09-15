package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tyspice/spicenet/view/page"
)

//go:embed public/assets/*
var assets embed.FS

func main() {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "public",
		Filesystem: http.FS(assets),
	}))

	e.GET("/", func(c echo.Context) error {
		return page.Home().Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/about", func(c echo.Context) error {
		return page.About().Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
