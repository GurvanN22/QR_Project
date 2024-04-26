package tools

// RenderTemplate & TemplateCache

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	var myCache, err = createTemplateCache()
	if err != nil {
		fmt.Println(err)
	}

	t, ok := myCache[tmpl]
	if !ok {
		fmt.Println("Could not get template from cache")
	}

	buffer := new(bytes.Buffer)
	t.Execute(buffer, p)
	buffer.WriteTo(w)
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("Template/page/*.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("Template/layout/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseFiles(matches...)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
