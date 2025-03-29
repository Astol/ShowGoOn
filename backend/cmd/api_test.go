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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	// Initialize the database connection
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	// Automatically migrate the schema for the Slide and Settings models
	db.AutoMigrate(&repository.Slide{}, &repository.Settings{})

	// Initialize repositories
	slideRepo := repository.NewSlideRepository(db)
	settingsRepo := repository.NewSettingsRepository(db)

	// Initialize services
	slideSvc := service.NewSlideService(slideRepo)
	settingsSvc := service.NewSettingsService(settingsRepo)

	// Initialize handlers
	r := gin.Default()
	handler.NewSlideHandler(slideSvc).RegisterRoutes(r)
	handler.NewSettingsHandler(settingsSvc).RegisterRoutes(r)

	return r
}

func TestCreateSlide(t *testing.T) {
	router := setupRouter()

	slide := service.Slide{
		Title:     "Test Slide",
		Text:      "This is a test slide",
		MediaPath: "/uploads/test.jpg",
		MediaType: 1, // ResourceTypeImage
	}
	body, _ := json.Marshal(slide)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/slides", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Test Slide")
}

func TestGetSlide(t *testing.T) {
	router := setupRouter()

	// First, create a slide to ensure there is something to get
	slide := service.Slide{
		Title:     "Test Slide",
		Text:      "This is a test slide",
		MediaPath: "/uploads/test.jpg",
		MediaType: 1, // ResourceTypeImage
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
	assert.Contains(t, w.Body.String(), "Test Slide")
}

func TestUpdateSlide(t *testing.T) {
	router := setupRouter()

	// First, create a slide to ensure there is something to update
	slide := service.Slide{
		Title:     "Test Slide",
		Text:      "This is a test slide",
		MediaPath: "/uploads/test.jpg",
		MediaType: 1, // ResourceTypeImage
	}
	body, _ := json.Marshal(slide)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/slides", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Now, update the slide
	updatedSlide := service.Slide{
		Title:     "Updated Slide",
		Text:      "This is an updated slide",
		MediaPath: "/uploads/updated.jpg",
		MediaType: 2, // ResourceTypeVideo
	}
	updatedBody, _ := json.Marshal(updatedSlide)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/slides/1", bytes.NewBuffer(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Slide")
}

func TestDeleteSlide(t *testing.T) {
	router := setupRouter()

	// First, create a slide to ensure there is something to delete
	slide := service.Slide{
		Title:     "Test Slide",
		Text:      "This is a test slide",
		MediaPath: "/uploads/test.jpg",
		MediaType: 1, // ResourceTypeImage
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
	assert.Contains(t, w.Body.String(), "Slide deleted successfully")
}

func TestGetAllSlides(t *testing.T) {
	router := setupRouter()

	// Create multiple slides
	slides := []service.Slide{
		{Title: "Slide 1", Text: "Text 1", MediaPath: "/uploads/slide1.jpg", MediaType: 1},
		{Title: "Slide 2", Text: "Text 2", MediaPath: "/uploads/slide2.jpg", MediaType: 2},
	}
	for _, slide := range slides {
		body, _ := json.Marshal(slide)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/slides", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
	}

	// Fetch all slides
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/slides", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Slide 1")
	assert.Contains(t, w.Body.String(), "Slide 2")
}

func TestCreateSettings(t *testing.T) {
	router := setupRouter()

	settings := service.Settings{
		SlideDuration:   5,
		SlideTransition: 2,
		UpdateDatetime:  time.Now(),
	}
	body, _ := json.Marshal(settings)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/settings", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "SlideDuration")
}

func TestGetSettings(t *testing.T) {
	router := setupRouter()

	// First, create settings to ensure there is something to get
	settings := service.Settings{
		SlideDuration:   5,
		SlideTransition: 2,
		UpdateDatetime:  time.Now(),
	}
	body, _ := json.Marshal(settings)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/settings", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Now, get the settings
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/settings/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "SlideDuration")
}

func TestUpdateSettings(t *testing.T) {
	router := setupRouter()

	// First, create settings to ensure there is something to update
	settings := service.Settings{
		SlideDuration:   5,
		SlideTransition: 2,
		UpdateDatetime:  time.Now(),
	}
	body, _ := json.Marshal(settings)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/settings", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Now, update the settings
	updatedSettings := service.Settings{
		SlideDuration:   10,
		SlideTransition: 3,
		UpdateDatetime:  time.Now(),
	}
	updatedBody, _ := json.Marshal(updatedSettings)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/settings/1", bytes.NewBuffer(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "SlideDuration")
	assert.Contains(t, w.Body.String(), "10")
}

func TestDeleteSettings(t *testing.T) {
	router := setupRouter()

	// First, create settings to ensure there is something to delete
	settings := service.Settings{
		SlideDuration:   5,
		SlideTransition: 2,
		UpdateDatetime:  time.Now(),
	}
	body, _ := json.Marshal(settings)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/settings", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Now, delete the settings
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/settings/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
