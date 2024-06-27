package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type FaviconTagData struct {
	Rel   template.URL
	Sizes string
	Type  template.URL
	Color string
	Href  template.URL
}

type ScriptTagData struct {
	Src template.URL
}

type MainLayoutData struct {
	Title       string
	Description string
	Keywords    []string
	OgImage     template.URL
	OgUrl       template.URL
	FaviconTags []FaviconTagData
	ScriptTags  []ScriptTagData
	BodyContent template.HTML
}

func main() {
	rootDir := filepath.Join(".")

	templateDir := filepath.Join(rootDir, "templates")
	viewDir := filepath.Join(templateDir, "views")
	partialDir := filepath.Join(templateDir, "partials")
	layoutDir := filepath.Join(templateDir, "layouts")

	staticDir := filepath.Join(rootDir, "static")
	cssDir := filepath.Join(staticDir, "css")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}


	})
}
