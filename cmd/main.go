package main

import (
	"anagramm/cmd/handlers/anagramm"
	"anagramm/cmd/handlers/health"
	"anagramm/cmd/router"
	"anagramm/internal/anagramm_service"
	"anagramm/pkg/logger"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		logger.Module,
		router.Module,
		health.Module,
		anagramm.Module,
		anagramm_service.Module,
	)

	app.Run()
}
