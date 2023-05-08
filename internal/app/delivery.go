package app

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	superDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"
)

type deliveries struct {
	super      superDelivery.SuperDelivery
	admin      adminDelivery.AdminDelivery
	adminSuper adminSuperDelivery.AdminSuperDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries

	deliveries.super = superDelivery.NewSuperDelivery(app.usecase.super)
	deliveries.admin = adminDelivery.NewAdminDelivery(app.usecase.admin)
	deliveries.adminSuper = adminSuperDelivery.NewAdminSuperDelivery(app.usecase.adminSuper)

	app.delivery = deliveries
}
