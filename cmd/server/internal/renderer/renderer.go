package renderer

import (
	"fmt"
	"html/template"
	"io"
	"math"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v5"
)

type TemplateRenderer struct {
	layouts *template.Template
	views   map[string]*template.Template
}

func NewTemplateRenderer(layouts string, views string) (*TemplateRenderer, error) {
	tmpl, err := template.New("bfbcs").
		Funcs(template.FuncMap{
			"seq":                 seq,
			"add":                 add,
			"mul":                 mul,
			"div":                 div,
			"gt":                  gt,
			"gte":                 gte,
			"min":                 math.Min,
			"formatNumber":        formatNumber,
			"formatDuration":      formatDuration,
			"formatTime":          formatTime,
			"getRankName":         getRankName,
			"getWeapons":          getWeapons,
			"getVehicles":         getVehicles,
			"getServiceStarClass": getServiceStarClass,
			"calculateProgress":   calculateProgress,
			"toUpper":             strings.ToUpper,
			"toLower":             strings.ToLower,
			"hasPrefix":           strings.HasPrefix,
		}).
		ParseGlob(layouts)
	if err != nil {
		return nil, err
	}

	matches, err := filepath.Glob(views)
	if err != nil {
		return nil, err
	}

	v := make(map[string]*template.Template, len(matches))
	for _, match := range matches {
		view, err2 := tmpl.Clone()
		if err2 != nil {
			return nil, err2
		}

		view, err2 = view.ParseFiles(match)
		if err2 != nil {
			return nil, err2
		}

		v[filepath.Base(match)] = view
	}

	return &TemplateRenderer{
		layouts: tmpl,
		views:   v,
	}, nil
}

func (r *TemplateRenderer) Render(_ *echo.Context, w io.Writer, name string, data any) error {
	layout, view, _ := strings.Cut(name, "/")

	tmpl, ok := r.views[view]
	if !ok {
		return fmt.Errorf("no such view: %s", view)
	}

	return tmpl.ExecuteTemplate(w, layout, data)
}
