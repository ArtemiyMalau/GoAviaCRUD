package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := GetConfig()
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	service := NewService(db)
	handler := NewHandler(*service)

	r := mux.NewRouter()
	r.HandleFunc("/airline", handler.handleAddAirline).Methods("POST")
	r.HandleFunc("/airline/{code}", handler.handleDeleteAirlineByCode).Methods("DELETE")
	r.HandleFunc("/airline/{code}/change-providers", handler.handleChangeAirlineProviders).Methods("PATCH")

	r.HandleFunc("/provider", handler.handleAddProvider).Methods("POST")
	r.HandleFunc("/provider/{id}", handler.handleDeleteProviderById).Methods("DELETE")
	r.HandleFunc("/provider/{id}/airlines", handler.handleGetProviderAirlines).Methods("GET")

	r.HandleFunc("/schema", handler.handleAddSchema).Methods("POST")
	r.HandleFunc("/schema/{name}", handler.handleGetSchemeByName).Methods("GET")
	r.HandleFunc("/schema/{id}", handler.handleUpdateSchemeById).Methods("PATCH")
	r.HandleFunc("/schema/{id}", handler.handleDeleteSchemeById).Methods("DELETE")

	r.HandleFunc("/account", handler.handleAddAccount).Methods("POST")
	r.HandleFunc("/account/{id}/set-scheme", handler.handleSetAccountScheme).Methods("POST")
	r.HandleFunc("/account/{id}", handler.handleDeleteAccountById).Methods("DELETE")
	r.HandleFunc("/account/{id}/airlines", handler.handleGetAccountAirlines).Methods("GET")

	log.Printf("Start listening port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
