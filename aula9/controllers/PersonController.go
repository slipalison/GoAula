package personcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPeople busca as pessoas
func GetPeople(w http.ResponseWriter, r *http.Request) {
	montLista()
	json.NewEncoder(w).Encode(people)
}

// GetPerson busca
func GetPerson(w http.ResponseWriter, r *http.Request) {
	montLista()
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	montLista()
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	montLista()
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func montLista() {

	if len(people) < 1 {

		people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
		people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
		people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	}

}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person
