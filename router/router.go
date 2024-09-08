package router

import (
	"log/slog"

	"github.com/fentezi/unit-converter/handler"
)

type Router struct {
	log     *slog.Logger
	handler *handler.Handler
}

func NewRouter(log *slog.Logger, handler *handler.Handler) *Router {
	return &Router{
		log:     log,
		handler: handler,
	}
}

func (r *Router) initRouter() {
}

func (r *Router) StartServer() {
	
}