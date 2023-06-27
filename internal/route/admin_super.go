package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for adminRouter or examRouter (for super), return scorerRouter
func AdminSuperRoutes(adminExamRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	ExamAdminSuperRoutes(adminExamRouter, delivery)

	scorerRouter := ScorerAdminSuperRoutes(adminExamRouter, delivery)

	shooterRouter := ShooterAllRoutes(scorerRouter, allDelivery)
	ShooterAdminSuperRoutes(scorerRouter, delivery, allDelivery)

	resultRouter := ResultAllRoutes(shooterRouter, allDelivery)
	ResultAdminSuperRoutes(resultRouter, delivery)

	stage0Router := Stage0AllRoutes(resultRouter, allDelivery)
	Stage0AdminSuperRoutes(stage0Router, delivery)

	stage1Router := Stage1AllRoutes(resultRouter, allDelivery)
	Stage1AdminSuperRoutes(stage1Router, delivery, allDelivery)

	stage2Router := Stage2AllRoutes(resultRouter, allDelivery)
	Stage2AdminSuperRoutes(stage2Router, delivery, allDelivery)

	stage3Router := Stage3AllRoutes(resultRouter, allDelivery)
	Stage3AdminSuperRoutes(stage3Router, delivery, allDelivery)

	stage4Router := Stage4AllRoutes(resultRouter, allDelivery)
	Stage4AdminSuperRoutes(stage4Router, delivery, allDelivery)

	stage5Router := Stage5AllRoutes(resultRouter, allDelivery)
	Stage5AdminSuperRoutes(stage5Router, delivery, allDelivery)

	stage6Router := Stage6AllRoutes(resultRouter, allDelivery)
	Stage6AdminSuperRoutes(stage6Router, delivery, allDelivery)
}
