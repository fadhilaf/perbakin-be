package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage0Finish(req model.UpdateStage0FinishRequest) model.WebServiceResponse {
	status, _ := usecase.Store.GetStage0Status(context.Background(), req.ID)
	if status != "5" {
		var msg string
		if status == "6" {
			msg = "Pendataan kualifikasi sudah selesai"
		} else {
			msg = "Tidak dapat menyelesaikan, masih pada seri ke-" + string(status)
		}
		return util.ToWebServiceResponse(msg, http.StatusForbidden, nil)
	}

	if !req.Body.Success {
		if err := usecase.Store.UpdateResultFailed(context.Background(), req.ResultID); err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate hasil ujian menjadi gagal: "+err.Error(), http.StatusInternalServerError, nil)
		}
	} else {
		if err := usecase.Store.UpdateResultNextStage(context.Background(), repositoryModel.UpdateResultNextStageParams{ID: req.ResultID, Stage: repositoryModel.NullStages{Stages: repositoryModel.Stages("1"), Valid: true}}); err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate hasil ujian ke tahap selanjutnya: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	finish, err := usecase.Store.UpdateStage0Finish(context.Background(), repositoryModel.UpdateStage0FinishParams{
		ID:          req.ID,
		ShooterSign: req.ShooterSign,
		ScorerSign:  req.ScorerSign,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal menyelesaikan pendataan kualifikasi: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menyelesaikan pendataan kualifikasi", http.StatusOK, gin.H{
		"stage_0": model.FinishStage0Response{
			Status:      string(finish.Status),
			ShooterSign: finish.ShooterSign,
			ScorerSign:  finish.ScorerSign,
		}})
}
