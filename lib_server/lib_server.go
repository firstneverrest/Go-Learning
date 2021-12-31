package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

func moviesHandler(c echo.Context) error {
	var f string = c.QueryParam("filter")
	if f == "" {
		return c.JSON(http.StatusOK, allMovies)
	}

	var movies []Movie
	for _, movie := range allMovies {
		if strings.Contains(movie.Name, f) {
			movies = append(movies, movie)
		}
	}

	return c.JSON(http.StatusOK, movies)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/movies", moviesHandler)

	port := "8080"
	log.Println("starting server at port:", port)
	log.Fatal(e.Start(":" + port))
}