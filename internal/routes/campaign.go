package routes

import (
	"github.com/as-ifn-at/targeting_engine/internal/handlers"
)

func (r *router) campaignRoutes() {

	campaignHandler := handlers.NewCampaignHandler(r.appConfig, r.logger)
	routerG := r.router.Group("v1/campaign")
	routerG.POST("/create", campaignHandler.Save)
	routerG.GET("/:id", campaignHandler.Get)
}
