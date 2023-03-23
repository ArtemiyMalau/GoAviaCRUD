package main

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{service: s}
}

func (h *handler) handleAddAirline(w http.ResponseWriter, r *http.Request) {
	var dto AirlineDTOAdd
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	v := validator.New()
	v.RegisterValidation("regexp", Regexp)
	if err := v.Struct(&dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.addAirline(context.TODO(), dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (h *handler) handleDeleteAirlineByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		writeJSON(w, http.StatusBadRequest, "Code is missing in parameters")
	}

	if err := h.service.deleteAirlineByCode(context.TODO(), code); err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleChangeAirlineProviders(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleAddProvider(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleDeleteProviderById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleGetProviderAirlines(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleAddSchema(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleGetSchemeByName(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleUpdateSchemeById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleDeleteSchemeById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleAddAccount(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleSetAccountScheme(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleDeleteAccountById(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented error")
}

func (h *handler) handleGetAccountAirlines(w http.ResponseWriter, r *http.Request) {
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
