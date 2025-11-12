// handler/user.go
package handler

import (
  "fmt"
	"net/http"

	"github.com/labstack/echo/v4"

  "codeberg.org/noreng-br/service"
  "codeberg.org/noreng-br/models"
)

type Handler struct {
  Service service.Service
}

func NewHandler(s service.Service) (*Handler, error) {
  return &Handler{
    Service: s,
  }, nil
}  

// GetUsers handles GET /api/v1/users
func (h *Handler) GetUsers(c echo.Context) error {
	// Simple JSON response example
	users := []string{"Alice", "Bob", "Charlie"}
	return c.JSON(http.StatusOK, users)
}

// GetUserByID handles GET /api/v1/users/:id
func (h *Handler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	// Respond with the ID parameter
	return c.String(http.StatusOK, "Viewing user: "+id)
}

func (h *Handler) CreateUser(c echo.Context) error {
  u := new(models.User)
  if err := c.Bind(u); err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Data")
  }

  if err := c.Validate(u); err != nil {
      // Returns a 400 Bad Request with validation errors
      return err 
  }

  // do not let user create admin through the register route
  u.IsAdmin = false

  fmt.Println("=============================In Handler=====================+++++")
  fmt.Println(u)
  fmt.Println("============================================================+++++")

  usr, err := h.Service.CreateUser(c.Request().Context(), *u)
  if err != nil {
    fmt.Println("=============================")
    fmt.Println(usr)
    fmt.Println(err.Error())
    fmt.Println("=============================")
    fmt.Println("=============================")
    return c.JSON(http.StatusInternalServerError, nil)
  }

  fmt.Println(usr)
	// Respond with the ID parameter
	return c.JSON(http.StatusCreated, u)
}
