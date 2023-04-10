package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/casperfj/bachelor/cmd/transaction-history/config"
	"github.com/casperfj/bachelor/cmd/transaction-history/handlers"
	"github.com/casperfj/bachelor/cmd/transaction-history/repository"
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

	// Initialize repository
	repo, err := repository.NewRepository(conf)
	if err != nil {
		panic("failed to initialize repository: " + err.Error())
	}

	// Initialize Handlers
	handler := handlers.NewHandler(conf, repo)

	// Initialize router
	router := gin.Default()

	// Setup routes
	router.POST("/transaction-history/", handler.CreateTransactionHistory)
	router.GET("/transaction-history/:id", handler.GetTransactionHistory)
	router.GET("/transaction-history/:id/status/", handler.GetTransactionHistoryStatus)
	router.GET("/transaction-histories/:accountid", handler.GetTransactionHistories)
	router.GET("/transaction/:transactionid/status/", handler.GetTransactionHistoryStatusByTransactionID)
	router.PUT("/transaction/:transactionid/status/", handler.UpdateTransactionHistoryStatusByTransactionID)
	router.GET("/statuses/", handler.GetStatuses)
	router.GET("/types/", handler.GetTypes)

	// Start HTTP server
	var address string = conf.Server.Host + ":" + fmt.Sprint(conf.Server.Port)
	log.Printf("starting transaction history service on: %s", address)
	log.Fatalf("server exited with error: %s", http.ListenAndServe(address, router).Error())
}
