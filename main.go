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

	/*c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"X-Requested-With"},
	})*/

	c := cors.AllowAll()

	handler := c.Handler(router)

	fmt.Printf("Server is running in port %s...", addr)

	log.Fatal(http.ListenAndServe(addr, handler)) //Listening the HTTP Requests

}
