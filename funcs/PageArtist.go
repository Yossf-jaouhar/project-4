package funcs

import (
	"html/template"
	"net/http"
)

func PageArtist(w http.ResponseWriter, r *http.Request) {
	// Handle only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get the artist name from the form data
	artistName := r.FormValue("artistName")
	if artistName == "" {
		http.Error(w, "Artist not found", http.StatusBadRequest)
		return
	}

	// Fetch data about the selected artist
	artist, err := FetchArtistByName(artistName) // You'll need to implement this function to fetch the artist by name
	if err != nil {
		http.Error(w, "Artist not found", http.StatusInternalServerError)
		return
	}

	// Parse and render the artist details page
	tmpl, err := template.ParseFiles("Templates/info.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artist)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
