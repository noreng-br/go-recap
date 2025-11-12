package handler

import (
  "net/http"


  "log"
  "github.com/labstack/echo/v4"

  "codeberg.org/noreng-br/models"
  "codeberg.org/noreng-br/service"
)

func (h *Handler) Login(c echo.Context) error {
  auth := new(models.Auth)
  if err := c.Bind(auth); err != nil {
    return c.String(http.StatusBadRequest, "Invalid json format")
  }

  token, err := h.Service.Login(c.Request().Context(), auth)
  if err != nil {
      log.Println("===================")
      log.Println(err.Error())
      if err == service.ErrNotFound {
        return c.JSON(http.StatusNotFound, "The user was not found")
      } else if err == service.ErrPassword {
          return c.JSON(http.StatusForbidden, "Not allowed")
      } else {
        return c.JSON(http.StatusInternalServerError, "An error ocurred")
      }
  }



	return c.JSON(http.StatusOK, token)
}
