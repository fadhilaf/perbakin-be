package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	// "github.com/gin-gonic/gin"
)

type AdminSuperDelivery interface {
}

var _ AdminSuperDelivery = &adminSuperHandler{}

func NewAdminSuperDelivery(usecase adminSuperUsecase.AdminSuperUsecase) AdminSuperDelivery {
	return &adminSuperHandler{
		Usecase: usecase,
	}
}

type adminSuperHandler struct {
	Usecase adminSuperUsecase.AdminSuperUsecase
}
