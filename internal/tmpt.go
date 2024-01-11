package internal

import (
	"fmt"
	"html/template"
	"path/filepath"

	internalTpl "training/webgo/internal/templates"
)

var templates = make(map[string]*template.Template)

func GetFSTemplate(name string) *template.Template {
	if cachedTpl, ok := templates[name]; ok {
		return cachedTpl
	}

	tplPaths := []string{fmt.Sprintf("%s.gohtml", "layout"), fmt.Sprintf("widgets/%s.gohtml", "navbar"), filepath.Join("pages", fmt.Sprintf("%s.gohtml", name))}
	// tpl, err := template.ParseFS(internalTpl.FS, tplPaths...)

	tpl := template.Must(template.ParseFS(internalTpl.FS, tplPaths...))

	// if err != nil {
	// 	panic(err)
	// }

	templates[name] = tpl

	return tpl
}

// func GetTemplate(name string) *template.Template {
// 	if cachedTpl, ok := templates[name]; ok {
// 		return cachedTpl
// 	}
// 	// If Widgets or blocks have its own folder, they can be included in the parser by collected
// 	// the files name
// 	tplPaths := []string{filepath.Join("templates", fmt.Sprintf("%s.gohtml", "layout")), filepath.Join("template/pages", fmt.Sprintf("%s.gohtml", name))}

// 	templates[name] = template.Must(template.ParseFiles(tplPaths...))

// 	return templates[name]
// }
