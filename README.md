# Arah <sub>by 🦊</sub> 

Library for managing host and route HTTP based on Echo.

___

<br>

## Content

1. [Quickplay](#quickplay)
1. [Host](#host)
    - [Hostname](#hostname)
    - [Config](#config)
    - [List Route](#list-route)
1. [Route](#route)
1. [Start HTTP](#start-http)

<br>

___

<br>

## Quickplay

[Back to top](#content)

<br>

You need to setup hosts and routes, then start the HTTP server, for example.

```go
package main

import (
    "net/http"

    "github.com/iamaredfoxx/arah"
    "github.com/labstack/echo/v4"
)

func main() {
    arah.Register(

        // register host
        arah.HostHandler{
            Route: []arah.RouteInterface{
                exampleRoute{}, // initiate the route
            },
        },

    )

    // serve HTTP
    arah.Bind(
        arah.Start{
            Port: "8080",
        },
    )
}

// This is route
type exampleRoute struct {}

func (exampleRoute) Create(rp arah.RoutePathInterface) {
    rp.Get("/example", func (c echo.Context) error {
        return c.String(http.StatusOK, "example")
    })
}
```

<br>

___

<br>

## Host
[Back to top](#content)

<br>

Don't you think you need different host for different section too 😉

<br>

___

<br>

### Hostname
[Back to top](#content)

<br>

You can specify the hostname by filling in the hostname option.

> Default host is required, to do so you can fill the hostname config with "" (empty string)

```go
arah.HostHandler{
    Hostname: "book.example.com"
}
```

<br><br>

### Config
[Back to top](#content)

<br>

More configuration for the host that you can do.

```go
type HostConfiguration struct {
	Debug            bool
	HideBanner       bool
	DisableHTTP2     bool
	ReadTimeout      int64
	WriteTimeout     int64
	Listener         net.Listener
	Validator        echo.Validator
	Binder           echo.Binder
	Renderer         echo.Renderer
	IPExtractor      echo.IPExtractor
	Middleware       []echo.MiddlewareFunc
	HTTPErrorHandler func(err error, c echo.Context)
}
```
for more detail check out https://echo.labstack.com/guide/customization

<br><br>

### List Route
[Back to top](#content)

<br>

Register all your routes in here, like so.

```go
arah.HostHandler{
    Route: []arah.RouteInterface{
        exampleRoute{},
    }
}
```

<br>

___

<br>

## Route
[Back to top](#content)

<br>

This is all you can do for now.
```go
Group(group string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc)

Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath

Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath

Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath

Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath

Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath

Any(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) []*routePath

Use(middleware ...echo.MiddlewareFunc)
```

<br>

Use all of that like this
```go
type exampleRoute struct {}

func (exampleRoute) Create(rp arah.RoutePathInterface) {

    rp.Use(func(next HandlerFunc) HandlerFunc{
        // should be return next
    })

    rp.Group("/example", f func(rg RoutePathInterface) {
        rg.Get("/" func (c echo.Context) error {
            // should be return Echo response
        })
    })

    rp.Get("/example", func (c echo.Context) error {
        // should be return Echo response
    })

    rp.Post("/example", func (c echo.Context) error {
        // should be return Echo response
    })

    rp.Put("/example", func (c echo.Context) error {
        // should be return Echo response
    })

    rp.Patch("/example", func (c echo.Context) error {
        // should be return Echo response
    })

    rp.Delete("/example", func (c echo.Context) error {
        // should be return Echo response
    })

    rp.Any("/example", func (c echo.Context) error {
        // should be return Echo response
    })

}
```

<br>

___

<br>

## Start HTTP
[Back to top](#content)

<br>

To start HTTP you need different method, for example

```go
arah.Bind(
    arah.Start{
        Port: "8080",
    },
)
```

<br>

All start struct you can use

```go
// Start HTTP
type Start struct {
	Port string
}

// Start HTTP with TLS
type StartTLS struct {
	Port              string
	CertFile, KeyFile interface{}
}


// Start HTTP with auto creating TLS
type StartAutoTLS struct {
	Port string
}


// Start HTTP 2
type StartH2CServer struct {
	Port  string
	HTTP2 http2.Server
}


// Start HTTP with custom configuration HTTP Server
type StartServer struct {
	Server http.Server
}

// If echo has custom Listener, run this one
type StartListener struct {
}
```
For more detail check out https://echo.labstack.com/guide/http_server