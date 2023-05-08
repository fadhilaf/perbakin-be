package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) GetAdminSuperByUserId(req model.GetByUserIdRequest) model.WebServiceResponse {
	_, err1 := usecase.Store.GetSuperByUserId(context.Background(), req.UserID)
	_, err2 := usecase.Store.GetAdminByUserId(context.Background(), req.UserID)

	if err1 != nil && err2 != nil {
		return util.ToWebServiceResponse("User tidak terdaftar sebagai admin maupun super admin", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Izin diberikan", http.StatusOK, nil)
}
