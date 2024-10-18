package funcs

import (
	"html/template"
	"net/http"
	"strconv"
)

func PageArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	IDD, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if IDD <= 0 || IDD >= 53 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	DATA, err := FetchArtistsData("https://groupietrackers.herokuapp.com/api/artists")

	Tmpl, err := template.ParseFiles("info.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = Tmpl.Execute(w, DATA)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
