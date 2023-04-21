package main

import (
	_ "embed"

	"github.com/casperfj/bachelor/cmd/failure-finder/config"
	"github.com/casperfj/bachelor/cmd/failure-finder/handlers"
	"github.com/casperfj/bachelor/cmd/failure-finder/queue"
	"github.com/casperfj/bachelor/cmd/failure-finder/repository"
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

	// Initialize repository
	repo, err := repository.NewRepository(conf)
	if err != nil {
		panic("failed to initialize repository: " + err.Error())
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
	handler := handlers.NewHandler(conf, repo)

	// Subscribe to queue
	queue.SubscribeToAccounts(handler, conf)
}
