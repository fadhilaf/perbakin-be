package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetScorerById(req model.OperatorByIdRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorerById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Id tidak terdaftar sebagai scorer", http.StatusUnauthorized, nil)
	}

	if scorer.ExamID != req.ExamID {
		return util.ToWebServiceResponse("Tidak diperbolehkan menampilkan data akun penguji ujian lain", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data scorer ditemukan", http.StatusOK, gin.H{
		"scorer": model.Operator{
			ID:     scorer.ID,
			ExamID: scorer.ExamID,
			User: model.User{
				ID:        scorer.UserID,
				Username:  scorer.Username,
				Name:      scorer.Name,
				CreatedAt: scorer.CreatedAt.Time,
				UpdatedAt: scorer.UpdatedAt.Time,
			},
		}})
}
