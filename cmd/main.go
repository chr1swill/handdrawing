package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type MainLayoutData struct {
	Title       string
	Description string
	Keywords    string
	BodyContent template.HTML
}

var tmpl *template.Template

func loadTemplates(paths []string) {
	tmpl = template.New("")

	for _, pattern := range paths {

		matches, err := filepath.Glob(pattern)
		if err != nil {
			log.Printf("Failed to glob pattern %s: %v", pattern, err)
			continue
		}

		if matches == nil {
			log.Printf("No templates matched the pattern %s: %v", pattern, err)
			continue
		}

		for _, match := range matches {
			_, err := tmpl.ParseFiles(match)
			if err != nil {
				log.Fatalf("Failed to parse template %s: %v", match, err)
			}
		}
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var bodyContent bytes.Buffer

	err := tmpl.ExecuteTemplate(&bodyContent, "home.html", nil)
	if err != nil {
		http.Error(w, "Failed to render content", http.StatusInternalServerError)
		return
	}

	homePageData := MainLayoutData{
		Title:       "Hand Drawing App",
		Description: "This is my hand drawing app",
		Keywords:    strings.Join([]string{"notebook", "hand draw", "drawings", "digital art"}, ","),
		BodyContent: template.HTML(bodyContent.String()),
	}

	err = tmpl.ExecuteTemplate(w, "main-layout.html", homePageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	rootDir := filepath.Join(".")

	templateDir := filepath.Join(rootDir, "templates")
	viewFiles := filepath.Join(templateDir, "views", "*.html")
	partialFiles := filepath.Join(templateDir, "partials", "*.html")
	layoutFiles := filepath.Join(templateDir, "layouts", "*.html")

	paths := []string{viewFiles, partialFiles, layoutFiles}
	loadTemplates(paths)

	staticFilesPath := filepath.Join(rootDir, "static")
	fs := http.FileServer(http.Dir(staticFilesPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePageHandler)

	fmt.Printf("Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
