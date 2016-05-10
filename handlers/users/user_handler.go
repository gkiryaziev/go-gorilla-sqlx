package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUsers return all users
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	users, err := uh.getUsers().ToJSON()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, errorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, users)
}

// GetUser return user by id
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, err.Error()))
		return
	}

	// get user by id
	user, err := uh.getUser(id)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, err.Error()))
		return
	}

	json, err := user.ToJSON()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, errorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, json)
}

// UpdateUser update user
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, http.StatusText(400)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, err.Error()))
		return
	}

	err = uh.updateUser(user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// DeleteUser delete user
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, err.Error()))
		return
	}

	err = uh.deleteUserByID(id)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, err.Error()))
		return
	}

	w.WriteHeader(200)
}

// InsertUser insert new user into database
func (uh *UserHandler) InsertUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, http.StatusText(400)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errorMessage(400, err.Error()))
		return
	}

	_, err = uh.insertUser(user)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, errorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
}
