package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) CreateScorer(req model.CreateUserRequest) model.WebServiceResponse {
	_, err := usecase.Store.GetUserByUsername(context.Background(), req.Username)
	if err == nil {
		return util.ToWebServiceResponse("Username sudah digunakan", http.StatusConflict, nil)
	}

	passwordHash, err := util.HashPassword(req.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	err = usecase.Store.CreateScorer(context.Background(), repositoryModel.CreateScorerParams{
		Username: req.Username,
		Password: passwordHash,
		Name:     req.Name,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat penguji: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat penguji", http.StatusCreated, nil)
}
