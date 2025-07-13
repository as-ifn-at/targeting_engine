package routes

import (
	"github.com/as-ifn-at/targeting_engine/internal/handlers"
)

func (r *router) targetRuleRoutes() {

	bookingHandler := handlers.NewRuleHandler(r.appConfig, r.logger)
	routerG := r.router.Group("/rule/v1")
	routerG.POST("/book", bookingHandler.Save)
	routerG.GET("/:id", bookingHandler.Get)
}
