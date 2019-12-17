package routes

import (
	"github.com/gorilla/mux"
	c "go-rest-examples/controllers"
)

func InitializeRoutes(router *mux.Router) {

	// Login Route
	router.HandleFunc("/login", SetMiddlewareJSON(c.Login)).Methods("POST")

	//Users routes
	router.HandleFunc("/users", SetMiddlewareJSON(c.CreateUser)).Methods("POST")
	router.HandleFunc("/users", SetMiddlewareJSON(c.GetUsers)).Methods("GET")
	router.HandleFunc("/users/{id}", SetMiddlewareJSON(c.GetUser)).Methods("GET")
	router.HandleFunc("/users/{id}", SetMiddlewareJSON(SetMiddlewareAuthentication(c.UpdateUser))).Methods("PUT")
	router.HandleFunc("/users/{id}", SetMiddlewareAuthentication(c.DeleteUser)).Methods("DELETE")

}
