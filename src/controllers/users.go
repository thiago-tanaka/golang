package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("create"); err != nil {
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
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("you can't update a user that is not yours"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("edit"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	if err = repository.Update(userID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("you can't delete a user that is not yours"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	if err = repository.Delete(userID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followedId, err := auth.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	followedId, err = strconv.ParseUint(params["id"], 10, 64)

	tokenUserId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if followedId == tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	if err = repository.Follow(tokenUserId, followedId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followedId, err := auth.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	followedId, err = strconv.ParseUint(params["id"], 10, 64)

	tokenUserId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if followedId == tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("you can't unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	if err = repository.Unfollow(tokenUserId, followedId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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
	followers, err := repository.GetFollowers(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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
	following, err := repository.GetFollowing(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

// Path: appil_sns\src\router\routes\user.go
