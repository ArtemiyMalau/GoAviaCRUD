package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/airline", handleAddAirline).Methods("POST")
	r.HandleFunc("/airline/{code}", handleDeleteAirlineByCode).Methods("DELETE")
	r.HandleFunc("/airline/{id}/change-providers", handleChangeAirlineProviders).Methods("PATCH")

	r.HandleFunc("/provider", handleAddProvider).Methods("POST")
	r.HandleFunc("/provider/{id}", handleDeleteProviderById).Methods("DELETE")
	r.HandleFunc("/provider/{id}/airlines", handleGetProviderAirlines).Methods("GET")

	r.HandleFunc("/schema", handleAddSchema).Methods("POST")
	r.HandleFunc("/schema/{name}", handleGetSchemeByName).Methods("GET")
	r.HandleFunc("/schema/{id}", handleUpdateSchemeById).Methods("PATCH")
	r.HandleFunc("/schema/{id}", handleDeleteSchemeById).Methods("DELETE")

	r.HandleFunc("/account", handleAddAccount).Methods("POST")
	r.HandleFunc("/account/{id}/set-scheme", handleSetAccountScheme).Methods("POST")
	r.HandleFunc("/account/{id}", handleDeleteAccountById).Methods("DELETE")
	r.HandleFunc("/account/{id}/airlines", handleGetAccountAirlines).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func handleAddAirline(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleDeleteAirlineByCode(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleChangeAirlineProviders(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleAddProvider(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleDeleteProviderById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleGetProviderAirlines(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleAddSchema(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleGetSchemeByName(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleUpdateSchemeById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleDeleteSchemeById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleAddAccount(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleSetAccountScheme(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleDeleteAccountById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func handleGetAccountAirlines(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}
