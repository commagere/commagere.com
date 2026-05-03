package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	// Serve static files from the static directory - only enable this if you aren't debugging via the dev_appserver.py file
	// http.Handle("/static/", http.FileServer(http.Dir(".")))

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
		"templates/partials/json_ld.html",
		"templates/partials/footer.html",
		"templates/partials/navigation.html",
	}

	templateFilenames := append([]string{layoutFilename, pageFilename}, partialsFilenames...)

	tmpl := template.Must(template.ParseFiles(templateFilenames...))

	err := tmpl.ExecuteTemplate(w, layout, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
