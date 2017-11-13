package Config

import (
	"github.com/gorilla/mux"
	controller "../Controllers"
)

func InitializeRoutes(router *mux.Router){
	router.HandleFunc("/people", controller.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", controller.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", controller.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", controller.DeletePersonEndpoint).Methods("DELETE")
}
