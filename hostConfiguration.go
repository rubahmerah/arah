package arah

import (
	"net"

	"github.com/labstack/echo/v4"
)

// Host configuration
type HostConfiguration struct {
	Debug            bool
	HideBanner       bool
	DisableHTTP2     bool
	ReadTimeout      int64
	WriteTimeout     int64
	Listener         *net.Listener
	Validator        *echo.Validator
	Binder           *echo.Binder
	Renderer         *echo.Renderer
	HTTPErrorHandler func(err error, c echo.Context)
}
