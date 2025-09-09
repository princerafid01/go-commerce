package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me a valid JSON", 400)
		return
	}
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	utils.SendData(w, newProduct, 201)
}
