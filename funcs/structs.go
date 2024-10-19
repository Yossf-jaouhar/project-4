package funcs

type ArtistData struct {
	ID              int      `json:"id"`
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
	RelationData    RelationsS
}


type Artists []ArtistData


type LocationsS struct {
	ID        int      `json:"id"`
	Data []string      `json:"locations"`
}

type ConcertDatesS struct {
	ID        int      `json:"id"`
	Data []string         `json:"dates"`
}

type RelationsS struct {
	ID        int      `json:"id"`
	Data map[string][]string  `json:"datesLocations"`
}
