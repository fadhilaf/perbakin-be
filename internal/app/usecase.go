package app

import (
	"github.com/FadhilAF/perbakin-be/internal/repository"
	admin_usecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
)

type usecases struct {
	admin admin_usecase.AdminUsecase
}

func (app *App) initUsecase() {
	store := repository.NewStore(app.dbPool)

	var usecases usecases

	usecases.admin = admin_usecase.NewAdminUsecase(store)

	app.usecase = usecases
}
