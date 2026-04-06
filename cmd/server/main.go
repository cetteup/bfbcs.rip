package main

import (
	"cmp"
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cetteup/bfbcs.rip/cmd/server/internal/handler"
	"github.com/cetteup/bfbcs.rip/cmd/server/internal/renderer"
	"github.com/cetteup/bfbcs.rip/internal/pkg/archive"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: time.RFC3339,
	})

	r, err := renderer.NewTemplateRenderer("public/layouts/*.html", "public/views/*.html")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize template renderer")
	}

	e := echo.NewWithConfig(echo.Config{
		HTTPErrorHandler: func(c *echo.Context, err error) {
			if res, _ := echo.UnwrapResponse(c.Response()); res != nil && res.Committed {
				return
			}

			code := http.StatusInternalServerError
			var sc echo.HTTPStatusCoder
			if errors.As(err, &sc) {
				code = cmp.Or(sc.StatusCode(), code)
			}

			var cerr error
			if c.Request().Method == http.MethodHead {
				cerr = c.NoContent(code)
			} else {
				cerr = c.Render(code, "default/error.html", renderer.NewPageContext(
					renderer.WithPath(c.Request().URL.Path),
					renderer.WithTitle("Error"),
					renderer.WithPlatform("pc"),
					renderer.With("Code", code),
				))
			}

			if cerr != nil {
				log.Error().
					Err(cerr).
					Msg("Failed to send error to client")
			}
		},
		Renderer:    r,
		IPExtractor: echo.ExtractIPFromXFFHeader(),
	})
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP:  true,
		LogMethod:    true,
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogUserAgent: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Err(v.Error).
				Str("remote", v.RemoteIP).
				Str("method", v.Method).
				Str("URI", v.URI).
				Int("status", v.Status).
				Stringer("latency", v.Latency).
				Str("agent", v.UserAgent).
				Msg("request")

			return nil
		},
	}))

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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         ":8080",
		HideBanner:      true,
		HidePort:        true,
		GracefulTimeout: 5 * time.Second,
		ListenerAddrFunc: func(addr net.Addr) {
			log.Info().
				Stringer("address", addr).
				Msg("Server started")
		},
	}
	if err = sc.Start(ctx, e); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
