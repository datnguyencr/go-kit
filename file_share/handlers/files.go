package handlers

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

func listFiles(dir string) template.HTML {
	var b strings.Builder

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		name := d.Name()
		b.WriteString(`<li><a href="/download/` + name + `">` + name + `</a></li>`)
		return nil
	})

	return template.HTML(b.String())
}
