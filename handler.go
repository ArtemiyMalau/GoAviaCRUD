package main

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

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
	r.HandleFunc("/provider/{code}", h.handleDeleteProviderByCode).Methods("DELETE")
	r.HandleFunc("/provider/{code}/airlines", h.handleGetProviderAirlinesByCode).Methods("GET")

	r.HandleFunc("/schema", h.handleAddSchema).Methods("POST")
	r.HandleFunc("/schema/{name}", h.handleGetSchemeByName).Methods("GET")
	r.HandleFunc("/schema/{id}", h.handleUpdateSchemeById).Methods("PATCH")
	r.HandleFunc("/schema/{id}", h.handleDeleteSchemeById).Methods("DELETE")

	r.HandleFunc("/account", h.handleAddAccount).Methods("POST")
	r.HandleFunc("/account/{id}/set-scheme", h.handleSetAccountScheme).Methods("PATCH")
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
	code := vars["code"]

	if err := h.s.deleteAirlineByCode(context.TODO(), code); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleChangeAirlineProviders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

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
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (h *handler) handleDeleteProviderByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	if err := h.s.deleteProviderByCode(context.TODO(), code); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleGetProviderAirlinesByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	airlines, err := h.s.getProviderAirlinesByCode(context.TODO(), code)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, airlines)
}

func (h *handler) handleAddSchema(w http.ResponseWriter, r *http.Request) {
	var dto SchemaDTOAdd
	if err := decodeAndValidate(&dto, r, h.v, nil); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.addSchema(context.TODO(), dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (h *handler) handleGetSchemeByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	scheme, err := h.s.getSchemeByName(context.TODO(), name)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, scheme)
}

func (h *handler) handleUpdateSchemeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var dto SchemaDTOUpdate
	if err := decodeAndValidate(&dto, r, h.v, nil); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	dto.Id = id

	if err := h.s.updateSchemeById(context.TODO(), dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleDeleteSchemeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.deleteSchemeById(context.TODO(), id); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleAddAccount(w http.ResponseWriter, r *http.Request) {
	var dto AccountDTOAdd
	if err := decodeAndValidate(&dto, r, h.v, nil); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.addAccount(context.TODO(), dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (h *handler) handleSetAccountScheme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var dto AccountDTOSetScheme
	if err := decodeAndValidate(&dto, r, h.v, nil); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	dto.Id = id

	if err := h.s.setAccountScheme(context.TODO(), dto); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleDeleteAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.deleteAccountById(context.TODO(), id); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (h *handler) handleGetAccountAirlines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	airlines, err := h.s.getAccountAirlinesById(context.TODO(), id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, airlines)
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
