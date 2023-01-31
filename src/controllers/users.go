package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/octoper/go-ray"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)

	_, err = repository.Create(user)

	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// return created user id with success message, like this: "user created with id: 1"
	responses.JSON(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ray.Ray(params)

	fmt.Sprintln(params)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	user, err := repository.GetByID(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.Search(nameOrNick)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}

// Path: appil_sns\src\router\routes\user.go
