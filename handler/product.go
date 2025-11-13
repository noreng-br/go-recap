package handler

import (
    "log"
    "net/http"

    "github.com/labstack/echo/v4"

    "codeberg.org/noreng-br/models"
)

func (h *Handler) CreateProduct(c echo.Context) error {
  productDto := new(models.ProductDTO)
  if err := c.Bind(productDto); err != nil {
    return c.String(http.StatusBadRequest, "Invalid json format")
  }
  
  var product models.Product 

  product.Name = productDto.Name
  product.Description = productDto.Description
  product.Price = productDto.Price

  product, err := h.Service.Repository.ProductRepo.CreateProduct(
    c.Request().Context(),
    product)

  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  log.Println("In Create Product handler")

  return c.JSON(http.StatusCreated, product)
}
