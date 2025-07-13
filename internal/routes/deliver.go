package routes

import (
	"github.com/as-ifn-at/targeting_engine/internal/handlers"
)

func (r *router) deliverRoutes() {

	deliverHandler := handlers.NewDeliverHandler(r.appConfig, r.logger)
	routerG := r.router.Group("v1/delivery")
	routerG.GET("", deliverHandler.Get)
}
