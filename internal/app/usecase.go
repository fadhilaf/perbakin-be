package app

import (
// "github.com/FadhilAF/perbakin-be/internal/repository"
)

// division_usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"

type usecases struct {
	// division   division_usecase.DivisionUsecase
}

func (app *App) initUsecase() {
	// store := repository.NewStore(app.dbPool)

	var usecases usecases

	// usecases.division = division_usecase.NewDivisionUsecase(app.store)

	app.usecase = usecases
}
