package main

import (
	"embed"
	"html/template"
	"io/fs"
	"path/filepath"
	"strconv"
	"time"

	"kaykodesigns.kpkaccounting.net/ui"
)

var Files embed.FS

type templateData struct {
	Form any
	Page any
}

func copyrightYear() string {
	return strconv.Itoa(time.Now().Year())
}

var functions = template.FuncMap{
	"copyrightYear": copyrightYear,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
