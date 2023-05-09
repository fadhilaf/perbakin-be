package usecase

import (
	"context"
	"fmt"
	"net/http"

	respositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateScorer(req model.UpdateUserRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorerById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan scorer dengan ID tersebut", http.StatusNotFound, nil)
	}

	passwordHash, err := util.HashPassword(req.Body.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	err = usecase.Store.UpdateScorer(context.Background(), respositoryModel.UpdateScorerParams{
		ID:       scorer.ID,
		Username: req.Body.Username,
		Password: passwordHash,
		Name:     req.Body.Name,
	})
	if err != nil {
		fmt.Println(err)
		return util.ToWebServiceResponse("Gagal mengubah scorer", http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah scorer", http.StatusOK, nil)
}
