package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	model "ExampleRESTApi/Models"
	db "ExampleRESTApi/DataSource"
	config "ExampleRESTApi/Config"
	"os"
	"fmt"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	config.InitializeRoutes(router) //Initializing The Routes Of Controllers

	db.GetInstance().People = append(db.GetInstance().People, model.Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &model.Address{City: "Dublin", State: "CA"}})
	db.GetInstance().People = append(db.GetInstance().People, model.Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})

	fmt.Printf("Server is running in port %s...", addr)

	log.Fatal(http.ListenAndServe(addr, router)) //Listening the HTTP Requests

}
