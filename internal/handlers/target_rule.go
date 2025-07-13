package handlers

import (
	"github.com/as-ifn-at/REST/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

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

}

func (h *ruleHandler) Save(ctx *gin.Context) {

}
