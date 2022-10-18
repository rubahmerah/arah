package route

import "github.com/labstack/echo/v4"

type route struct {
	host *host
	echo *echo.Echo
}
