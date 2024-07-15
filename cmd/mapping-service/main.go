package main

import (
	"context"
	handlers "mapping-service/pkg/handler"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	h := handlers.NewHandler()

	e := handlers.SetupRoutes(h)

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	// Custom Middleware

	// CORS Middleware
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Start server
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if err := e.Start(":" + port); err != nil {
			e.Logger.Fatal("Error starting Echo server:", err)
		}
	}()

	// Set up channel on which to receive SIGINT signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds
	<-quit
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		e.Logger.Fatal("Server shutdown failed:", err)
	}
	e.Logger.Info("Server gracefully stopped")

}
