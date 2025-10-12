package handlers

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	usr := database.Find(reqLogin.Email, reqLogin.Password)

	if usr == nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := utils.CreateJwt(cnf.JwtSecretKey, utils.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})

	if err != nil {
		http.Error(w, "Internel Server Error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusCreated)
}
