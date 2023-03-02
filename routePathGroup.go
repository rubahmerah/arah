package arah

import (
	"github.com/labstack/echo/v4"
)

// Route Path Grouping
type routePathGroup struct {
	route *group
}

func (r routePathGroup) Echo() *echo.Group {
	return r.route.echo
}

// func (r routePathGroup) Use(middleware ...echo.MiddlewareFunc) {
// 	r.route.echo.Use(middleware...)
// }

func (r routePathGroup) Group(groupName string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc) {
	g := routePathGroup{route: &group{
		host: r.route.host,
		echo: r.Echo().Group(groupName, m...),
	}}
	f(g)
}

func (r routePathGroup) Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().GET(prefix, f, m...),
	}
	return &rp
}

func (r routePathGroup) Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().POST(prefix, f, m...),
	}
	return &rp
}

func (r routePathGroup) Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().PUT(prefix, f, m...),
	}
	return &rp
}

func (r routePathGroup) Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().PATCH(prefix, f, m...),
	}
	return &rp
}

func (r routePathGroup) Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().DELETE(prefix, f, m...),
	}
	return &rp
}

func (r routePathGroup) Any(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) []*routePath {
	rp := []*routePath{}
	ro := r.Echo().Any(prefix, f, m...)
	for _, v := range ro {
		rp = append(rp, &routePath{
			host:  r.route.host,
			route: v,
		})
	}
	return rp
}
