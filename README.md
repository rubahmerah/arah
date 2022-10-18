# Route <sup style="font-size:0.4em">(Not done yet)</sup> <sub>by iamredfoxx 🦊</sub> 

<br>

Content

- Example

<br>
___

<br>

## Example

<br>

Single host
```go

package main

import (
	redfoxroute "github.com/iamaredfoxx/route"

    "net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	/**
	 * Register the Host first
	 */
	redfoxroute.Register(
		redfoxroute.HostHandler{
			Hostname: "main",
			Config:   nil,
			Route: []redfoxroute.RouteInterface{
				routes.MainRoute{},
			},
		}
	)

	/**
	 * Bind route
	 */
	redfoxroute.Bind("8080")
}

type MainRoute struct {
}

func (mr MainRoute) Create(rp redfoxroute.RoutePathInterface) {
	rp.Get("", Home)
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Home")
}

```
<br>

Multiple Host

```go

package main

import (
	redfoxroute "github.com/iamaredfoxx/route"

    "net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	/**
	 * Register the Host first
	 */
	redfoxroute.Register(
		redfoxroute.HostHandler{
			Hostname: "main",
			Config:   nil,
			Route: []redfoxroute.RouteInterface{
				routes.MainRoute{},
			},
		},
		redfoxroute.HostHandler{
			Hostname: "user",
			Config:   nil,
			Route: []redfoxroute.RouteInterface{
				routes.UserRoute{},
			},
		}
	)

	/**
	 * Bind route
	 */
	redfoxroute.Bind("8080")
}


/**
 * Main Route
 */

type MainRoute struct {
}

func (mr MainRoute) Create(rp redfoxroute.RoutePathInterface) {
	rp.Get("", Home)

    rp.Any("/user/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host, err := routeExample.Host("user")
		if err != nil {
			err = echo.ErrNotFound
		} else {
			host.ServeHTTP(res, req)
		}
		return
	})
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Home")
}

/**
 * User Route
 */

type MainRoute struct {
}

func (mr MainRoute) Create(rp redfoxroute.RoutePathInterface) {
	rp.Group("/user", func(rg routeExample.RoutePathInterface) {
		rg.Get("", User)
	})
}

func User(c echo.Context) error {
	return c.String(http.StatusOK, "User")
}


```