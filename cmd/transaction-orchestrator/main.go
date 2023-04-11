package main

import (
	_ "embed"

	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/cache"
	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/config"
	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/handlers"
	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/queue"
	commonConfig "github.com/casperfj/bachelor/pkg/common/config"
)

//go:embed config/config.yaml
var configFile []byte

func main() {
	// Load configuration
	confRaw, err := commonConfig.LoadConfig(configFile, &config.Configuration{})
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}
	conf, ok := confRaw.(*config.Configuration)
	if !ok {
		panic("failed to cast configuration to *config.Configuration")
	}

	// Initialize queue
	queue, err := queue.NewQueue(conf)
	if err != nil {
		panic("failed to initialize queue: " + err.Error())
	}

	// Close queue on exit
	defer queue.Connection.Close()
	defer queue.Channel.Close()

	// Initialize Handlers
	handler := handlers.NewHandler(conf)

	// Initialize cache
	cache, err := cache.NewCache(conf)
	if err != nil {
		panic("failed to initialize cache: " + err.Error())
	}

	// Subscribe to queue
	queue.SubscribeToTransaction(handler, cache, conf)
}
