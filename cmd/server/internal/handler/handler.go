package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v5"

	"github.com/cetteup/bfbcs.rip/internal/pkg/archive"
)

type client interface {
	GetStats(ctx context.Context, platform string, name string) (archive.StatsResponse, error)
}

type Handler struct {
	client client
}

func NewHandler(client client) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) HandleStatsGET(c *echo.Context) error {
	params := struct {
		Platform string `param:"platform"`
		Name     string `param:"name"`
	}{}

	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Wrap(err)
	}

	stats, err := h.client.GetStats(c.Request().Context(), params.Platform, params.Name)
	if err != nil {
		if errors.Is(err, archive.ErrPlayerNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Player not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)).Wrap(err)
	}

	return c.Render(http.StatusOK, "stats.html", stats)
}

func (h *Handler) HandleStatsPOST(c *echo.Context) error {
	params := struct {
		Platform string `param:"platform"`
		Name     string `form:"searchplayer[name]"`
	}{}

	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Wrap(err)
	}

	return c.Redirect(
		http.StatusFound,
		fmt.Sprintf("/stats_%s/%s", url.PathEscape(params.Platform), url.PathEscape(params.Name)),
	)
}
