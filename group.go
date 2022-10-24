package route

import "github.com/labstack/echo/v4"

// Route for grouping
type group struct {
	host *host
	echo *echo.Group
}
