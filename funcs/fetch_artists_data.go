package funcs

import (
	"encoding/json"
	"net/http"
)

func fetchArtists() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		// err
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&Artists)
	if err != nil {
		// errr
		return
	}
}

func (a *ArtistData) fetchRelations() {
	resp, err := http.Get(a.Relations)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&a.RelationData)
	if err != nil {
		return
	}
}

func (a *ArtistData) fetchLocation() {
	resp, err := http.Get(a.Locations)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&a.LocationsDATA)
	if err != nil {
		return
	}
}

func (a *ArtistData) fetchDates() {
	resp, err := http.Get(a.ConcertDates)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&a.ConcertDatesDATA)
	if err != nil {
		return
	}
}

func (a *ArtistData) FeatchAll() {
	a.fetchRelations()
	a.fetchDates()
	a.fetchLocation()
}
