package arah

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

/**
 * Mockup
 */

type RouteExample struct {
}

func (re RouteExample) Create(rp RoutePathInterface) {
	rp.Get("/", func(c echo.Context) error {
		return nil
	}).Name("user")
}

type StartExample struct {
	Emot string
}

func (se StartExample) Start(e *echo.Echo) error {
	fmt.Println(se.Emot)
	return nil
}

/**
 * Test Drive 😃
 */

func TestHostRegister(t *testing.T) {
	hostHandler := []HostHandler{
		{
			Hostname: "user",
			Route: []RouteInterface{
				RouteExample{},
			},
		},
		{},
	}

	err := Register(hostHandler...)
	assert.Nil(t, err)
}

func TestHostGetter(t *testing.T) {
	h, err := Host("user")
	assert.Nil(t, err)
	assert.NotNil(t, h)
}

func TestRoutePathGetter(t *testing.T) {
	routePath, err := Name("user")
	assert.Nil(t, err)
	assert.NotEmpty(t, routePath)
}

func TestBinder(t *testing.T) {
	err := Bind(StartExample{Emot: "😄"})
	assert.Nil(t, err)
}
