package handlers

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/adminpage.html
var adminpageFile []byte

func (h *Handler) Adminpage(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", adminpageFile)
}
