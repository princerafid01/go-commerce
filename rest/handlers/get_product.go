package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give a valid product ID", 400)
		return
	}

	product := database.Get(id)

	if product == nil {
		utils.SendError(w, 404, "Product Not Found")
		return
	}

	utils.SendData(w, product, 200)

}
