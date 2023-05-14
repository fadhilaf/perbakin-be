package app

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	scorerDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/scorer"
	superDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"
)

type deliveries struct {
	super  superDelivery.SuperDelivery
	admin  adminDelivery.AdminDelivery
	scorer scorerDelivery.ScorerDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries

	deliveries.super = superDelivery.NewSuperDelivery(app.usecase.super, app.usecase.adminSuper, app.usecase.session)
	deliveries.admin = adminDelivery.NewAdminDelivery(app.usecase.admin, app.usecase.adminSuper, app.usecase.session)
	deliveries.scorer = scorerDelivery.NewScorerDelivery(app.usecase.scorer, app.usecase.session)

	app.delivery = deliveries
}
