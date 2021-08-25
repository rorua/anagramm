package router

import (
	"anagramm/cmd/handlers/anagramm"
	"anagramm/cmd/handlers/health"
	"anagramm/pkg/logger"
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Invoke(
		NewRouter,
	),
)

type Params struct {
	fx.In
	Lifecycle      fx.Lifecycle
	Logger         logger.ILogger
	HealthHandler  health.Handler
	AnagramHandler anagramm.Handler
}

func NewRouter(params Params) {

	router := mux.NewRouter()

	router.HandleFunc("/health", params.HealthHandler.Health).Methods("GET")

	router.HandleFunc("/load", params.AnagramHandler.Load).Methods("POST")
	router.HandleFunc("/get", params.AnagramHandler.Get).Methods("GET")

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	params.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				params.Logger.Info("Application started")
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				params.Logger.Info("Application stopped")
				return server.Shutdown(ctx)
			},
		},
	)
}
