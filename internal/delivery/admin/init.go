package delivery

import (
	usecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	// "github.com/gin-gonic/gin"
)

type AdminDelivery interface {
}

var _ AdminDelivery = &adminHandler{}

func NewAdminDelivery(usecase usecase.AdminUsecase) AdminDelivery {
	return &adminHandler{
		Usecase: usecase,
	}
}

type adminHandler struct {
	Usecase usecase.AdminUsecase
}
