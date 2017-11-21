package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	config "ExampleRESTApi/Config"
	"os"
	"fmt"
	"github.com/rs/cors"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		//return "", fmt.Errorf("$PORT not set")
		port = "5000"
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

	handler := cors.Default().Handler(router)

	fmt.Printf("Server is running in port %s...", addr)

	log.Fatal(http.ListenAndServe(addr, handler)) //Listening the HTTP Requests

}
