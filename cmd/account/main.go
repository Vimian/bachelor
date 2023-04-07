package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/casperfj/bachelor/cmd/account/config"
	"github.com/casperfj/bachelor/cmd/account/handlers"
	"github.com/casperfj/bachelor/cmd/account/repository"
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
	router.POST("/account/", handler.CreateAccount)
	router.GET("/account/:id", handler.GetAccount)
	router.GET("/account/:id/balance/", handler.GetBalance)
	router.PUT("/account/:id/balance/", handler.UpdateBalance)
	router.GET("/accounts/:ownerid", handler.GetAccounts)

	// Start HTTP server
	var address string = conf.Server.Host + ":" + fmt.Sprint(conf.Server.Port)
	log.Printf("starting account service on: %s", address)
	log.Fatalf("server exited with error: %s", http.ListenAndServe(address, router).Error())
}
