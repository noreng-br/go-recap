// handler/handler.go
package handler

import (
	"github.com/labstack/echo/v4"
)

// InitRoutes sets up all API routes on the provided Echo instance.
func InitRoutes(e *echo.Echo, h *Handler) {
	// Group routes under a common prefix, like /v1
	v1 := e.Group("/v1")

	// --- User Routes ---
	// Since GetUsers and GetUserByID are in the same package (handler),
	// they can be called directly without any prefix.
	v1.GET("/users", h.GetUsers)
	v1.GET("/users/:id", h.GetUserByID)
	v1.POST("/users", h.CreateUser)

	// You can add other resources' routes here:
	// v1.GET("/products", GetProducts)
}
