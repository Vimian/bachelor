package handlers

import (
	"github.com/casperfj/bachelor/cmd/frontend/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	conf *config.Configuration
	//repo *repository.Repository
}

type Handlers struct {
	Homepage  func(c *gin.Context)
	Adminpage func(c *gin.Context)
}

func NewHandler(conf *config.Configuration /*, repo *repository.Repository*/) *Handlers {
	// Initialize handlers
	h := &Handler{
		conf: conf,
		//repo: repo,
	}

	// Return handlers
	return &Handlers{
		Homepage:  h.Homepage,
		Adminpage: h.Adminpage,
	}
}
