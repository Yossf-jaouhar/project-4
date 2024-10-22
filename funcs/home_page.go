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

	fetchArtists()
	
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	errr := tmpl.Execute(w, Artists)
	if errr != nil {
		http.Error(w, "Err in Execute", http.StatusInternalServerError)
		return
	}
	
}
