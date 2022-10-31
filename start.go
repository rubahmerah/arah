package arah

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

// Start HTTP
type Start struct {
	Port string
}

func (s Start) Start(e *echo.Echo) error {
	if err := e.Start(strings.Join([]string{":", s.Port}, "")); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Start HTTP with TLS
type StartTLS struct {
	Port              string
	CertFile, KeyFile interface{}
}

func (s StartTLS) Start(e *echo.Echo) error {
	if err := e.StartTLS(strings.Join([]string{":", s.Port}, ""), s.CertFile, s.KeyFile); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Start HTTP with auto creating TLS
type StartAutoTLS struct {
	Port string
}

func (s StartAutoTLS) Start(e *echo.Echo) error {
	if err := e.StartAutoTLS(strings.Join([]string{":", s.Port}, "")); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Start HTTP 2
type StartH2CServer struct {
	Port  string
	HTTP2 http2.Server
}

func (s StartH2CServer) Start(e *echo.Echo) error {
	if err := e.StartH2CServer(strings.Join([]string{":", s.Port}, ""), &s.HTTP2); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Start HTTP with custom configuration HTTP Server
type StartServer struct {
	Server http.Server
}

func (s StartServer) Start(e *echo.Echo) error {
	if err := e.StartServer(&s.Server); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// If echo has custom Listener, run this one
type StartListener struct {
}

func (s StartListener) Start(e *echo.Echo) error {
	if err := e.Start(""); err != http.ErrServerClosed {
		return err
	}
	return nil
}
