package handler

import (
	"ShowGoOn/backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SlideHandler struct {
	service *service.SlideService
}

func NewSlideHandler(s *service.SlideService) *SlideHandler {
	return &SlideHandler{service: s}
}

func (h *SlideHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/slides", h.CreateSlide)
	r.GET("/slides/:id", h.GetSlide)
	r.PUT("/slides/:id", h.UpdateSlide)
	r.DELETE("/slides/:id", h.DeleteSlide)
	r.GET("/slides", h.GetAllSlides) // New route for getting all slides
}

func (h *SlideHandler) CreateSlide(c *gin.Context) {
	var slide service.Slide
	if err := c.ShouldBindJSON(&slide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateSlide(&slide); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, slide)
}

func (h *SlideHandler) GetSlide(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	slide, err := h.service.GetSlideByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slide)
}

func (h *SlideHandler) UpdateSlide(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var slide service.Slide
	if err := c.ShouldBindJSON(&slide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slide.ID = intID
	if err := h.service.UpdateSlide(&slide); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slide)
}

func (h *SlideHandler) DeleteSlide(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	if err := h.service.DeleteSlide(intID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Slide deleted successfully"})
}

func (h *SlideHandler) GetAllSlides(c *gin.Context) {
	slides, err := h.service.GetAllSlides()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slides)
}
