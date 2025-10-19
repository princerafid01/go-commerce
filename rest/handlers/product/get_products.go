package product

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendData(w, http.StatusOK, productList)
}
