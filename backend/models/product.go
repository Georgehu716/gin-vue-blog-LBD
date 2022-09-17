package models

type Product struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Qty int `json:"qty"`
	LastUpdated string `json:"last_updated"`
}