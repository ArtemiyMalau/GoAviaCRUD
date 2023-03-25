package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	InitDB()

	config := GetConfig()
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := NewService(db)
	v := validator.New()
	handler := NewHandler(*service, v)

	r := mux.NewRouter()
	handler.RegisterHandlers(r)

	log.Printf("Start listening port %s", config.ListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ListenPort), r))
}
