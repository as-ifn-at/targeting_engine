package handlers

import (
	"github.com/as-ifn-at/targeting_engine/internal/config"
	"github.com/as-ifn-at/targeting_engine/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var Campaigns = make(map[string]models.Campaign, 0)

type campaignHandler struct {
	Handler
	config config.Config
	logger zerolog.Logger
}

func NewCampaignHandler(config config.Config, logger zerolog.Logger) Handler {
	return &campaignHandler{
		config: config,
		logger: logger,
	}
}

func (h *campaignHandler) Get(ctx *gin.Context) {

}

func (h *campaignHandler) Save(ctx *gin.Context) {

}
