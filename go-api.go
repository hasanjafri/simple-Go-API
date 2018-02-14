package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City     string `json:"city,omitempty"`
	Province string `json:"state,omitempty"`
}

var users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()
	users = append(users, User{ID: "1", Firstname: "Hasan", Lastname: "Jafri", Address: &Address{City: "Montreal", Province: "Quebec"}})
	users = append(users, User{ID: "2", Firstname: "Sultan", Lastname: "Moni", Address: &Address{City: "Peterborough", Province: "Ontario"}})
	users = append(users, User{ID: "3", Firstname: "Dion", Lastname: "Owen"})
	router.HandleFunc("/users", GetUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
