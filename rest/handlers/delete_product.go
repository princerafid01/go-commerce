package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give a valid product ID", 400)
		return
	}

	database.Delete(id)

	utils.SendData(w, "Successfully Deleted the product", 200)
}
