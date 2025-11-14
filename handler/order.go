package handler

import (
    "log"
    "fmt"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
  "github.com/golang-jwt/jwt/v5"

    "codeberg.org/noreng-br/models"
)

func (h *Handler) CreateOrder(c echo.Context) error {
  orderRequest := new(models.OrderRequest)
  if err := c.Bind(orderRequest); err != nil {
    return c.String(http.StatusBadRequest, "invalid json format")
  }

  ctx := c.Request().Context()

  reqId, err := h.Service.Repository.OrderRepo.CreateOrder(ctx, *orderRequest)
  if err != nil {
    return c.String(http.StatusInternalServerError, "internal error")
  }

  log.Println("Create order------")

  return c.String(http.StatusCreated, fmt.Sprintf("New order %d created", reqId))
}

func (h *Handler) GetUserOrders(c echo.Context) error {
  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*models.JWTCustomClaims)

  userId := claims.UserID

  uId, err := strconv.Atoi(userId)
  if err != nil {
    return c.String(http.StatusBadRequest, "user_id wrong format")
  }

  ctx := c.Request().Context()

  userOrders, err := h.Service.Repository.OrderRepo.GetUserOrders(ctx, uId)
  if err != nil {
    return c.String(http.StatusInternalServerError, "internal server error")
  }

  return c.JSON(http.StatusOK, userOrders)
}

func (h *Handler) GetAllOrders(c echo.Context) error {
	// Simple JSON response example
  orders, err := h.Service.Repository.OrderRepo.ListOrders(c.Request().Context())
  if err != nil {
      c.String(http.StatusInternalServerError, "Error listing users")
  }

  return c.JSON(http.StatusOK, orders)
}
