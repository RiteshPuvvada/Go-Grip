package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
) 

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Tittle string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname`
}

var movies []Movie

func getMovies (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies {
		if items.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}


func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "145", Tittle: "Test", Director: &Director{FirstName: "J0hn", LastName: "charlie"}})
	movies = append(movies, Movie{ID: "2", Isbn: "146", Tittle: "Test 2", Director: &Director{FirstName: "olive", LastName: "richards"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")
	
	fmt.Printf("Starting the server at port 3000\n")
	log.Fatal(http.ListenAndServe(":3000",r))

}