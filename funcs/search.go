package funcs

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	Searchh := r.FormValue("SeArch")

	var ArtTest []ArtistData




	var relationsData struct {
		Index []struct {
			ID             int                 `json:"id"`
			DatesLocations map[string][]string `json:"datesLocations"`
		} `json:"index"`
	}

	
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
	
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&relationsData)
	if err != nil {

		return
	}


	for index, artist := range Artists {
		found := false

		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(Searchh)) ||
			strings.Contains(strconv.Itoa(artist.CreationDate), Searchh) ||
			strings.Contains(artist.FirstAlbum, Searchh) {
				
			ArtTest = append(ArtTest, Artists[index])
			continue
		}

		locations := relationsData.Index[index].DatesLocations

		for location := range locations {
			if strings.Contains(location, Searchh) {
				ArtTest = append(ArtTest, Artists[index])
				found = true
				break
			}
		}

		if found {
			continue
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(Searchh)) {
				ArtTest = append(ArtTest, Artists[index])
				break
			}
		}
	}

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	errr := tmpl.Execute(w, ArtTest)
	if errr != nil {
		http.Error(w, "Err in Execute", http.StatusInternalServerError)
		return
	}
}
