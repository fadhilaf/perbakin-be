package app

import (
	"github.com/FadhilAF/perbakin-be/internal/repository"

	admin_usecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	auth_usecase "github.com/FadhilAF/perbakin-be/internal/usecase/auth"
	super_usecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"
)

type usecases struct {
	super super_usecase.SuperUsecase
	auth  auth_usecase.AuthUsecase
	admin admin_usecase.AdminUsecase
}

func (app *App) initUsecase() {
	store := repository.NewStore(app.dbPool)

	var usecases usecases

	usecases.super = super_usecase.NewSuperUsecase(store)
	usecases.auth = auth_usecase.NewAuthUsecase(store)
	usecases.admin = admin_usecase.NewAdminUsecase(store)

	app.usecase = usecases
}
