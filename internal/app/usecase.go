package app

import (
	"github.com/FadhilAF/perbakin-be/internal/repository"

	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	superUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"
)

type usecases struct {
	super      superUsecase.SuperUsecase
	admin      adminUsecase.AdminUsecase
	adminSuper adminSuperUsecase.AdminSuperUsecase
}

func (app *App) initUsecase() {
	store := repository.NewStore(app.dbPool)

	var usecases usecases

	usecases.super = superUsecase.NewSuperUsecase(store)
	usecases.admin = adminUsecase.NewAdminUsecase(store)
	usecases.adminSuper = adminSuperUsecase.NewAdminSuperUsecase(store)

	app.usecase = usecases
}
