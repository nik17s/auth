package handler

import "github.com/nik17s/auth/pkg/service"

type Handler struct {
	service service.Implementation
}

func NewHandler(service service.Implementation) *Handler {
	return &Handler{
		service: service,
	}
}
