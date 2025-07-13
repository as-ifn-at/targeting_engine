package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/as-ifn-at/targeting_engine/internal/config"
	"github.com/as-ifn-at/targeting_engine/internal/routes"
	"github.com/rs/zerolog"
)

func main() {

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	config := config.Load()
	router := routes.NewRouter(config, logger).SetRouters()
	listenPort := fmt.Sprintf(":%v", config.Port)
	httpServer := &http.Server{
		Addr:    listenPort,
		Handler: router,
	}

	logger.Info().Msg(fmt.Sprintf("starting server on port %v", listenPort))
	if err := httpServer.ListenAndServe(); err != nil {
		logger.Error().Msg(fmt.Sprintf("unable to start server on port %v", listenPort))
		panic(err)
	}
}
