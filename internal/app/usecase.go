package app

import (
	// division_usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"
	// generation_usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/generation"
	// role_usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/role"
)

type usecases struct {
	// division   division_usecase.DivisionUsecase
}

func (app *App) initUsecase() {
	var usecases usecases

	// usecases.division = division_usecase.NewDivisionUsecase(app.store)

	app.usecase = usecases
}
