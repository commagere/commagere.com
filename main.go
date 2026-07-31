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

	// We want to redirect all requests to HTTPS and remove the www prefix if present.
	// This is currently handled by Cloudflare, but we can also handle it here for local development or if Cloudflare is not used.
	// // Get the host and protocol from the request
	// host := r.Host
	// proto := r.Header.Get("X-Forwarded-Proto")

	// // Redirect to HTTPS and remove www prefix
	// if strings.HasPrefix(host, "www.") || proto == "http" {
	// 	cleanHost := strings.TrimPrefix(host, "www.")
	// 	http.Redirect(w, r, "https://"+cleanHost+r.URL.Path, http.StatusMovedPermanently)
	// 	return
	// }

	// Handle 404 for any path other than "/"
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
