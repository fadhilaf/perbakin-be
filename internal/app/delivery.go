package app

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	superDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"
)

type deliveries struct {
	super superDelivery.SuperDelivery
	admin adminDelivery.AdminDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries

	deliveries.super = superDelivery.NewSuperDelivery(app.usecase.super)
	deliveries.admin = adminDelivery.NewAdminDelivery(app.usecase.admin)

	app.delivery = deliveries
}
