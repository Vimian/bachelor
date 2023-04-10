package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/casperfj/bachelor/cmd/user/config"
	"github.com/casperfj/bachelor/cmd/user/handlers"
	"github.com/casperfj/bachelor/cmd/user/repository"
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
	handler := handlers.NewHandler(repo)

	// Initialize router
	router := gin.Default()

	// Setup routes
	router.POST("/user/", handler.CreateUser)
	router.GET("/user/:id", handler.GetUser)

	// Start HTTP server
	var address string = conf.Server.Host + ":" + fmt.Sprint(conf.Server.Port)
	log.Printf("starting user service on: %s", address)
	log.Fatalf("server exited with error: %s", http.ListenAndServe(address, router).Error())
}
