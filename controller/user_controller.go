package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"

	"../model"
	"../model/message"
	"../service"
	"../util"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(db *sqlx.DB) *UserController {
	return &UserController{service.NewUserService(db)}
}

// ========================
// get all users
// ========================
func (this *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

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

// ========================
// get user by id
// ========================
func (this *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)

	// get user by id
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

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

// ========================
// insert new user into database
// ========================
func (this *UserController) InsertUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, http.StatusText(400)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user model.User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, err.Error()))
		return
	}

	id, err := this.service.InsertUser(user)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, ErrorMessage(500, err.Error()))
		return
	}

	w.Header().Set("Location", util.BaseURL(r) + r.RequestURI + "/" + strconv.FormatInt(id, 10))
	w.WriteHeader(200)
}

// ========================
// update user
// ========================
func (this *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprint(w, ErrorMessage(400, http.StatusText(400)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user model.User

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

	w.Header().Set("Location",  util.BaseURL(r) + r.RequestURI + "/" + strconv.FormatInt(user.Id, 10))
	w.WriteHeader(200)
}

// ========================
// delete user
// ========================
func (this *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

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

	w.Header().Set("Location", util.BaseURL(r) + "/api/v2/users")
	w.WriteHeader(200)
}

// ========================
// return message as json string
// ========================
func ErrorMessage(status int, msg string) string {
	msg_final := &message.ResponseMessage{status, msg, "/docs/api/errors"}
	result, _ := util.NewResultTransformer(msg_final).ToJson()
	return result
}
