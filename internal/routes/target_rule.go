package routes

import (
	"github.com/as-ifn-at/targeting_engine/internal/handlers"
)

func (r *router) targetRuleRoutes() {

	bookingHandler := handlers.NewRuleHandler(r.appConfig, r.logger)
	routerG := r.router.Group("v1/rule")
	routerG.POST("/create", bookingHandler.Save)
	routerG.GET("/:id", bookingHandler.Get)
}
