package models

type Cart struct {
	Id      uint `json: "id"`
	Owner   int  `json: "cart_owner"`
	Product int  `json: "product_id"`
}
