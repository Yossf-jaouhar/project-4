package funcs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FetchArtists(url string)(any, error) {
	Response, err := http.Get(url)
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



