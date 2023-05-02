package app

import (
	"github.com/FadhilAF/perbakin-be/internal/repository"

	super_usecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"
)

type usecases struct {
	super super_usecase.SuperUsecase
}

func (app *App) initUsecase() {
	store := repository.NewStore(app.dbPool)

	var usecases usecases

	usecases.super = super_usecase.NewSuperUsecase(store)

	app.usecase = usecases
}
