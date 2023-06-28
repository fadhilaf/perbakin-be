package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateShooterImage(req model.UpdateImageRequest) model.WebServiceResponse {
	imagePath, err := usecase.Store.UpdateShooterImage(context.Background(), repositoryModel.UpdateShooterImageParams{
		ID:        req.ID,
		ImagePath: req.ImagePath,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah gambar penembak: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah gambar penembak", http.StatusOK, gin.H{
		"image_path": imagePath,
	})
}
