package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/as-ifn-at/targeting_engine/internal/config"
	"github.com/as-ifn-at/targeting_engine/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var ErrViolatingIncExRule = errors.New("inclusion and exclusion rule cannot be applied at the same time")

type ruleHandler struct {
	Handler
	config config.Config
	logger zerolog.Logger
}

func NewRuleHandler(config config.Config, logger zerolog.Logger) Handler {
	return &ruleHandler{
		config: config,
		logger: logger,
	}
}

func (h *ruleHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, rule := range Rules {
		if rule.CampaignId == id {
			ctx.IndentedJSON(http.StatusOK, rule)
			return
		}
	}

	h.logger.Error().Msg(fmt.Sprintf("rule not found: %v", id))
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "rule not found"})
}

func (h *ruleHandler) Save(ctx *gin.Context) {
	var newTargetRules models.TargetRules
	if err := ctx.BindJSON(&newTargetRules); err != nil {
		h.logger.Error().Msg(fmt.Sprintf("error while unmarshaling rule details: %v", err.Error()))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTargetRules.Rules.Normalize()
	if _, ok := Campaigns[newTargetRules.CampaignId]; !ok {
		h.logger.Error().Msg(fmt.Sprintf("no campaign present with id: [%v] to apply rules", newTargetRules.CampaignId))
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "no campaign present to apply rules"})
		return
	}

	if err := validateRules(newTargetRules.Rules); err != nil {
		h.logger.Error().Msg(fmt.Sprintf("error applying rules to campaign; %v", err))
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Rules[newTargetRules.CampaignId] = newTargetRules
	ctx.IndentedJSON(http.StatusCreated, gin.H{"Rule created for campaign": newTargetRules.CampaignId})
}

func validateRules(rules models.RuleSet) error {
	valid := isValidRule(rules.IncludeCountry, rules.ExcludeCountry)
	if !valid {
		return ErrViolatingIncExRule
	}
	valid = isValidRule(rules.IncludeOS, rules.ExcludeOS)
	if !valid {
		return ErrViolatingIncExRule
	}
	valid = isValidRule(rules.IncludeApp, rules.ExcludeApp)
	if !valid {
		return ErrViolatingIncExRule
	}

	return nil
}

func isValidRule[T comparable](ip1, ip2 []T) bool {
	return !(len(ip1) > 0 && len(ip2) > 0)
}
