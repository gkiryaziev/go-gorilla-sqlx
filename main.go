package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/conf"
	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/handlers/users"
)

// checkError check errors
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// load config
	config, err := conf.NewConfig("config.yaml").Load()
	checkError(err)

	// mysql connection string
	mysqlBind := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.DB.UserName,
		config.DB.UserPassword,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)

	// http server address and port
	hostBind := fmt.Sprintf("%s:%s",
		config.Host.IP,
		config.Host.Port,
	)

	// open connection to database
	db, err := sqlx.Connect("mysql", mysqlBind)
	checkError(err)
	db.SetMaxIdleConns(100)
	defer db.Close()

	// handlers
	userHandler := users.NewUserHandler(db)

	mx := mux.NewRouter()

	// user handler
	mx.HandleFunc("/api/v1/users", userHandler.GetUsers).Methods("GET")
	mx.HandleFunc("/api/v1/users/{id}", userHandler.GetUser).Methods("GET")
	mx.HandleFunc("/api/v1/users", userHandler.UpdateUser).Methods("PUT")
	mx.HandleFunc("/api/v1/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	mx.HandleFunc("/api/v1/users", userHandler.InsertUser).Methods("POST")

	// static
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	// negroni
	ng := negroni.New()
	ng.UseHandler(mx)

	// start server
	log.Println("Listening on", hostBind)
	err = http.ListenAndServe(hostBind, ng)
	checkError(err)
}
