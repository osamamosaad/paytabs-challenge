package utils

import "github.com/labstack/echo/v4"

func WriteError(c echo.Context, status int, msg string) error {
	errorResponse := struct {
		Status  int
		Message string
	}{status, msg}

	return c.JSON(status, errorResponse)
}
