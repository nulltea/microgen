package cmd

import "github.com/timoth-y/microgen/core/service"

type Handler struct {
	service service.OperationService
}

func NewCommander(operations service.OperationService) *Handler {
	return &Handler{
		operations,
	}
}