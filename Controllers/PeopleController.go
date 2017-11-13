package Controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	model "../Models"
	db "../DataSource"
)


func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range db.GetInstance().People {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(db.GetInstance().People)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person model.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	db.GetInstance().People = append(db.GetInstance().People, person)
	json.NewEncoder(w).Encode(db.GetInstance().People)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range db.GetInstance().People {
		if item.ID == params["id"] {
			db.GetInstance().People = append(db.GetInstance().People[:index], db.GetInstance().People[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(db.GetInstance().People)
}
