package app

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"
	scorerDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/scorer"
	superDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"
)

type deliveries struct {
	super      superDelivery.SuperDelivery
	admin      adminDelivery.AdminDelivery
	scorer     scorerDelivery.ScorerDelivery
	all        allDelivery.AllDelivery
	adminSuper adminSuperDelivery.AdminSuperDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries

	deliveries.super = superDelivery.NewSuperDelivery(app.usecase.super)
	deliveries.admin = adminDelivery.NewAdminDelivery(app.usecase.admin)
	deliveries.scorer = scorerDelivery.NewScorerDelivery(app.usecase.scorer)

	deliveries.adminSuper = adminSuperDelivery.NewAdminSuperDelivery(app.usecase.adminSuper)
	deliveries.all = allDelivery.NewAllDelivery(app.usecase.all)

	app.delivery = deliveries
}
