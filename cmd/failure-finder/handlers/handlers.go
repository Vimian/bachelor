package handlers

import (
	"github.com/casperfj/bachelor/cmd/failure-finder/config"
	"github.com/casperfj/bachelor/cmd/failure-finder/repository"
	"github.com/google/uuid"
)

type Handler struct {
	conf *config.Configuration
	repo *repository.Repository
}

type Handlers struct {
	ProcessAccount func(uuid.UUID)
}

func NewHandler(conf *config.Configuration, repo *repository.Repository) *Handlers {
	// Initialize handlers
	h := &Handler{
		conf: conf,
		repo: repo,
	}

	// Return handlers
	return &Handlers{
		ProcessAccount: h.ProcessAccount,
	}
}
