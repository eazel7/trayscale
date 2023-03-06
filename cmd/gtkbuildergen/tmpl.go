package main

import (
	"embed"
	_ "embed"
	"text/template"

	"deedles.dev/trayscale/internal/set"
)

var (
	//go:embed *.tmpl
	tmplFS embed.FS

	tmpl *template.Template
)

func init() {
	tmpl = template.New("")
	tmpl.Funcs(map[string]any{
		"requires": func(ui []Interface) []Requires {
			// TODO: Return better values.
			return ui[0].Requires
		},
		"hasMenus": func(ui []Interface) bool {
			for _, i := range ui {
				if len(i.Menus) != 0 {
					return true
				}
			}
			return false
		},
		"newValueSet": func() set.Set[string] {
			return make(set.Set[string])
		},
	})
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "*.tmpl"))
}