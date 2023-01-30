package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// return created user id with success message, like this: "user created with id: 1"
	responses.JSON(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}

// Path: appil_sns\src\router\routes\user.go
