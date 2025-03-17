package main

import (
	"ShowGoOn/backend/internal/handler"
	"ShowGoOn/backend/internal/repository"
	"ShowGoOn/backend/internal/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	// Initialize the database connection
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	// Initialize repository
	repo := repository.NewSlideRepository(db)

	// Initialize service
	svc := service.NewSlideService(repo)

	// Initialize handler
	r := gin.Default()
	handler.NewSlideHandler(svc).RegisterRoutes(r)

	return r
}

func TestCreateSlide(t *testing.T) {
	router := setupRouter()

	slide := repository.Slide{
		Title: "Test Slide",
		Text:  "This is a test slide",
	}
	body, _ := json.Marshal(slide)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/slides", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Create item")
}

func TestGetSlide(t *testing.T) {
	router := setupRouter()

	// First, create a slide to ensure there is something to get
	slide := repository.Slide{
		Title: "Test Slide",
		Text:  "This is a test slide",
	}
	body, _ := json.Marshal(slide)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/slides", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Now, get the slide
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/slides/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Get item")
}

func TestUpdateSlide(t *testing.T) {
	router := setupRouter()

	// First, create a slide to ensure there is something to update
	slide := repository.Slide{
		Title: "Test Slide",
		Text:  "This is a test slide",
	}
	body, _ := json.Marshal(slide)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/slides", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Now, update the slide
	updatedSlide := repository.Slide{
		Title: "Updated Slide",
		Text:  "This is an updated slide",
	}
	updatedBody, _ := json.Marshal(updatedSlide)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/slides/1", bytes.NewBuffer(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Update item")
}

func TestDeleteSlide(t *testing.T) {
	router := setupRouter()

	// First, create a slide to ensure there is something to delete
	slide := repository.Slide{
		Title: "Test Slide",
		Text:  "This is a test slide",
	}
	body, _ := json.Marshal(slide)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/slides", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Now, delete the slide
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/slides/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Delete item")
}
