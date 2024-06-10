package main

import (
	"log"
	"net/http"

	"github.com/quienn/gosuarios/internal/gosuarios/controllers"
	"github.com/quienn/gosuarios/internal/gosuarios/models"
)

func main() {
	models.ConnectDatabase()

	http.HandleFunc("POST /users", controllers.CreateUser)
	http.HandleFunc("GET /users", controllers.FindUsers)
	http.HandleFunc("GET /users/{id}", controllers.FindUser)
	http.HandleFunc("PATCH /users/{id}", controllers.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
