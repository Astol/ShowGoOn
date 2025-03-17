package handler

import (
	"ShowGoOn/backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.SlideService
}

func NewSlideHandler(s *service.SlideService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/slides", h.CreateItem)
	r.GET("/slides/:id", h.GetItem)
	r.PUT("/slides/:id", h.UpdateItem)
	r.DELETE("/slides/:id", h.DeleteItem)
}

func (h *Handler) CreateItem(c *gin.Context) {
	// Implementation for creating an item
	c.JSON(http.StatusOK, gin.H{"message": "Create item"})
}

func (h *Handler) GetItem(c *gin.Context) {
	// Implementation for getting an item by ID
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get item", "id": id})
}

func (h *Handler) UpdateItem(c *gin.Context) {
	// Implementation for updating an item by ID
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update item", "id": id})
}

func (h *Handler) DeleteItem(c *gin.Context) {
	// Implementation for deleting an item by ID
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete item", "id": id})
}
