package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg"
)

type userRouter struct {
	userService root.UserService
}

func NewUserRouter(u root.UserService, router *mux.Router) *mux.Router {
	userRouter := userRouter{u}

	router.HandleFunc("/", userRouter.createUserHandler).Methods("PUT")
	router.HandleFunc("/{username}", userRouter.getUserHandler).Methods("GET")
	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.userService.Create(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	JSON(w, http.StatusOK, err)
}

func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	username := vars["username"]

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	JSON(w, http.StatusOK, user)
}

func (ur *userRouter) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandler")
	err, credentials := decodeCredentials(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user root.User
	err, user = ur.userService.Login(credentials)
	if err == nil {
		JSON(w, http.StatusOK, user)
	} else {
		Error(w, http.StatusInternalServerError, "Incorrect password")
	}
}

func decodeUser(ur *http.Request) (root.User, error) {
	var u root.User
	if ur.Body == nil {
		return u, errors.New("no request body")
	}
	decoder := json.NewDecoder(ur.Body)
	err := decoder.Decode(&u)
	return u, err
}

func decodeCredentials(r *http.Request) (error, root.Credentials) {
	var c root.Credentials
	if r.Body == nil {
		return errors.New("no request body"), c
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	return err, c
}
