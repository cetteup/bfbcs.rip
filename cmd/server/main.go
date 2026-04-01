package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/cetteup/bfbcs.rip/cmd/server/internal/handler"
	"github.com/cetteup/bfbcs.rip/cmd/server/internal/renderer"
	"github.com/cetteup/bfbcs.rip/internal/pkg/archive"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	r, err := renderer.NewTemplateRenderer("public/views/*.html")
	if err != nil {
		panic(err)
	}
	e.Renderer = r

	client := archive.NewClient(archive.BaseURL)
	h := handler.NewHandler(client)

	// Serve static files
	e.Static("/static", "public/static")

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/stats_:platform/:name", h.HandleStatsGET)
	e.POST("/stats_:platform", h.HandleStatsPOST)

	if err = e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
