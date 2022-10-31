package arah

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// The host
type host struct {
	hostname string
	echo     *echo.Echo
	config   HostConfiguration
}

// Setup configuration for host
func (r *host) prep() {
	r.echo.HideBanner = r.config.HideBanner
	r.echo.DisableHTTP2 = r.config.DisableHTTP2
	r.echo.Debug = r.config.Debug
	r.echo.Server.ReadTimeout = time.Duration(r.config.ReadTimeout)
	r.echo.Server.WriteTimeout = time.Duration(r.config.WriteTimeout)

	validator := r.config.Validator
	r.echo.Validator = validator

	listener := r.config.Listener
	r.echo.Listener = listener

	ipextractor := r.config.IPExtractor
	r.echo.IPExtractor = ipextractor
}

func (h *host) ServeHTTP(res *echo.Response, req *http.Request) {
	h.echo.ServeHTTP(res, req)
}
