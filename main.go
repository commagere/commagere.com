package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	layout := "base"
	page := "index"

	layoutFilename := "templates/layouts/" + layout + ".html"
	pageFilename := "templates/pages/" + page + ".html"

	partialsFilenames := []string{
		"templates/partials/header.html",
		"templates/partials/footer.html",
		"templates/partials/navigation.html",
	}

	templateFilenames := append([]string{layoutFilename, pageFilename}, partialsFilenames...)

	tmpl := template.Must(template.ParseFiles(templateFilenames...))

	//tmpl := template.Must(template.ParseFiles(layoutFilename, pageFilename, partialsFilenames...))
	// tmpl := template.Must(template.ParseGlob("templates/index.html"))
	// err := tmpl.Execute(w, nil)
	err := tmpl.ExecuteTemplate(w, layout, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
