package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
)
type Artist struct {
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

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	fmt.Println("loaded successfully")

	if response.StatusCode == http.StatusOK {
		var artist []Artist

		err = json.NewDecoder(response.Body).Decode(&artist); if err != nil{
			log.Fatal(err)
		}

		fmt.Printf("Total artist: %d\n", len(artist))

		for _,list := range artist{
			fmt.Printf("Name: %s, Date: %d\n", list.Name, list.CreationDate)
		}
	} else {
		log.Printf("server returned an error status code: %d", response.StatusCode)

	}
}
