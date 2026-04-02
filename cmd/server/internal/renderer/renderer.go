package renderer

import (
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v5"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(glob string) (*TemplateRenderer, error) {
	tmpl, err := template.New("bfbcs").
		Funcs(template.FuncMap{
			"add":                add,
			"mul":                mul,
			"div":                div,
			"gt":                 gt,
			"formatNumber":       formatNumber,
			"formatDuration":     formatDuration,
			"formatTime":         formatTime,
			"getRankName":        getRankName,
			"getWeaponStarClass": getWeaponStarClass,
			"toUpper":            strings.ToUpper,
		}).
		ParseGlob(glob)
	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{
		templates: tmpl,
	}, nil
}

func (r *TemplateRenderer) Render(_ *echo.Context, w io.Writer, name string, data any) error {
	return r.templates.ExecuteTemplate(w, name, data)
}
