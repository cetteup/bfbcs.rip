package handler

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/labstack/echo/v5"

	"github.com/cetteup/bfbcs.rip/cmd/server/internal/renderer"
	"github.com/cetteup/bfbcs.rip/internal/pkg/archive"
)

type client interface {
	GetStats(ctx context.Context, platform string, name string) (archive.StatsResponse, error)
	GetDogtags(ctx context.Context, platform string, name string) (archive.DogtagsResponse, error)
}

type Handler struct {
	client client
}

func NewHandler(client client) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) HandleHomeGET(platform string) echo.HandlerFunc {
	players := []struct {
		Name     string
		Platform string
		Rank     float64
		Score    float64
		Updated  string
	}{
		{Name: "jjbird 104", Platform: "xbox360", Rank: 16.0, Score: 268230.0, Updated: "2023-12-08T22:31:04Z"},
		{Name: "Hi im Swat", Platform: "xbox360", Rank: 23.0, Score: 448024.0, Updated: "2023-12-08T22:30:48Z"},
		{Name: "EpicNeener", Platform: "xbox360", Rank: 0.0, Score: 4585.0, Updated: "2023-12-08T22:30:07Z"},
		{Name: "pretendeffort", Platform: "xbox360", Rank: 6.0, Score: 65742.0, Updated: "2023-12-08T22:25:32Z"},
		{Name: "Macedon1006", Platform: "xbox360", Rank: 22.0, Score: 429885.0, Updated: "2023-12-08T22:25:30Z"},
		{Name: "Cyrus Dorian", Platform: "xbox360", Rank: 33.0, Score: 1600010.0, Updated: "2023-12-08T22:25:21Z"},
		{Name: "It Is Empty", Platform: "xbox360", Rank: 0.0, Score: 1340.0, Updated: "2023-12-08T22:25:20Z"},
		{Name: "FaithWarrior07", Platform: "xbox360", Rank: 0.0, Score: 4926.0, Updated: "2023-12-08T22:25:06Z"},
		{Name: "Iamomenchild", Platform: "xbox360", Rank: 14.0, Score: 223793.0, Updated: "2023-12-08T22:20:42Z"},
		{Name: "Arebos", Platform: "xbox360", Rank: 25.0, Score: 540946.0, Updated: "2023-12-08T22:18:30Z"},
		{Name: "I3ruce W4yne", Platform: "xbox360", Rank: 6.0, Score: 54827.0, Updated: "2023-12-08T22:18:22Z"},
		{Name: "DubGang386", Platform: "xbox360", Rank: 50.0, Score: 7233960.0, Updated: "2023-12-08T22:17:39Z"},
		{Name: "Boosting Drone", Platform: "xbox360", Rank: 1.0, Score: 8400.0, Updated: "2023-12-08T22:17:36Z"},
		{Name: "GeekHard88", Platform: "xbox360", Rank: 0.0, Score: 1920.0, Updated: "2023-12-08T22:17:26Z"},
		{Name: "Fenixxrd", Platform: "xbox360", Rank: 8.0, Score: 94505.0, Updated: "2023-12-08T22:17:12Z"},
		{Name: "howbowdahcash", Platform: "xbox360", Rank: 10.0, Score: 129582.0, Updated: "2023-12-08T22:16:53Z"},
		{Name: "ixNEGANxi", Platform: "xbox360", Rank: 50.0, Score: 41599400, Updated: "2023-12-08T22:16:13Z"},
		{Name: "Koluban", Platform: "pc", Rank: 50.0, Score: 6084580.0, Updated: "2023-12-08T22:11:13Z"},
		{Name: "Mayxhaha", Platform: "pc", Rank: 50.0, Score: 5738690.0, Updated: "2023-12-08T22:11:00Z"},
		{Name: "KhusainovMoH2010", Platform: "pc", Rank: 50.0, Score: 8032250.0, Updated: "2023-12-08T22:10:58Z"},
		{Name: "gefaut", Platform: "pc", Rank: 34.0, Score: 1769880.0, Updated: "2023-12-08T22:10:41Z"},
		{Name: "comdy", Platform: "pc", Rank: 18.0, Score: 309270.0, Updated: "2023-12-08T22:10:38Z"},
		{Name: "MistyForest14", Platform: "ps3", Rank: 3.0, Score: 23475.0, Updated: "2023-12-08T22:10:36Z"},
		{Name: "DasM3h", Platform: "pc", Rank: 6.0, Score: 54909.0, Updated: "2023-12-08T22:10:34Z"},
		{Name: "Kormi94", Platform: "pc", Rank: 39.0, Score: 2884960.0, Updated: "2023-12-08T22:10:31Z"},
		{Name: "honda-militar0", Platform: "ps3", Rank: 12.0, Score: 176663.0, Updated: "2023-12-08T22:10:31Z"},
		{Name: "XalaX_Tu", Platform: "ps3", Rank: 0.0, Score: 1150.0, Updated: "2023-12-08T22:10:26Z"},
		{Name: "leonard_k", Platform: "ps3", Rank: 50.0, Score: 13259800, Updated: "2023-12-08T22:05:31Z"},
		{Name: "OscarNovember", Platform: "ps3", Rank: 5.0, Score: 47936.0, Updated: "2023-12-08T22:05:28Z"},
		{Name: "digital_fly0", Platform: "ps3", Rank: 37.0, Score: 2353240.0, Updated: "2023-12-08T22:05:12Z"},
		{Name: "NYE-AM0R", Platform: "ps3", Rank: 50.0, Score: 8872560.0, Updated: "2023-12-08T22:00:23Z"},
		{Name: "jazzmarce", Platform: "ps3", Rank: 23.0, Score: 459145.0, Updated: "2023-12-08T22:00:12Z"},
		{Name: "20RAAP01", Platform: "pc", Rank: 31.0, Score: 1261160.0, Updated: "2023-12-08T21:55:25Z"},
		{Name: "truegrit952", Platform: "pc", Rank: 50.0, Score: 52132800, Updated: "2023-12-08T21:55:23Z"},
		{Name: "rcrlj", Platform: "ps3", Rank: 2.0, Score: 14542.0, Updated: "2023-12-08T21:55:15Z"},
		{Name: "Influencer", Platform: "pc", Rank: 50.0, Score: 7843080.0, Updated: "2023-12-08T21:50:50Z"},
		{Name: "tom-wlkp", Platform: "pc", Rank: 50.0, Score: 44314000, Updated: "2023-12-08T21:50:49Z"},
		{Name: "tkasch95", Platform: "pc", Rank: 50.0, Score: 6509620.0, Updated: "2023-12-08T21:50:47Z"},
		{Name: "ofkin312", Platform: "pc", Rank: 50.0, Score: 24947000, Updated: "2023-12-08T21:50:43Z"},
		{Name: "Don_Maulos", Platform: "ps3", Rank: 43.0, Score: 3693220.0, Updated: "2023-12-08T21:50:38Z"},
		{Name: "Matanuska", Platform: "ps3", Rank: 28.0, Score: 930132.0, Updated: "2023-12-08T21:50:35Z"},
		{Name: "[[ XX ]]", Platform: "pc", Rank: 50.0, Score: 511573000, Updated: "2023-12-08T21:50:32Z"},
		{Name: "ToNiNoo", Platform: "pc", Rank: 50.0, Score: 9035710.0, Updated: "2023-12-08T21:50:31Z"},
		{Name: "soviet_onion981", Platform: "ps3", Rank: 25.0, Score: 568911.0, Updated: "2023-12-08T21:50:31Z"},
		{Name: "pincaman66", Platform: "ps3", Rank: 40.0, Score: 2917580.0, Updated: "2023-12-08T21:50:18Z"},
		{Name: "I-Zurvan-I", Platform: "ps3", Rank: 46.0, Score: 4563000.0, Updated: "2023-12-08T21:45:35Z"},
		{Name: "Shhua94", Platform: "ps3", Rank: 38.0, Score: 2600770.0, Updated: "2023-12-08T21:45:20Z"},
		{Name: "Driven-007_", Platform: "ps3", Rank: 13.0, Score: 194758.0, Updated: "2023-12-08T21:40:42Z"},
		{Name: "Vovca51rus", Platform: "pc", Rank: 50.0, Score: 6923390.0, Updated: "2023-12-08T21:40:34Z"},
		{Name: "Aspero Vargos", Platform: "pc", Rank: 10.0, Score: 123529.0, Updated: "2023-12-08T21:40:30Z"},
		{Name: "killer_945000", Platform: "ps3", Rank: 39.0, Score: 2783240.0, Updated: "2023-12-08T21:40:30Z"},
		{Name: "DonnyScott", Platform: "pc", Rank: 24.0, Score: 511548.0, Updated: "2023-12-08T21:40:25Z"},
		{Name: "REssa", Platform: "pc", Rank: 50.0, Score: 69890500, Updated: "2023-12-08T21:36:14Z"},
		{Name: "Sgt_Feels", Platform: "pc", Rank: 50.0, Score: 407991000, Updated: "2023-12-08T21:36:05Z"},
		{Name: "KristopherVector", Platform: "pc", Rank: 40.0, Score: 3074080.0, Updated: "2023-12-08T21:35:55Z"},
		{Name: "brzydula1982", Platform: "ps3", Rank: 10.0, Score: 137567.0, Updated: "2023-12-08T21:35:53Z"},
		{Name: "NoSupper", Platform: "pc", Rank: 48.0, Score: 4922460.0, Updated: "2023-12-08T21:35:53Z"},
		{Name: "Yiaah", Platform: "pc", Rank: 50.0, Score: 12894100, Updated: "2023-12-08T21:35:50Z"},
		{Name: "Aft15525", Platform: "xbox360", Rank: 4.0, Score: 30635.0, Updated: "2023-12-08T21:35:33Z"},
		{Name: "Sgt WarDaddy98", Platform: "xbox360", Rank: 30.0, Score: 1213500.0, Updated: "2023-12-08T21:35:28Z"},
		{Name: "MR_SLYDING", Platform: "ps3", Rank: 39.0, Score: 2783910.0, Updated: "2023-12-08T21:35:24Z"},
		{Name: "Ghostmonkey_vzl", Platform: "ps3", Rank: 21.0, Score: 393239.0, Updated: "2023-12-08T21:35:21Z"},
		{Name: "Jaolet67", Platform: "ps3", Rank: 42.0, Score: 3384000.0, Updated: "2023-12-08T21:35:16Z"},
		{Name: "Basis8892", Platform: "ps3", Rank: 21.0, Score: 378316.0, Updated: "2023-12-08T21:35:15Z"},
		{Name: "infidel 11207", Platform: "pc", Rank: 35.0, Score: 2000490.0, Updated: "2023-12-08T21:30:56Z"},
		{Name: "John_Hennry_III", Platform: "ps3", Rank: 22.0, Score: 435470.0, Updated: "2023-12-08T21:30:48Z"},
		{Name: "...seb...__", Platform: "pc", Rank: 50.0, Score: 29726900, Updated: "2023-12-08T21:30:45Z"},
		{Name: "mister249", Platform: "pc", Rank: 50.0, Score: 32142900, Updated: "2023-12-08T21:30:32Z"},
		{Name: "Kirilitos", Platform: "pc", Rank: 4.0, Score: 33036.0, Updated: "2023-12-08T21:30:24Z"},
		{Name: "Archades003", Platform: "ps3", Rank: 25.0, Score: 611337.0, Updated: "2023-12-08T21:30:18Z"},
		{Name: "KuniBoost", Platform: "ps3", Rank: 0.0, Score: 0.0, Updated: "2023-12-08T21:30:15Z"},
		{Name: "xSENZURICHAMPx", Platform: "pc", Rank: 50.0, Score: 9067740.0, Updated: "2023-12-08T21:25:44Z"},
		{Name: "Yolorito81", Platform: "ps3", Rank: 49.0, Score: 5171960.0, Updated: "2023-12-08T21:25:20Z"},
		{Name: "WidowsSon23", Platform: "ps3", Rank: 25.0, Score: 561812.0, Updated: "2023-12-08T21:25:11Z"},
		{Name: "The ZarbaZan", Platform: "pc", Rank: 49.0, Score: 5344070.0, Updated: "2023-12-08T21:21:06Z"},
		{Name: "colin_wilson_5", Platform: "pc", Rank: 50.0, Score: 15313300, Updated: "2023-12-08T21:21:04Z"},
		{Name: "Garpoonov", Platform: "pc", Rank: 50.0, Score: 7382380.0, Updated: "2023-12-08T21:21:02Z"},
		{Name: "JadeSergal", Platform: "ps3", Rank: 3.0, Score: 23351.0, Updated: "2023-12-08T21:20:41Z"},
		{Name: "svvs", Platform: "pc", Rank: 50.0, Score: 8608180.0, Updated: "2023-12-08T21:20:39Z"},
		{Name: "Zzz_SIBAY_zzZ", Platform: "ps3", Rank: 37.0, Score: 2311280.0, Updated: "2023-12-08T21:20:33Z"},
		{Name: "On2022", Platform: "pc", Rank: 50.0, Score: 66841200, Updated: "2023-12-08T21:20:28Z"},
		{Name: "ROBZI77A", Platform: "ps3", Rank: 25.0, Score: 541406.0, Updated: "2023-12-08T21:16:27Z"},
		{Name: "Svahren", Platform: "ps3", Rank: 37.0, Score: 2321260.0, Updated: "2023-12-08T21:16:00Z"},
		{Name: "killer_hippi", Platform: "ps3", Rank: 2.0, Score: 16887.0, Updated: "2023-12-08T21:15:28Z"},
		{Name: "TR3BORIS", Platform: "xbox360", Rank: 39.0, Score: 2894490.0, Updated: "2023-12-08T21:15:18Z"},
		{Name: "guion-mundial2", Platform: "ps3", Rank: 24.0, Score: 507762.0, Updated: "2023-12-08T21:15:15Z"},
		{Name: "GuardianAAngel", Platform: "pc", Rank: 30.0, Score: 1093260.0, Updated: "2023-12-08T21:10:31Z"},
		{Name: "HHOURSEPKA13", Platform: "xbox360", Rank: 29.0, Score: 1080160.0, Updated: "2023-12-08T20:50:19Z"},
		{Name: "Becky Thump", Platform: "xbox360", Rank: 15.0, Score: 243554.0, Updated: "2023-12-08T20:25:24Z"},
		{Name: "Damips9334", Platform: "xbox360", Rank: 3.0, Score: 23043.0, Updated: "2023-12-08T20:25:10Z"},
		{Name: "Bilbo swagin574", Platform: "xbox360", Rank: 46.0, Score: 4335910.0, Updated: "2023-12-08T18:25:13Z"},
		{Name: "GreenPhantom259", Platform: "xbox360", Rank: 40.0, Score: 2937440.0, Updated: "2023-12-08T09:55:14Z"},
		{Name: "Spencecilence", Platform: "xbox360", Rank: 3.0, Score: 23592.0, Updated: "2023-12-08T06:05:39Z"},
		{Name: "TreeGoClimb", Platform: "xbox360", Rank: 23.0, Score: 440445.0, Updated: "2023-12-08T05:35:29Z"},
		{Name: "GodlyHam", Platform: "xbox360", Rank: 25.0, Score: 584616.0, Updated: "2023-12-08T05:30:31Z"},
		{Name: "FastZander101", Platform: "xbox360", Rank: 1.0, Score: 7955.0, Updated: "2023-12-08T03:45:07Z"},
		{Name: "Richar Aamon", Platform: "xbox360", Rank: 40.0, Score: 2983790.0, Updated: "2023-12-08T01:25:17Z"},
		{Name: "Fabio1279", Platform: "xbox360", Rank: 33.0, Score: 1675590.0, Updated: "2023-12-07T23:50:08Z"},
		{Name: "Xenocry90", Platform: "xbox360", Rank: 0.0, Score: 0.0, Updated: "2023-12-07T00:25:11Z"},
		{Name: "Falling Cos", Platform: "xbox360", Rank: 3.0, Score: 26792.0, Updated: "2023-11-15T18:49:05Z"},
	}

	return func(c *echo.Context) error {
		return c.Render(http.StatusOK, "default/home.html", renderer.NewPageContext(
			renderer.WithPath(c.Request().URL.Path),
			renderer.WithTitle(fmt.Sprintf("%s - Battlefield Bad Company 2 Stats", strings.ToUpper(platform))),
			renderer.WithPlatform(platform),
			renderer.With("Players", players),
		))
	}
}

func (h *Handler) HandleLeaderboardGET(c *echo.Context) error {
	params := struct {
		Platform string `param:"platform"`
	}{}

	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Wrap(err)
	}

	return c.Render(http.StatusOK, "default/leaderboard.html", renderer.NewPageContext(
		renderer.WithPath(c.Request().URL.Path),
		renderer.WithTitle("Leaderboard - BFBC2 Stats"),
		renderer.WithPlatform(params.Platform),
	))
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
			return c.Render(http.StatusNotFound, "default/stats-not-found.html", renderer.NewPageContext(
				renderer.WithPath(c.Request().URL.Path),
				renderer.WithTitle(fmt.Sprintf("%s - BFBC2 Stats", params.Name)),
				renderer.WithPlatform(params.Platform),
				renderer.With("Player", archive.Player{
					Name:     params.Name,
					Platform: params.Platform,
				}),
			))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)).Wrap(err)
	}

	return c.Render(http.StatusOK, "default/stats.html", renderer.NewPageContext(
		renderer.WithPath(c.Request().URL.Path),
		renderer.WithTitle(fmt.Sprintf("%s - BFBC2 Stats", stats.Player.Name)),
		renderer.WithPlatform(stats.Player.Platform),
		renderer.With("Player", stats.Player),
		renderer.With("Values", stats.Values),
	))
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

func (h *Handler) HandleDogtagsGET(c *echo.Context) error {
	params := struct {
		Platform string `param:"platform"`
		Name     string `param:"name"`
	}{}

	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Wrap(err)
	}

	stats, serr := h.client.GetStats(c.Request().Context(), params.Platform, params.Name)
	dogtags, derr := h.client.GetDogtags(c.Request().Context(), params.Platform, params.Name)
	if err := cmp.Or(serr, derr); err != nil {
		if errors.Is(err, archive.ErrPlayerNotFound) {
			return c.Render(http.StatusNotFound, "default/dogtags-not-found.html", renderer.NewPageContext(
				renderer.WithPath(c.Request().URL.Path),
				renderer.WithTitle(fmt.Sprintf("%s - Dogtags", params.Name)),
				renderer.WithPlatform(params.Platform),
				renderer.With("Player", archive.Player{
					Name:     params.Name,
					Platform: params.Platform,
				}),
			))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)).Wrap(err)
	}

	// Maintain original site's sort order (total dogtags descending, no tiebreaker)
	slices.SortStableFunc(dogtags.Records, func(a, b archive.DogtagRecord) int {
		return cmp.Compare(b.Total, a.Total)
	})

	summary := struct {
		Total        int
		TotalUnique  int
		Bronze       int
		BronzeUnique int
		Silver       int
		SilverUnique int
		Gold         int
		GoldUnique   int
	}{}

	for _, record := range dogtags.Records {
		summary.Total += record.Total
		summary.Bronze += record.Bronze
		summary.Silver += record.Silver
		summary.Gold += record.Gold

		if record.Total > 0 {
			summary.TotalUnique += 1
		}
		if record.Bronze > 0 {
			summary.BronzeUnique += 1
		}
		if record.Silver > 0 {
			summary.SilverUnique += 1
		}
		if record.Gold > 0 {
			summary.GoldUnique += 1
		}
	}

	return c.Render(http.StatusOK, "default/dogtags.html", renderer.NewPageContext(
		renderer.WithPath(c.Request().URL.Path),
		renderer.WithTitle(fmt.Sprintf("%s - Dogtags", stats.Player.Name)),
		renderer.WithPlatform(stats.Player.Platform),
		renderer.With("Player", stats.Player),
		renderer.With("Values", stats.Values),
		renderer.With("Records", dogtags.Records),
		renderer.With("Summary", summary),
	))
}
