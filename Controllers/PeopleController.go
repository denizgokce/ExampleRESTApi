package Controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	db "ExampleRESTApi/DataSource"
	model "ExampleRESTApi/Models"
)

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person = db.GetPerson(params["id"])
	//json.NewEncoder(w).Encode(&model.Person{})
	json.NewEncoder(w).Encode(person)
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(db.GetPeople())
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person model.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	db.AddPerson(&person)
	json.NewEncoder(w).Encode(db.GetPeople())
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	db.DeletePerson(params["id"])
	json.NewEncoder(w).Encode(db.GetPeople())
}

func UpdatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person model.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	db.UpdatePerson(params["id"], &person)
	json.NewEncoder(w).Encode(db.GetPeople())
}
