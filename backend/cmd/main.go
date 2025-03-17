package main

import (
	"ShowGoOn/backend/internal/handler"
	"ShowGoOn/backend/internal/repository"
	"ShowGoOn/backend/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	// DB
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Router
	r := gin.Default()

	// Set trusted proxies TODO: move to a config file, now its localhost
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("failed to set trusted proxies:", err)
	}

	// Repositories
	repo := repository.NewSlideRepository(db)

	// Services
	svc := service.NewSlideService(repo)

	// Handlers
	handler.NewSlideHandler(svc).RegisterRoutes(r)

	// Run
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
