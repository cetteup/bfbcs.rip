package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/cetteup/bfbcs.rip/cmd/server/internal/handler"
	"github.com/cetteup/bfbcs.rip/cmd/server/internal/renderer"
	"github.com/cetteup/bfbcs.rip/internal/pkg/archive"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	r, err := renderer.NewTemplateRenderer("public/layouts/*.html", "public/views/*.html")
	if err != nil {
		panic(err)
	}
	e.Renderer = r

	client := archive.NewClient(archive.BaseURL)
	h := handler.NewHandler(client)

	// Serve static files
	e.Static("/static", "public/static")

	// These all showed the same context, only differing by which platform
	// you'd search on and which leaderboard you'd be linked to in the navigation
	e.GET("/", h.HandleHomeGET("pc"))
	e.GET("/pc", h.HandleHomeGET("pc"))
	e.GET("/xbox360", h.HandleHomeGET("xbox360"))
	e.GET("/ps3", h.HandleHomeGET("ps3"))

	// Redirect old URLs for Xbox 360, which was originally referred to as just "360" in stats URLs
	e.GET("/stats_360/*", func(c *echo.Context) error {
		return c.Redirect(http.StatusFound, strings.Replace(c.Request().URL.RequestURI(), "/stats_360/", "/stats_xbox360/", 1))
	})
	e.GET("/BFBC2_leaderboard/360", func(c *echo.Context) error {
		return c.Redirect(http.StatusFound, strings.Replace(c.Request().URL.RequestURI(), "/BFBC2_leaderboard/360", "/BFBC2_leaderboard/xbox360", 1))
	})

	e.GET("/BFBC2_leaderboard/:platform", h.HandleLeaderboardGET)

	e.POST("/stats_:platform", h.HandleStatsPOST)
	e.GET("/stats_:platform/:name", h.HandleStatsGET)
	e.GET("/stats_:platform/:name/dogtags", h.HandleDogtagsGET)
	e.GET("/stats_:platform/:name/nemesis_dogtags", h.HandleNemesisDogtagsGET)

	if err = e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
