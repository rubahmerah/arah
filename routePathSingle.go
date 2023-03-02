package arah

import (
	"github.com/labstack/echo/v4"
)

// Route path single
type routePathSingle struct {
	route *route
}

func (r routePathSingle) Echo() *echo.Echo {
	return r.route.echo
}

// func (r routePathSingle) Use(middleware ...echo.MiddlewareFunc) {
// 	r.route.echo.Use(middleware...)
// }

func (r routePathSingle) Group(groupName string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc) {
	g := routePathGroup{route: &group{
		host: r.route.host,
		echo: r.Echo().Group(groupName, m...),
	}}
	f(g)
}

func (r routePathSingle) Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().GET(prefix, f, m...),
	}
	return &rp
}

func (r routePathSingle) Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().POST(prefix, f, m...),
	}
	return &rp
}

func (r routePathSingle) Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().PUT(prefix, f, m...),
	}
	return &rp
}

func (r routePathSingle) Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().PATCH(prefix, f, m...),
	}
	return &rp
}

func (r routePathSingle) Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath {
	rp := routePath{
		host:  r.route.host,
		route: r.Echo().DELETE(prefix, f, m...),
	}
	return &rp
}

func (r routePathSingle) Any(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) []*routePath {
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
