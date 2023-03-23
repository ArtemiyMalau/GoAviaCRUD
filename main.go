package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	config := GetConfig()

	r := mux.NewRouter()
	r.HandleFunc("/airline", handleAddAirline).Methods("POST")
	r.HandleFunc("/airline/{code}", handleDeleteAirlineByCode).Methods("DELETE")
	r.HandleFunc("/airline/{code}/change-providers", handleChangeAirlineProviders).Methods("PATCH")

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

	log.Printf("Start listening port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

func handleAddAirline(w http.ResponseWriter, r *http.Request) {
	var dto AirlineDTOAdd
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeJSON(w, 400, err.Error())
		return
	}

	v := validator.New()
	v.RegisterValidation("regexp", Regexp)
	if err := v.Struct(&dto); err != nil {
		writeJSON(w, 400, err.Error())
		return
	}

	// Business logic
	config := GetConfig()
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO airline (code, name) VALUES (:code, :name)", &dto)
	tx.Commit()
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

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func Regexp(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(fl.Param())
	return re.MatchString(fl.Field().String())
}

// type apiError struct {
// 	Err string
// 	Status int
// }

// func (e apiError) Error() string {
// 	return e.Err
// }

// type apiFunc func(w http.ResponseWriter, r *http.Request) error

// func makeHTTPHandler(f apiFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := f(w, r); err != nil {
// 			if e, ok := err.(apiError); ok {
// 				writeJSON(w, e.Status, e)
// 				return
// 			}

// 			writeJSON(w, http.StatusInternalServerError, {Error: "Server"})
// 		}
// 	}
// }
