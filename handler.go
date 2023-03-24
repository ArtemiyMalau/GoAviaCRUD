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
	s Service
	v *validator.Validate
}

func NewHandler(s Service, v *validator.Validate) *handler {
	v.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(fl.Param())
		return re.MatchString(fl.Field().String())
	})
	return &handler{s: s, v: v}
}

func (h *handler) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/airline", h.handleAddAirline).Methods("POST")
	r.HandleFunc("/airline/{code}", h.handleDeleteAirlineByCode).Methods("DELETE")
	r.HandleFunc("/airline/{code}/change-providers", h.handleChangeAirlineProviders).Methods("PATCH")

	r.HandleFunc("/provider", h.handleAddProvider).Methods("POST")
	r.HandleFunc("/provider/{id}", h.handleDeleteProviderByCode).Methods("DELETE")
	r.HandleFunc("/provider/{id}/airlines", h.handleGetProviderAirlines).Methods("GET")

	r.HandleFunc("/schema", h.handleAddSchema).Methods("POST")
	r.HandleFunc("/schema/{name}", h.handleGetSchemeByName).Methods("GET")
	r.HandleFunc("/schema/{id}", h.handleUpdateSchemeById).Methods("PATCH")
	r.HandleFunc("/schema/{id}", h.handleDeleteSchemeById).Methods("DELETE")

	r.HandleFunc("/account", h.handleAddAccount).Methods("POST")
	r.HandleFunc("/account/{id}/set-scheme", h.handleSetAccountScheme).Methods("POST")
	r.HandleFunc("/account/{id}", h.handleDeleteAccountById).Methods("DELETE")
	r.HandleFunc("/account/{id}/airlines", h.handleGetAccountAirlines).Methods("GET")
}

func (h *handler) handleAddAirline(w http.ResponseWriter, r *http.Request) {
	var dto AirlineDTOAdd
	if err := decodeAndValidate(&dto, r, h.v, nil); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.addAirline(context.TODO(), dto); err != nil {
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

	if err := h.s.deleteAirlineByCode(context.TODO(), code); err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleChangeAirlineProviders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		writeJSON(w, http.StatusBadRequest, "Code is missing in parameters")
	}

	var dto AirlineDTOChangeProviders
	if err := decodeAndValidate(&dto, r, h.v, func(object *AirlineDTOChangeProviders) error {
		object.Code = code
		return nil
	}); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.changeAirlineProviders(context.TODO(), dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleAddProvider(w http.ResponseWriter, r *http.Request) {
	var dto ProviderDTOAdd
	if err := decodeAndValidate(&dto, r, h.v, nil); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.addProvider(context.TODO(), dto); err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleDeleteProviderByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		writeJSON(w, http.StatusBadRequest, "Code is missing in parameters")
	}

	if err := h.s.deleteProviderByCode(context.TODO(), code); err != nil {
		writeJSON(w, http.StatusBadRequest, nil)
		return
	}

	writeJSON(w, http.StatusOK, nil)
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

func decodeAndValidate[T any](object T, r *http.Request, v *validator.Validate, addFields func(object T) error) error {
	if err := json.NewDecoder(r.Body).Decode(object); err != nil {
		return err
	}

	if addFields != nil {
		if err := addFields(object); err != nil {
			return err
		}
	}

	if err := v.Struct(object); err != nil {
		return err
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
