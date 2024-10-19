package funcs

import (
	"html/template"
	"net/http"
	"strconv"
)

func PageArtist(w http.ResponseWriter, r *http.Request) {
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

	apiURL := "https://groupietrackers.herokuapp.com/api/artists/" + id

	artistData, er := fetchArtist(apiURL)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}

	artistData, err = fetchRelations(artistData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}
	artistData, errr := fetchLocation(artistData)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}
	artistData, errrr := fetchDates(artistData)
	if errrr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}


	tmpl, err := template.ParseFiles("Templates/info.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w,artistData)

	
	if errr != nil {
		http.Error(w, "Err in Execute", http.StatusInternalServerError)
		return
	}
}
