package funcs

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	artists, err := FetchArtists()

	if err != nil {
		http.Error(w, "Failedllllll", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
