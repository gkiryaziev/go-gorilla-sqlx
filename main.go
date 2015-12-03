package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"

	"./controller"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	mx := mux.NewRouter()

	// variables
	db_user_name := "admin"
	db_user_password := "admin"
	db_host := "192.168.2.11"
	db_port := "3306"
	db_database := "db_social"
	http_host := ""
	http_port := "8008"

	// mysql connection string
	mysql_bind := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		db_user_name, db_user_password, db_host, db_port, db_database)

	// http server address and port
	http_bind := fmt.Sprintf("%s:%s", http_host, http_port)

	// open connection to database
	db, err := sqlx.Connect("mysql", mysql_bind)
	db.SetMaxIdleConns(100)
	CheckError(err)
	defer db.Close()

	// controllers
	userController := controller.NewUserController(db)

	// user handler
	mx.HandleFunc("/api/v2/users",userController.GetUsers).Methods("GET")
	mx.HandleFunc("/api/v2/users/{id}",userController.GetUser).Methods("GET")
	mx.HandleFunc("/api/v2/users",userController.InsertUser).Methods("POST")
	mx.HandleFunc("/api/v2/users",userController.UpdateUser).Methods("PUT")
	mx.HandleFunc("/api/v2/users/{id}",userController.DeleteUser).Methods("DELETE")

	// static
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	// start http server
	fmt.Println("Listening on " + http_bind)
	err = http.ListenAndServe(http_bind, mx)
	CheckError(err)
}
