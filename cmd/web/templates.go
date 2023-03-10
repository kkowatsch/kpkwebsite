package main

import (
	"embed"
	"html/template"
	"io/fs"
	"path/filepath"
)

var Files embed.FS

type templateData struct {
	Flash string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(Files, "ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"ui/html/base.tmpl",
			"ui/html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).ParseFS(Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
