package funcs

import (
	"html/template"
	"net/http"
)

// Home handles requests for the home page.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, 404)
		// http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	fetchArtists()
	SearchForAllLocations()

	Artis.SearchArt = Artis.Artists

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	

	errr := tmpl.Execute(w, Artis)
	if errr != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
