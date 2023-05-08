package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetAllScorers() model.WebServiceResponse {
	scorers, err := usecase.Store.GetAllScorers(context.Background())
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data penguji", http.StatusInternalServerError, nil)
	}

	var scorersDisplayData []model.UserDisplayData
	for _, admin := range scorers {
		scorersDisplayData = append(scorersDisplayData, model.UserDisplayData{
			ID:        admin.ID,
			Name:      admin.Name,
			CreatedAt: admin.CreatedAt.Time,
			UpdatedAt: admin.UpdatedAt.Time,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penguji", http.StatusOK, gin.H{
		"scorers": scorersDisplayData,
	})
}
