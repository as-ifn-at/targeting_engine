package handlers

import (
	"github.com/as-ifn-at/REST/internal/config"
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

}

func (h *deliverHandler) Save(ctx *gin.Context) {

}
