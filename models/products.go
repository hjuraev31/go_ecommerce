package models

type Products struct {
	Id uint `json: "id"`
	Title string `json: "title"`
	Description string `json: "description"`
	Quantity int `json: "quantity"`
	Price uint `json: "price"`
}