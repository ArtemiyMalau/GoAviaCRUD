package main

type providerCode string

type schemaId int

type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AirlineModel struct {
	Id int
	Airline
}

type AirlineDTOAdd struct {
	Code string `json:"code" validate:"required,regexp=^[A-ZА-Я0-9]{2}$"`
	Name string `json:"name" validate:"required"`
}

type AirlineDTOChangeProviders struct {
	Code      string         `json:"code" validate:"required,regexp=^[A-ZА-Я0-9]{2}$"`
	Providers []providerCode `json:"providers" validate:"required,dive,regexp=^[A-Z]{2}$"`
}

type Provider struct {
	Code providerCode `json:"id"`
	Name string       `json:"name"`
}

type ProviderModel struct {
	Id int
	Airline
}

type ProviderDTOAdd struct {
	Id   providerCode `json:"id" validate:"required,regexp=^[A-Z]{2}$"`
	Name string       `json:"name" validate:"required"`
}

type Schema struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Providers []providerCode `json:"providers"`
}

type Account struct {
	Id       int      `json:"id"`
	SchemaId schemaId `json:"schema_id"`
}
