// pkg/handlers/handler.go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
}
