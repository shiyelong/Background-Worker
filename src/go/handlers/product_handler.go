package handlers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Product 1", Desc: "This is the first product"},
		{ID: 2, Name: "Product 2", Desc: "This is the second product"},
		{ID: 3, Name: "Product 3", Desc: "This is the third product"},
	}

	json.NewEncoder(w).Encode(products)
}
