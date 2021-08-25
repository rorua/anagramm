package anagramm

import (
	"anagramm/internal/anagramm_service"
	"anagramm/pkg/logger"
	"anagramm/pkg/reply"
	"encoding/json"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Provide(New)

type Param struct {
	fx.In
	Logger          logger.ILogger
	AnagrammService anagramm_service.Service
}

type Handler interface {
	Load(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger          logger.ILogger
	anagrammService anagramm_service.Service
}

func New(p Param) Handler {
	return &handler{
		logger:          p.Logger,
		anagrammService: p.AnagrammService,
	}
}

func (h *handler) Load(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("%s", "load")

	var request []string

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error("err on parsing req %v", err)
		reply.JSON(w, 400, err)
	}

	err = h.anagrammService.Load(request)
	if err != nil {
		h.logger.Error("err on saving dict %v", err)
		reply.JSON(w, 400, err)
	}

	reply.JSON(w, 200, nil)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("%s", "Get")

	word := r.URL.Query().Get("word")

	dict, err := h.anagrammService.Get(word)
	if err != nil {
		h.logger.Error("err on saving dict %v", err)
		reply.JSON(w, 400, err)
	}

	if len(dict) == 0 {
		reply.JSON(w, 200, nil)
		return
	}

	reply.JSON(w, 200, dict)
}
