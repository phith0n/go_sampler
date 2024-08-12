package web

import "go_sampler/providers/config"

func NewAPI(cfg *config.Config) *APIHandler {
	return &APIHandler{
		config: cfg,
	}
}

type APIHandler struct {
	config *config.Config
}
