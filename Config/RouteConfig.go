package Config

import (
	"github.com/gorilla/mux"
	controller "ExampleRESTApi/Controllers"
	"net/http"
)

func InitializeRoutes(router *mux.Router){
	//Configuring Controller's Routes
	router.HandleFunc("/people", controller.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", controller.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", controller.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", controller.DeletePersonEndpoint).Methods("DELETE")
	//Configuring Static Page's Routes
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./Views/")))
}
