package handler

import (
	"log/slog"

	"github.com/fentezi/unit-converter/service"
)

type Handler struct {
	log     *slog.Logger
	service *service.Service
}

func NewHandler(log *slog.Logger, service *service.Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}
