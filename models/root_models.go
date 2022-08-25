package models

type Transaction struct {
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
	Hour   string  `json:"hour" validate:"required"`
	Action *bool   `json:"action" validate:"required"`
}
