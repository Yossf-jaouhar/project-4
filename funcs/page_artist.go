package funcs

import (
	"html/template"
	"net/http"
	"strconv"
)

// PageArtist handles requests for individual artist pages.
func PageArtist(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	IDD, _ := strconv.Atoi(id)

	if IDD <= 0 || IDD >= 53 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if len(Artis.Artists) != 0 {

		Artis.Artists[IDD-1].FeatchAll()

		tmpl, err := template.ParseFiles("Templates/info.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		errr := tmpl.Execute(w, Artis.Artists[IDD-1])
		if errr != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	} else {
		// http.Redirect(w, r, "/404", http.StatusFound)
		ErrorHandler(w, http.StatusNotFound)
		return
	}
}
