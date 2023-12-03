package internal

import "StorageService/internal/service"

type API struct {
	EventPublisher *service.EventPublisher
}

func NewApi(service.EventPublisher) *API {
	return &API{}
}
