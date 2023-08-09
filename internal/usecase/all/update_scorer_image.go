package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *allUsecaseImpl) UpdateScorerImage(req model.UpdateImageRequest) model.WebServiceResponse {
	imagePath, err := usecase.Store.UpdateScorerImage(context.Background(), repositoryModel.UpdateScorerImageParams{
		ID:        req.ID,
		ImagePath: req.ImagePath,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah gambar penguji: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah gambar penguji", http.StatusOK, gin.H{
		"image_path": imagePath,
	})
}
