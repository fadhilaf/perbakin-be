package app

import (
	super_delivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"
)

type deliveries struct {
	super super_delivery.SuperDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries

	deliveries.super = super_delivery.NewSuperDelivery(app.usecase.super)

	app.delivery = deliveries
}
