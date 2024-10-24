package funcs

// ArtistData holds information about an artist.
type ArtistData struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Locations        string   `json:"locations"`
	LocationsDATA    LocationsS
	ConcertDates     string `json:"concertDates"`
	ConcertDatesDATA ConcertDatesS
	Relations        string `json:"relations"`
	RelationData     RelationsS
	FEATCHINGerror   error
}

// Artists holds a slice of ArtistData.
var Artis struct {
	Artists []ArtistData
	Locat   LocationsDATA

	SearchArt []ArtistData
}

// LocationsS hoArtistslds location data related to artists.
type LocationsS struct {
	Data []string `json:"locations"`
}

// ConcertDatesS holds concert date data related to artists.
type ConcertDatesS struct {
	Data []string `json:"dates"`
}

// RelationsS holds relation data for artists.
type RelationsS struct {
	Data map[string][]string `json:"datesLocations"`
}

type LocationsDATA struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

