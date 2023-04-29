package handlers

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/homepage.html
var homepageFile []byte

func (h *Handler) Homepage(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", homepageFile)
}
