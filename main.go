package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	model "ExampleRESTApi/Models"
	db "ExampleRESTApi/DataSource"
	config "ExampleRESTApi/Config"
)

func main() {
	router := mux.NewRouter()
	config.InitializeRoutes(router) //Initializing The Routes Of Controllers

	db.GetInstance().People = append(db.GetInstance().People, model.Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &model.Address{City: "Dublin", State: "CA"}})
	db.GetInstance().People = append(db.GetInstance().People, model.Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})


	log.Fatal(http.ListenAndServe(":5000", router)) //Listening the HTTP Requests at Port 5000

}
