package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage0Series(req model.UpdateStage0SeriesRequest) model.WebServiceResponse {
	var scores string
	var err error

	status, _ := usecase.Store.GetStage0Status(context.Background(), req.ID)
	if status != repositoryModel.Stage0Status(req.Series) {
		return util.ToWebServiceResponse("Gagal mengupdate kualifikasi seri "+req.Series+", sekarang sedang seri "+string(status), http.StatusForbidden, nil)
	}

	switch req.Series {
	case "1":
		scores, err = usecase.Store.UpdateStage0Series1(context.Background(), repositoryModel.UpdateStage0Series1Params{
			ID:      req.ID,
			Series1: req.Scores,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate kualifikasi seri 1: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "2":
		scores, err = usecase.Store.UpdateStage0Series2(context.Background(), repositoryModel.UpdateStage0Series2Params{
			ID:      req.ID,
			Series2: req.Scores,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate kualifikasi seri 2: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "3":
		scores, err = usecase.Store.UpdateStage0Series3(context.Background(), repositoryModel.UpdateStage0Series3Params{
			ID:      req.ID,
			Series3: req.Scores,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate kualifikasi seri 3: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "4":
		scores, err = usecase.Store.UpdateStage0Series4(context.Background(), repositoryModel.UpdateStage0Series4Params{
			ID:      req.ID,
			Series4: req.Scores,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate kualifikasi seri 4: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "5":
		scores, err = usecase.Store.UpdateStage0Series5(context.Background(), repositoryModel.UpdateStage0Series5Params{
			ID:      req.ID,
			Series5: req.Scores,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate kualifikasi seri 5: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	return util.ToWebServiceResponse("Berhasil mengupdate kualifikasi seri "+req.Series, http.StatusOK, gin.H{
		"scores": util.ScoresToIntArray(scores),
	})
}
