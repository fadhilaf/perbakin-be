package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage5No(req model.UpdateStage123456NoRequest) model.WebServiceResponse {
	var scores string
	var err error

	switch req.Try {
	case "1":
		status, _ := usecase.Store.GetStage5try1Status(context.Background(), req.ID)
		if status != repositoryModel.Stage5Status(req.No) {
			return util.ToWebServiceResponse("Gagal mengupdate stage 5 percobaan 1 no "+req.No+", sekarang sedang no "+string(status), http.StatusForbidden, nil)
		}

		switch req.No {
		case "1":
			scores, err = usecase.Store.UpdateStage5try1No1(context.Background(), repositoryModel.UpdateStage5try1No1Params{
				ID:  req.ID,
				No1: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 5 percobaan 1 no 1: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "2":
			scores, err = usecase.Store.UpdateStage5try1No2(context.Background(), repositoryModel.UpdateStage5try1No2Params{
				ID:  req.ID,
				No2: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 5 percobaan 1 no 2: "+err.Error(), http.StatusInternalServerError, nil)
			}
		}
	case "2":
		status, _ := usecase.Store.GetStage5try2Status(context.Background(), req.ID)
		if status != repositoryModel.Stage5Status(req.No) {
			return util.ToWebServiceResponse("Gagal mengupdate stage 5 percobaan 2 no "+req.No+", sekarang sedang no "+string(status), http.StatusForbidden, nil)
		}

		switch req.No {
		case "1":
			scores, err = usecase.Store.UpdateStage5try2No1(context.Background(), repositoryModel.UpdateStage5try2No1Params{
				ID:  req.ID,
				No1: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 5 percobaan 2 no 1: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "2":
			scores, err = usecase.Store.UpdateStage5try2No2(context.Background(), repositoryModel.UpdateStage5try2No2Params{
				ID:  req.ID,
				No2: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 5 percobaan 2 no 2: "+err.Error(), http.StatusInternalServerError, nil)
			}
		}
	}

	return util.ToWebServiceResponse("Berhasil mengupdate stage 5 percobaan "+req.Try+" no "+req.No, http.StatusOK, gin.H{
		"no": util.Stage5DatabaseNumbersToStruct(scores),
	})
}
