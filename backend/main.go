package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"BakeMeets/api/handlers"
	"BakeMeets/repository"
	"BakeMeets/config"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"net/http"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database connection
	cfg := config.GetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Error connecting to database:", err)
	}

	// Repository
	breadRepo := repository.BreadRepository{DB: db}
	userRepo := repository.UserRepository{DB: db}

	// Handlers
	breadHandler := handlers.BreadHandler{Br: &breadRepo}
	userHandler := handlers.UserHandler{Ur: &userRepo}

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to BreadMeets!")
	})
	e.GET("/breads", breadHandler.GetAllBreads)
	e.GET("/breads/:id", breadHandler.GetBreadByID)
	e.POST("/breads", breadHandler.CreateBread)
	e.PUT("/breads/edit/:id", breadHandler.UpdateBread)
	e.DELETE("/breads/:id", breadHandler.DeleteBread)

	e.POST("/signup", userHandler.SignUp)
	e.POST("/login", userHandler.Login)
	e.POST("/logout", userHandler.Logout)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
