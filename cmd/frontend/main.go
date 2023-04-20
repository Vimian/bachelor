package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/casperfj/bachelor/cmd/frontend/config"
	"github.com/casperfj/bachelor/cmd/frontend/handlers"
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

	/*
		// Initialize repository
		repo, err := repository.NewRepository(conf)
		if err != nil {
			panic("failed to initialize repository: " + err.Error())
		}
	*/

	// Initialize Handlers
	handler := handlers.NewHandler(conf)

	// Initialize router
	router := gin.Default()

	// Setup routes
	router.GET("/", handler.Homepage)
	router.GET("/admin", handler.Adminpage)

	// Start HTTP server
	var address string = conf.Server.Host + ":" + fmt.Sprint(conf.Server.Port)
	log.Printf("starting user service on: %s", address)
	log.Fatalf("server exited with error: %s", http.ListenAndServe(address, router).Error())
}
