package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/quienn/gosuarios/internal/gosuarios/models"
)

// TODO llenar los métodos del controlador

// GET /users
func FindUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	models.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func FindUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

type UserInput struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input UserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := models.User{Nickname: input.Nickname, Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// PATCH /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var userUpdate UserInput
	json.NewDecoder(r.Body).Decode(&userUpdate)

	models.DB.Model(&user).Updates(userUpdate)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	models.DB.Delete(&user)

	w.WriteHeader(http.StatusNoContent)
}
