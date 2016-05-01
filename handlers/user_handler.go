package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"../models"
	"../models/message"
	"../services"
	"../utils"
)

type userHandler struct {
	service *services.UserService
}

func NewUserHandler(db *sqlx.DB) *userHandler {
	return &userHandler{services.NewUserService(db)}
}

// get all users
func (this *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	users, err := this.service.GetUsers().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, users)
}

// get user by id
func (this *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	// get user by id
	user, err := this.service.GetUser(id)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	json, err := user.ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, json)
}

// update user
func (this *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, http.StatusText(400)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user models.User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	err = this.service.UpdateUser(user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// delete user
func (this *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	err = this.service.DeleteUserById(id)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// insert new user into database
func (this *userHandler) InsertUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, http.StatusText(400)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user models.User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	_, err = this.service.InsertUser(user)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// return message as json string
func ErrorMessage(status int, msg string) string {
	msg_final := &message.ResponseMessage{status, msg, "/docs/api/errors"}
	result, _ := utils.NewResultTransformer(msg_final).ToJson()
	return result
}
