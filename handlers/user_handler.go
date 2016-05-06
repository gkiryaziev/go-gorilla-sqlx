package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/models"
	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/models/message"
	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/services"
	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/utils"
)

type userHandler struct {
	service *services.UserService
}

// NewUserHandler return new userHandler object
func NewUserHandler(db *sqlx.DB) *userHandler {
	return &userHandler{services.NewUserService(db)}
}

// GetUsers return all users
func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	users, err := uh.service.GetUsers().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, users)
}

// GetUser return user by id
func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	// get user by id
	user, err := uh.service.GetUser(id)
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

// UpdateUser update user
func (uh *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

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

	err = uh.service.UpdateUser(user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// DeleteUser delete user
func (uh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	err = uh.service.DeleteUserById(id)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// InsertUser insert new user into database
func (uh *userHandler) InsertUser(w http.ResponseWriter, r *http.Request) {

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

	_, err = uh.service.InsertUser(user)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// ErrorMessage return error message as json string
func ErrorMessage(status int, msg string) string {
	msg_final := &message.ResponseMessage{status, msg, "/docs/api/errors"}
	result, _ := utils.NewResultTransformer(msg_final).ToJson()
	return result
}
