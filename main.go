package main

import (
	"encoding/json" // Perbaikan: Penulisan "encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string   `json:"id"`
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) { // Perbaikan: Penulisan tipe `ResponseWriter` dan `Request`
	w.Header().Set("Content-Type", "application/json") // Perbaikan: Penulisan `Set` huruf besar
	json.NewEncoder(w).Encode(movies) // Perbaikan: Penulisan `NewEncoder` huruf besar
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] { // Perbaikan: `item.id` menjadi `item.ID`
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Perbaikan: Penulisan fungsi `strconv.Itoa(rand.Intn(...))`
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) { // Perbaikan: Posisi `{`
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "23345", Title: "Ambaruwo", Director: &Director{Firstname: "ci", Lastname: "imut"}})
	movies = append(movies, Movie{ID: "2", Isbn: "20345", Title: "Movie 2", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET") // Perbaikan: `HendleFunc` menjadi `HandleFunc`
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // Perbaikan: Tambahkan `/` di route

	fmt.Printf("Starting server at port 8000\n") // Perbaikan: Escape karakter `\n`
	log.Fatal(http.ListenAndServe(":8000", r))
}
