package funcs

import (
	"encoding/json"
	"net/http"
)

func FetchArtistsData(ArtistUrl string) ([]string, error) {

	var ArtistInfo []ArtistData
	Response , err :=  http.Get(ArtistUrl)
	if err != nil {
		return nil , err
	}

	err = json.NewDecoder(Response.Body).Decode(&ArtistInfo)
	if err != nil {
		return nil , err
	}

	return nil, nil
}
