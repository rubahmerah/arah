package route

import "github.com/labstack/echo/v4"

type RoutePathInterface interface {
	Group(group string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc)
	Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Any(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) []*routePath
}

type RouteInterface interface {
	Create(RoutePathInterface)
}
