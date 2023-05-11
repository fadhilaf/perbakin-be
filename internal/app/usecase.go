package app

import (
	"github.com/FadhilAF/perbakin-be/internal/repository"

	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	scorerUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/scorer"
	superUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"

	sessionUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/session"
)

type usecases struct {
	super      superUsecase.SuperUsecase
	admin      adminUsecase.AdminUsecase
	adminSuper adminSuperUsecase.AdminSuperUsecase
	scorer     scorerUsecase.ScorerUsecase

	session sessionUsecase.SessionUsecase
}

func (app *App) initUsecase() {
	store := repository.NewStore(app.dbPool)

	var usecases usecases

	usecases.super = superUsecase.NewSuperUsecase(store)
	usecases.admin = adminUsecase.NewAdminUsecase(store)
	usecases.scorer = scorerUsecase.NewScorerUsecase(store)
	usecases.adminSuper = adminSuperUsecase.NewAdminSuperUsecase(store)

	usecases.session = sessionUsecase.NewSessionUsecase(store)

	app.usecase = usecases
}
