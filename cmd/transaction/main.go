package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/casperfj/bachelor/cmd/transaction/config"
	"github.com/casperfj/bachelor/cmd/transaction/handlers"
	"github.com/casperfj/bachelor/cmd/transaction/queue"
	commonConfig "github.com/casperfj/bachelor/pkg/common/config"
	"github.com/gin-gonic/gin"
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
	handler := handlers.NewHandler(queue)

	// Initialize router
	router := gin.Default()

	// Setup routes
	router.POST("/transaction/", handler.CreateTransaction)

	// Start HTTP server
	var address string = conf.Server.Host + ":" + fmt.Sprint(conf.Server.Port)
	log.Printf("starting transaction service on: %s", address)
	log.Fatalf("server exited with error: %s", http.ListenAndServe(address, router).Error())
}
