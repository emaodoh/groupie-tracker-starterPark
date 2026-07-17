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

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Relation `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Dates struct {
	Index []Date `json:"index"`
}

type Index struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Locations struct {
	Index []Index `json:"index"`
}

type DataStore struct {
	Artists   []Artist
	Locations []Index
	Dates     []Date
	Relations []Relation
}

var Store DataStore

// --- Fetching Functions ---

func fetchArtists(client *http.Client) ([]Artist, error) {
	resp, err := client.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func fetchLocations(client *http.Client) ([]Index, error) {
	resp, err := client.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var loc Locations
	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
		return nil, err
	}
	return loc.Index, nil
}

func fetchDates(client *http.Client) ([]Date, error) {
	resp, err := client.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() 

	var d Dates
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}
	return d.Index, nil
}

func fetchRelations(client *http.Client) ([]Relation, error) {
	resp, err := client.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rel Relations
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return nil, err
	}
	return rel.Index, nil
}

// Single unified fetchAll function returning errors if any fail
func fetchAll() error {
	client := &http.Client{}

	var err error

	Store.Artists, err = fetchArtists(client)
	if err != nil {
		return fmt.Errorf("failed fetching artists: %v", err)
	}

	Store.Locations, err = fetchLocations(client)
	if err != nil {
		return fmt.Errorf("failed fetching locations: %v", err)
	}

	Store.Dates, err = fetchDates(client)
	if err != nil {
		return fmt.Errorf("failed fetching dates: %v", err)
	}

	Store.Relations, err = fetchRelations(client)
	if err != nil {
		return fmt.Errorf("failed fetching relations: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("Loading data into memory...")
	if err := fetchAll(); err != nil {
		log.Fatalf("Data loading initialization panic: %v", err)
	}

	fmt.Println("Data loaded successfully!")
	fmt.Printf("Total Records Loaded -> Artists: %d, Locations: %d, Dates: %d, Relations: %d\n",
		len(Store.Artists), len(Store.Locations), len(Store.Dates), len(Store.Relations))

	
}
