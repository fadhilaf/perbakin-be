package app

import (
	admin_delivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	auth_delivery "github.com/FadhilAF/perbakin-be/internal/delivery/auth"
	super_delivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"
)

type deliveries struct {
	super super_delivery.SuperDelivery
	auth  auth_delivery.AuthDelivery
	admin admin_delivery.AdminDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries

	deliveries.super = super_delivery.NewSuperDelivery(app.usecase.super)
	deliveries.auth = auth_delivery.NewAuthDelivery(app.usecase.auth)
	deliveries.admin = admin_delivery.NewAdminDelivery(app.usecase.admin)

	app.delivery = deliveries
}
