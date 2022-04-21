package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Movie struct {
	Name string `json:"name"`
	Genre string `json:"genre"`
	description	string `json:"description"`
}

var allMovies = []Movie{
	{
		Name: "The Shawshank Redemption",
		Genre: "Drama",
		description: "Two imprisoned",
	},
	{
		Name: "The Godfather",
		Genre: "Drama",
		description: "The aging patriarch",
	},
	{
		Name: "The Godfather: Part II",
		Genre: "Drama",
		description: "The early life and career of",
	},
} 

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("movies handler", r.Method, r.URL)
	// if the request is GET Method
	if r.Method == http.MethodGet {
		var q = r.URL.Query()
		f := q.Get("filter")

		// no filter
		if (f == "") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(allMovies)
			return
		}

		// has filter
		var movies []Movie
		for _, movie := range allMovies {
			if strings.Contains(movie.Name, f) {
				movies = append(movies, movie)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	} else {
		msg := http.StatusText(http.StatusMethodNotAllowed)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/movies", moviesHandler)

	port := "8080"
	log.Println("starting server at port:", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}