package arah

import "github.com/labstack/echo/v4"

// Route
type route struct {
	host *host
	echo *echo.Echo
}
