package funcs

import (
	"net/http"
	"os"
)

// CssHandler handles requests for CSS files in the "assets" directory.
func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		files, err := os.ReadDir("assets")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		for _, file := range files {
			if r.URL.Path == "/assets/"+file.Name() {

				fs := http.Dir("assets")
				http.StripPrefix("/assets", http.FileServer(fs)).ServeHTTP(w, r)
				return
			}
		}
		ErrorHandler(w, http.StatusNotFound)
		return
	} else {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
