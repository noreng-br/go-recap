
package handler

import (
    "log"
    "net/http"

    "github.com/labstack/echo/v4"

    "codeberg.org/noreng-br/models"
)

func (h *Handler) CreateCategory(c echo.Context) error {
  categoryDto := new(models.CategoryDTO)
  if err := c.Bind(categoryDto); err != nil {
    return c.String(http.StatusBadRequest, "Invalid json format")
  }

  if categoryDto.Name == "" {
    return c.String(http.StatusBadRequest, "name required")
  }
  

  category, err := h.Service.Repository.CategoryRepo.CreateCategory(
    c.Request().Context(),
    categoryDto.Name)

  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  log.Println("In Create Category handler")

  return c.JSON(http.StatusCreated, category)
}
