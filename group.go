package route

import "github.com/labstack/echo/v4"

type group struct {
	host *host
	echo *echo.Group
}
