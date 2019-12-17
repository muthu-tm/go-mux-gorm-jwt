package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"go-rest-examples/db"
	"go-rest-examples/routes"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("unable to load .env file. WARNING!")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting environment values, %v", err)
	} else {
		fmt.Println("We are getting the environment values")
	}

	db.Initialize(os.Getenv("driver"), os.Getenv("user"), os.Getenv("password"), os.Getenv("port"), os.Getenv("host"), os.Getenv("name"))
	db.SeedData(db.GetDB())

	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	fmt.Println("Listening to port 8080")
	err = http.ListenAndServe(":8080", router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
