package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetAdminById(req model.ByIdRequest) model.WebServiceResponse {
	admin, _ := usecase.Store.GetAdminById(context.Background(), req.ID)

	return util.ToWebServiceResponse("Data admin ditemukan", http.StatusOK, gin.H{
		"admin": model.Operator{
			ID:     admin.ID,
			ExamID: admin.ExamID,
			User: model.User{
				ID:       admin.UserID,
				Username: admin.Username,
				Name:     admin.Name,
			},
		}})
}
