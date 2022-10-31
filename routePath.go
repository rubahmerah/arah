package arah

import "github.com/labstack/echo/v4"

// Route path
type routePath struct {
	host   *host
	route  *echo.Route
	path   string
	name   string
	method string
}

func (rp *routePath) Name(name string) {
	rp.route.Name = name
	rp.name = name
	routeLists.routes[name] = rp.host
}
