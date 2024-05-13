package api

import (
	"payment-system-four/internal/util"

	"github.com/gin-gonic/gin"
)

// Readiness is to check if server is up
func (u *HTTPHandler) Readiness(c *gin.Context) {
	data := "server is up and running"

	// healthcheck
	util.Response(c, "Ready to go", 200, data, nil)
}
