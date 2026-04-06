package renderer

import (
	"strings"
)

type PageContext map[string]any

type Setter func(c PageContext)

func WithTitle(title string) Setter {
	return func(c PageContext) {
		c["Title"] = title
	}
}

func WithPlatform(platform string) Setter {
	return func(c PageContext) {
		c["Platform"] = strings.ToLower(platform)
	}
}

func With(key string, value any) Setter {
	return func(c PageContext) {
		c[key] = value
	}
}

func NewPageContext(setters ...Setter) PageContext {
	c := make(PageContext, len(setters))
	for _, setter := range setters {
		setter(c)
	}

	return c
}
