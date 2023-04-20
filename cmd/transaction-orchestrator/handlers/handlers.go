package handlers

import (
	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/config"
	"github.com/casperfj/bachelor/pkg/transaction"
)

type Handler struct {
	conf *config.Configuration
}

type Handlers struct {
	ProcessTransaction func(transaction *transaction.Transaction, configuration *config.Configuration)
}

func NewHandler(conf *config.Configuration) *Handlers {
	// Initialize handlers
	h := &Handler{
		conf: conf,
	}

	// Return handlers
	return &Handlers{
		ProcessTransaction: h.ProcessTransaction,
	}
}
