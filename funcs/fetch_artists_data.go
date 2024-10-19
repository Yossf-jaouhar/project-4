package funcs

import (
	"encoding/json"
	"net/http"
)

func fetchArtists(Url string)(Artists, error){
	resp, err := http.Get(Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists Artists
	
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}



func fetchArtist(apiURL string) (ArtistData, error) {
	var artists ArtistData
	resp, err := http.Get(apiURL)
	if err != nil {
		return artists, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return artists, err
	}
	return artists, nil
}


func fetchRelations(artist ArtistData) (ArtistData, error) {
	resp, err := http.Get(artist.Relations)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&artist.RelationData)
	if err != nil {
		return artist, err
	}
	return artist, nil
}
func fetchLocation(artist ArtistData) (ArtistData, error) {
	resp, err := http.Get(artist.Locations)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artist.LocationsDATA)
	if err != nil {
		return artist, err
	}
	return artist, nil
} 


func fetchDates(artist ArtistData) (ArtistData, error) {
	resp, err := http.Get(artist.ConcertDates)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artist.ConcertDatesDATA)
	if err != nil {
		return artist, err
	}
	return artist, nil
}