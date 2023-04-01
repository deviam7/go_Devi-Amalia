package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title string `json:"Title"`
	Year  string `json:"Year"`
	Plot  string `json:"Plot"`
}

func handleMovie(w http.ResponseWriter, r *http.Request) {
	imdbID := mux.Vars(r)["imdbID"]
	url := fmt.Sprintf("https://www.omdbapi.com/?apikey=a818034f&i=%s", imdbID)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Akan return data movie utk The True Story of Jesse James")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{imdbID}", handleMovie).Methods("GET")
	http.ListenAndServe(":8080", r)
}
