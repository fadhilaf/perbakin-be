package app

import (
	"github.com/FadhilAF/perbakin-be/internal/repository"

	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"
	scorerUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/scorer"
	superUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"
	utilUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/util"
)

type usecases struct {
	super      superUsecase.SuperUsecase
	admin      adminUsecase.AdminUsecase
	adminSuper adminSuperUsecase.AdminSuperUsecase
	scorer     scorerUsecase.ScorerUsecase
	all        allUsecase.AllUsecase
	util       utilUsecase.UtilUsecase
}

func (app *App) initUsecase() {
	store := repository.NewStore(app.dbPool)

	var usecases usecases

	usecases.super = superUsecase.NewSuperUsecase(store)
	usecases.admin = adminUsecase.NewAdminUsecase(store)
	usecases.scorer = scorerUsecase.NewScorerUsecase(store)
	usecases.adminSuper = adminSuperUsecase.NewAdminSuperUsecase(store)
	usecases.all = allUsecase.NewAllUsecase(store)
	usecases.util = utilUsecase.NewUtilUsecase(store)

	app.usecase = usecases
}
