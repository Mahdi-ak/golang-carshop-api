package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health handles the GET endpoint
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "working!")
	return
}

// HealthPost handles the POST endpoint
func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "working! Post")
	return
}

// HealthPostById handles the POST by id endpoint
func (h *HealthHandler) HealthPostById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("working! Post by id: %s", id))
	return
}
