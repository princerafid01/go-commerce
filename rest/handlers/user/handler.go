package user

import (
	"ecommerce/config"
)

type Handler struct {
	// userRepo repo.UserRepo
	cnf *config.Config
	svc Service
}

func NewHandler(cnf *config.Config, svc Service) *Handler {
	return &Handler{
		svc: svc,
		cnf: cnf,
	}
}
