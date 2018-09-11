package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/solairerove/go-simple-mux/pkg/entity"
)

var people []entity.Person

// GetPeople ...
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}

func main() {
	people = append(people, entity.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &entity.Address{City: "City X", State: "State X"}})
	people = append(people, entity.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &entity.Address{City: "City Z", State: "State Y"}})
	people = append(people, entity.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
