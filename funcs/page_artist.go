package funcs

import (
	"html/template"
	"net/http"
	"strconv"
)

func PageArtist(w http.ResponseWriter, r *http.Request) {
	
	id := r.PathValue("id")

	IDD, _ := strconv.Atoi(id)

	if IDD <= 0 || IDD >= 53 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Artists[IDD-1].FeatchAll()


	tmpl, err := template.ParseFiles("Templates/info.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	errr := tmpl.Execute(w,Artists[IDD-1])
	if errr != nil {
		http.Error(w, "Err in Execute", http.StatusInternalServerError)
		return
	}
}
