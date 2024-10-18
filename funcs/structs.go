package funcs

type Artist struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	FirstAlbum string `json:"firstAlbum"`
}

type ArtistData struct {
	Id               int      `json:"id"`
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
	RelationsDATA    RelationsS
}

type LocationsS struct {
	Data []string
}

type ConcertDatesS struct {
	Data []string
}

type RelationsS struct {
	Data map[string]string
}
