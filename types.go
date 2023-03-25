package main

type providerCode string

type schemaId int

// Airline-related
type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AirlineModel struct {
	Id int
	Airline
}

type AirlineDTOAdd struct {
	Code string `json:"code" validate:"required,regexp=^[A-ZА-Я0-9]{2}$" db:"code"`
	Name string `json:"name" validate:"required" db:"name"`
}

type AirlineDTOChangeProviders struct {
	Code      string         `validate:"required,regexp=^[A-ZА-Я0-9]{2}$"`
	Providers []providerCode `json:"providers" validate:"required,dive,regexp=^[A-Z]{2}$"`
}

// Provider-related
type Provider struct {
	Code providerCode `json:"id"`
	Name string       `json:"name"`
}

type ProviderModel struct {
	Id int
	Provider
}

type ProviderDTOAdd struct {
	Code providerCode `json:"id" validate:"required,regexp=^[A-Z]{2}$" db:"code"`
	Name string       `json:"name" validate:"required" db:"name"`
}

// Schema-related
type Schema struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Providers []providerCode `json:"providers"`
}

type SchemaDTOAdd struct {
	Name      string         `json:"name" validate:"required"`
	Providers []providerCode `json:"providers" validate:"required,dive,regexp=^[A-Z]{2}$"`
}

type SchemaDTOUpdate struct {
	Id        int
	Name      string         `json:"name" validate:"omitempty"`
	Providers []providerCode `json:"providers" validate:"omitempty,dive,regexp=^[A-Z]{2}$"`
}

// Account-related
type Account struct {
	Id       int      `json:"id"`
	SchemaId schemaId `json:"schema_id"`
}

type AccountDTOAdd struct {
	SchemaId schemaId `json:"schema_id" validate:"required"`
}

type AccountDTOSetScheme struct {
	Id       int      `db:"id"`
	SchemaId schemaId `json:"schema_id" validate:"required" db:"scheme_id"`
}
