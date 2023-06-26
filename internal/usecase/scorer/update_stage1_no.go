package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage1No(req model.UpdateStage123456NoRequest) model.WebServiceResponse {
	var scores string
	var err error

	switch req.Try {
	case "1":
		status, _ := usecase.Store.GetStage1try1Status(context.Background(), req.ID)
		if status != repositoryModel.Stage13Status(req.No) {
			return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 no "+req.No+", sekarang sedang no "+string(status), http.StatusForbidden, nil)
		}

		switch req.No {
		case "1":
			scores, err = usecase.Store.UpdateStage1try1No1(context.Background(), repositoryModel.UpdateStage1try1No1Params{
				ID:  req.ID,
				No1: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 no 1: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "2":
			scores, err = usecase.Store.UpdateStage1try1No2(context.Background(), repositoryModel.UpdateStage1try1No2Params{
				ID:  req.ID,
				No2: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 no 2: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "3":
			scores, err = usecase.Store.UpdateStage1try1No3(context.Background(), repositoryModel.UpdateStage1try1No3Params{
				ID:  req.ID,
				No3: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 no 3: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "4":
			scores, err = usecase.Store.UpdateStage1try1No4(context.Background(), repositoryModel.UpdateStage1try1No4Params{
				ID:  req.ID,
				No4: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 no 4: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "5":
			scores, err = usecase.Store.UpdateStage1try1No5(context.Background(), repositoryModel.UpdateStage1try1No5Params{
				ID:  req.ID,
				No5: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 no 5: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "6":
			scores, err = usecase.Store.UpdateStage1try1No6(context.Background(), repositoryModel.UpdateStage1try1No6Params{
				ID:  req.ID,
				No6: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 no 6: "+err.Error(), http.StatusInternalServerError, nil)
			}
		}
	case "2":
		status, _ := usecase.Store.GetStage1try2Status(context.Background(), req.ID)
		if status != repositoryModel.Stage13Status(req.No) {
			return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 no "+req.No+", sekarang sedang no "+string(status), http.StatusForbidden, nil)
		}

		switch req.No {
		case "1":
			scores, err = usecase.Store.UpdateStage1try2No1(context.Background(), repositoryModel.UpdateStage1try2No1Params{
				ID:  req.ID,
				No1: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 no 1: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "2":
			scores, err = usecase.Store.UpdateStage1try2No2(context.Background(), repositoryModel.UpdateStage1try2No2Params{
				ID:  req.ID,
				No2: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 no 2: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "3":
			scores, err = usecase.Store.UpdateStage1try2No3(context.Background(), repositoryModel.UpdateStage1try2No3Params{
				ID:  req.ID,
				No3: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 no 3: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "4":
			scores, err = usecase.Store.UpdateStage1try2No4(context.Background(), repositoryModel.UpdateStage1try2No4Params{
				ID:  req.ID,
				No4: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 no 4: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "5":
			scores, err = usecase.Store.UpdateStage1try2No5(context.Background(), repositoryModel.UpdateStage1try2No5Params{
				ID:  req.ID,
				No5: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 no 5: "+err.Error(), http.StatusInternalServerError, nil)
			}
		case "6":
			scores, err = usecase.Store.UpdateStage1try2No6(context.Background(), repositoryModel.UpdateStage1try2No6Params{
				ID:  req.ID,
				No6: req.ScoresAndDuration,
			})
			if err != nil {
				return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 no 6: "+err.Error(), http.StatusInternalServerError, nil)
			}
		}
	}

	return util.ToWebServiceResponse("Berhasil mengupdate stage 1 percobaan "+req.Try+" no "+req.No, http.StatusOK, gin.H{
		"scores": util.Stage123DatabaseNumbersToStruct(scores),
	})
}
