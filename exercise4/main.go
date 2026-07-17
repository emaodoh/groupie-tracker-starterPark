package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



var store []Artists

func main() {
	var err error
	store, err = fetchArtists()

	if err != nil {
		log.Fatal(err)
	}
// Serve static files from the static directory using http.FileServer
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Artist loaded successfully")
	http.HandleFunc("/", homeHandler)
	fmt.Println("server is live")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("static/home.html")

	if err != nil {
		log.Println(err)

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	

	tmpl.Execute(w, store)
}
