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


func (h *Handler) GetProducts(c echo.Context) error {
  ctx := c.Request().Context()

  products, err := h.Service.Repository.ProductRepo.GetProducts(ctx)
  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  return c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProduct(c echo.Context) error {
  ctx := c.Request().Context()
  idx := c.QueryParam("id")

  product, err := h.Service.Repository.ProductRepo.GetProductById(ctx, idx)
  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  return c.JSON(http.StatusOK, product)
}

func (h *Handler) AddProductCategories(c echo.Context) error {
  pCatDto := new(models.ProductCategoryDTO)
  if err := c.Bind(pCatDto); err != nil {
    return c.String(http.StatusBadRequest, "Invalid json format")
  }

  // empty array for category
  if len(pCatDto.CategoryIDs) == 0 {
    return c.String(http.StatusBadRequest, "Categories required")
  }

  ctx := c.Request().Context()

  err := h.Service.Repository.ProductRepo.AddCategoriesToProduct(ctx, pCatDto.ProductID, pCatDto.CategoryIDs)
  if err !=  nil {
    return c.String(http.StatusInternalServerError, "Internal server error")
  }

  return c.String(http.StatusOK, "Categories of the product have been updated")
}

func (h *Handler) UpdateProduct(c echo.Context) error {
  productDto := new(models.ProductDTO)
  if err := c.Bind(productDto); err != nil {
    return c.String(http.StatusBadRequest, "Invalid json format")
  }

  idx := c.QueryParam("id")
  
  var product models.Product 

  product.Name = productDto.Name
  product.Description = productDto.Description
  product.Price = productDto.Price

  product, err := h.Service.Repository.ProductRepo.UpdateProduct(
    c.Request().Context(),
    idx,
    product,
  )

  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  log.Println("In update product handler")

  return c.JSON(http.StatusCreated, product)
}

func (h *Handler) DeleteProduct(c echo.Context) error {
  ctx := c.Request().Context()
  idx := c.QueryParam("id")

  err := h.Service.Repository.ProductRepo.DeleteProduct(ctx, idx)
  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  return c.JSON(http.StatusOK, "successfully deleted product")
}
