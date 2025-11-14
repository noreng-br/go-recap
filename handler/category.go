
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


func (h *Handler) GetCategories(c echo.Context) error {
	// Simple JSON response example
  categories, err := h.Service.Repository.CategoryRepo.ListCategories(c.Request().Context())
  if err != nil {
      c.String(http.StatusInternalServerError, "Error listing users")
  }

  return c.JSON(http.StatusOK, categories)
}

func (h *Handler) DeleteCategory(c echo.Context) error {
  ctx := c.Request().Context()
  idx := c.QueryParam("id")

  err := h.Service.Repository.CategoryRepo.DeleteCategory(ctx, idx)
  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  return c.String(http.StatusOK, "category was succesfully removed")
}

