package app

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"
	scorerDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/scorer"
	superDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"
)

type deliveries struct {
	super  superDelivery.SuperDelivery
	admin  adminDelivery.AdminDelivery
	scorer scorerDelivery.ScorerDelivery
	all    allDelivery.AllDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries

	deliveries.super = superDelivery.NewSuperDelivery(app.usecase.super, app.usecase.adminSuper, app.usecase.all)
	deliveries.admin = adminDelivery.NewAdminDelivery(app.usecase.admin, app.usecase.adminSuper, app.usecase.all)
	deliveries.scorer = scorerDelivery.NewScorerDelivery(app.usecase.scorer, app.usecase.all)
	deliveries.all = allDelivery.NewAllDelivery(app.usecase.all)

	app.delivery = deliveries
}
