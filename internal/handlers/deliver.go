package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/as-ifn-at/targeting_engine/common"
	"github.com/as-ifn-at/targeting_engine/internal/config"
	"github.com/as-ifn-at/targeting_engine/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type deliverHandler struct {
	Handler
	config config.Config
	logger zerolog.Logger
}

func NewDeliverHandler(config config.Config, logger zerolog.Logger) Handler {
	return &deliverHandler{
		config: config,
		logger: logger,
	}
}

func (h *deliverHandler) Get(ctx *gin.Context) {

	urlFields := models.UrlFields{
		App:     ctx.Query("app"),
		Country: ctx.Query("country"),
		OS:      ctx.Query("os"),
	}

	if err := validateDeliverRequest(urlFields); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logger.Error().Msg(fmt.Sprintf("deliver request validation failed :%v", err))
		return
	}

	res := []models.DeliverResponse{}
	for _, campaign := range Campaigns {
		rule, ok := Rules[campaign.CampaignId]
		if ok && campaign.Status == common.CampaignActiveStatus &&
			rulesApplicable(rule.Rules, urlFields) {
			res = append(res, models.DeliverResponse{
				CampaignId: campaign.CampaignId,
				Image:      campaign.Image,
				CTA:        campaign.CTA,
			})
		}
	}
	if len(res) == 0 {
		ctx.IndentedJSON(http.StatusNoContent, res)
		h.logger.Error().Msg("no records found")
		return
	}
	ctx.IndentedJSON(http.StatusOK, res)
}

func (h *deliverHandler) Save(ctx *gin.Context) {
	// not required for now
}

func validateDeliverRequest(urlfields models.UrlFields) error {
	if urlfields.App == "" {
		return errors.New("missing app param")
	}
	if urlfields.Country == "" {
		return errors.New("missing country param")
	}
	if urlfields.OS == "" {
		return errors.New("missing os param")
	}

	return nil
}

func rulesApplicable(rules models.RuleSet, urlfields models.UrlFields) bool {
	v1 := applyRule(urlfields.App, rules.ExcludeApp, rules.IncludeApp)
	v2 := applyRule(urlfields.OS, rules.ExcludeOS, rules.IncludeOS)
	v3 := applyRule(urlfields.Country, rules.ExcludeCountry, rules.IncludeCountry)

	return v1 && v2 && v3
}

func applyRule(urlField string, exclude, include []string) bool {
	if slices.Contains(exclude, strings.ToLower(urlField)) {
		return false
	}
	if len(include) > 0 {
		return slices.Contains(include, strings.ToLower(urlField))
	}

	return true
}
