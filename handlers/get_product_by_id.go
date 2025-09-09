package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give a valid product ID", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			utils.SendData(w, product, 200)
			return
		}
	}

	utils.SendData(w, "Product Not Found 200", 404)

}
