package handlers

import (
	"fmt"
	"net/http"

	"github.com/as-ifn-at/targeting_engine/internal/config"
	"github.com/as-ifn-at/targeting_engine/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var Campaigns = make(map[string]models.Campaign, 0)
var Rules = make(map[string]models.TargetRules, 0)

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
	id := ctx.Param("id")
	for _, campaign := range Campaigns {
		if campaign.CampaignId == id {
			ctx.IndentedJSON(http.StatusOK, campaign)
			return
		}
	}
	h.logger.Error().Msg(fmt.Sprintf("campaign not found: %v", id))
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "campaign not found"})
}

func (h *campaignHandler) Save(ctx *gin.Context) {
	var newCampaign models.Campaign
	if err := ctx.BindJSON(&newCampaign); err != nil {
		h.logger.Error().Msg(fmt.Sprintf("error while unmarshaling campaign details: %v", err.Error()))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Campaigns[newCampaign.CampaignId] = newCampaign
	ctx.IndentedJSON(http.StatusCreated, gin.H{"Campaign created": newCampaign.CampaignId})
}
