package funcs

import (
	"html/template"
	"net/http"
)



func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	
	artists, err := FetchArtists("https://groupietrackers.herokuapp.com/api/artists")
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
