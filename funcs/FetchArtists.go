package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchArtists() ([]Artist, error) {
	Response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer Response.Body.Close()

	Body, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		return nil, err
	}

	var artists []Artist

	err = json.Unmarshal(Body, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func FetchArtistByName(name string) (Artist, error) {
	artistsList, err := FetchArtists()
	if err != nil {
		
		return Artist{}, err
	}
	for _, artist := range artistsList {
		if artist.Name == name {
			return artist, nil
		}
	}
	return Artist{}, fmt.Errorf("artist not found")
}
