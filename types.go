package main

type providerId string

type schemaId int

type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AirlineDTOAdd struct {
	Code string `json:"code" validate:"required,regexp=^[A-ZА-Я0-9]{2}$"`
	Name string `json:"name" validate:"required"`
}

type AirlineDTOChangeProviders struct {
	Providers []providerId `json:"providers"`
}

type Provider struct {
	Id   providerId `json:"id"`
	Name string     `json:"name"`
}

type Schema struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Providers []providerId `json:"providers"`
}

type Account struct {
	Id       int      `json:"id"`
	SchemaId schemaId `json:"schema_id"`
}
