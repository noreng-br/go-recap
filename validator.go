package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/go-playground/validator/v10"
)
// CustomValidator holds the validator instance
type CustomValidator struct {
    validator *validator.Validate
}
// Validate implements the echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        // Optionally customize the error message structure
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
