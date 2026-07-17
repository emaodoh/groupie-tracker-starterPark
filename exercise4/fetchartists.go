package main

import(
	"encoding/json"
	"net/http"
)


type Artists struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}


func fetchArtists()([]Artists, error){

	url := "https://groupietrackers.herokuapp.com/api/artists"
	response,err := http.Get(url)

	if err != nil{
		return nil,err
	}

	defer response.Body.Close()
		var artists []Artists

	if response.StatusCode == http.StatusOK{


		if err = json.NewDecoder(response.Body).Decode(&artists); err != nil{
			return nil,err
		}
	}

	return artists,nil
}