package app

import (
	// division_delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/division"
)

type deliveries struct {
	// division   division_delivery.DivisionDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	// deliveries.division = division_delivery.NewDivisionDelivery(app.usecase.division)

	app.delivery = deliveries
}
