package health

import (
	"anagramm/pkg/logger"
	"anagramm/pkg/reply"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Provide(NewHealthHandler)

type Param struct {
	fx.In
	Logger logger.ILogger
}

type Handler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger logger.ILogger
}

func NewHealthHandler(p Param) Handler {
	return &handler{
		logger: p.Logger,
	}
}

func (h *handler) Health(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("%s", "health")
	reply.JSON(w, 200, "service is live")
}
