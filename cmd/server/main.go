package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Player struct {
	Pid       int    `json:"pid"`
	Name      string `json:"name"`
	Platform  string `json:"platform"`
	Namespace string `json:"namespace"`
	Added     string `json:"added"`
	Updated   string `json:"updated"`
}

type PlayerStats struct {
	Player Player                 `json:"player"`
	Values map[string]interface{} `json:"values"`
}

type TemplateRenderer struct {
	templates *template.Template
}

func (tr *TemplateRenderer) Render(_ *echo.Context, w io.Writer, name string, data any) error {
	return tr.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	// Template functions
	funcMap := template.FuncMap{
		"div": func(a, b float64) float64 { return a / b },
		"mul": func(a, b float64) float64 { return a * b },
		"gt":  func(a, b float64) bool { return a > b },
		"formatNumber": func(n float64) string {
			i := int64(n)
			s := strconv.FormatInt(i, 10)
			var parts []string
			for len(s) > 3 {
				parts = append([]string{s[len(s)-3:]}, parts...)
				s = s[:len(s)-3]
			}
			if s != "" {
				parts = append([]string{s}, parts...)
			}
			return strings.Join(parts, " ")
		},
		"formatTime": func(seconds float64) string {
			totalSeconds := int(seconds)
			hours := totalSeconds / 3600
			minutes := (totalSeconds % 3600) / 60
			secs := totalSeconds % 60
			return fmt.Sprintf("%0.2dh %0.2dm %0.2ds", hours, minutes, secs)
		},
		"formatDate": func(dateStr string) string {
			t, err := time.Parse(time.RFC3339, dateStr)
			if err != nil {
				return dateStr
			}
			return t.Format("2006-01-02 15:04")
		},
		"getRankName": func(rank float64) string {
			rankNames := map[int]string{
				1:  "Private I",
				2:  "Private II",
				3:  "Private III",
				4:  "Specialist I",
				5:  "Specialist II",
				6:  "Specialist III",
				7:  "Corporal I",
				8:  "Corporal II",
				9:  "Corporal III",
				10: "Sergeant I",
				11: "Sergeant II",
				12: "Sergeant III",
				13: "Staff Sergeant I",
				14: "Staff Sergeant II",
				15: "Staff Sergeant III",
				16: "Master Sergeant I",
				17: "Master Sergeant II",
				18: "Master Sergeant III",
				19: "First Sergeant I",
				20: "First Sergeant II",
				21: "First Sergeant III",
				22: "Warrant Officer I",
				23: "Warrant Officer II",
				24: "Warrant Officer III",
				25: "Chief Warrant Officer I",
				26: "Chief Warrant Officer II",
				27: "Chief Warrant Officer III",
				28: "Second Lieutenant I",
				29: "Second Lieutenant II",
				30: "Second Lieutenant III",
				31: "First Lieutenant I",
				32: "First Lieutenant II",
				33: "First Lieutenant III",
				34: "Captain I",
				35: "Captain II",
				36: "Captain III",
				37: "Major I",
				38: "Major II",
				39: "Major III",
				40: "Lieutenant Colonel I",
				41: "Lieutenant Colonel II",
				42: "Lieutenant Colonel III",
				43: "Colonel I",
				44: "Colonel II",
				45: "Colonel III",
				46: "Brigadier General I",
				47: "Brigadier General II",
				48: "Brigadier General III",
				49: "General",
				50: "General of the Army",
			}
			if name, ok := rankNames[int(rank)]; ok {
				return strings.ToUpper(name)
			}
			return fmt.Sprintf("%.0f", rank)
		},
		"toUpper": strings.ToUpper,
	}

	// Load templates
	templates := template.Must(template.New("").Funcs(funcMap).ParseGlob("public/views/*.html"))
	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// Serve static files
	e.Static("/static", "public/static")

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/stats_pc/:name", func(c *echo.Context) error {
		name := c.Param("name")
		url := "https://api.battlefield.rip/archive/bfbc2/players/pc/" + name + "/stats?key=all"
		resp, err := http.Get(url)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error fetching stats")
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return c.String(http.StatusNotFound, "Player not found")
		}
		var stats PlayerStats
		if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
			return c.String(http.StatusInternalServerError, "Error parsing stats")
		}
		return c.Render(http.StatusOK, "stats.html", stats)
	})

	e.POST("/stats_pc", func(c *echo.Context) error {
		name := c.FormValue("searchplayer[name]")
		if name == "" {
			return c.String(http.StatusBadRequest, "Player name is required")
		}
		return c.Redirect(http.StatusFound, "/stats_pc/"+name)
	})

	e.GET("/stats_xbox360/:name", func(c *echo.Context) error {
		name := c.Param("name")
		url := "https://api.battlefield.rip/archive/bfbc2/players/xbox360/" + name + "/stats?key=all"
		resp, err := http.Get(url)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error fetching stats")
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return c.String(http.StatusNotFound, "Player not found")
		}
		var stats PlayerStats
		if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
			return c.String(http.StatusInternalServerError, "Error parsing stats")
		}
		return c.Render(http.StatusOK, "stats.html", stats)
	})

	e.POST("/stats_xbox360", func(c *echo.Context) error {
		name := c.FormValue("searchplayer[name]")
		if name == "" {
			return c.String(http.StatusBadRequest, "Player name is required")
		}
		return c.Redirect(http.StatusFound, "/stats_xbox360/"+name)
	})

	e.GET("/stats_ps3/:name", func(c *echo.Context) error {
		name := c.Param("name")
		url := "https://api.battlefield.rip/archive/bfbc2/players/ps3/" + name + "/stats?key=all"
		resp, err := http.Get(url)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error fetching stats")
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return c.String(http.StatusNotFound, "Player not found")
		}
		var stats PlayerStats
		if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
			return c.String(http.StatusInternalServerError, "Error parsing stats")
		}
		return c.Render(http.StatusOK, "stats.html", stats)
	})

	e.POST("/stats_ps3", func(c *echo.Context) error {
		name := c.FormValue("searchplayer[name]")
		if name == "" {
			return c.String(http.StatusBadRequest, "Player name is required")
		}
		return c.Redirect(http.StatusFound, "/stats_ps3/"+name)
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
