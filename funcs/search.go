package funcs

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Search handles search requests for artists.
func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	Searchh := r.FormValue("SeArch")
	Searchh = strings.ToLower(Searchh)
	var ArtTest []ArtistData

	for index, artist := range Artis.Artists {
		found := false

		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(Searchh)) ||
			strings.Contains(strconv.Itoa(artist.CreationDate), Searchh) ||
			artist.FirstAlbum == Searchh {

			ArtTest = append(ArtTest, Artis.Artists[index])
			continue
		}
		locations := Artis.Locat.Index[index].Locations
		for _, location := range locations {
			if strings.Contains(location, Searchh) {
				ArtTest = append(ArtTest, Artis.Artists[index])
				found = true
				break
			}
		}

		if found {
			continue
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(Searchh)) {
				ArtTest = append(ArtTest, Artis.Artists[index])
				break
			}
		}
	}

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	Artis.SearchArt = ArtTest
	errr := tmpl.Execute(w, Artis)
	if errr != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
