package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		content, err := io.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(content))
	}else{
		log.Printf("server returned an error status code: %d", response.StatusCode)
	}
}
