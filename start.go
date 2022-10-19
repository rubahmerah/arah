package route

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

type Start struct {
	Port string
}

func (s Start) Start(e *echo.Echo) {
	if err := e.Start(strings.Join([]string{":", s.Port}, "")); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type StartTLS struct {
	Port              string
	CertFile, KeyFile interface{}
}

func (s StartTLS) Start(e *echo.Echo) {
	if err := e.StartTLS(strings.Join([]string{":", s.Port}, ""), s.CertFile, s.KeyFile); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type StartAutoTLS struct {
	Port string
}

func (s StartAutoTLS) Start(e *echo.Echo) {
	if err := e.StartAutoTLS(strings.Join([]string{":", s.Port}, "")); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type StartH2CServer struct {
	Port  string
	HTTP2 http2.Server
}

func (s StartH2CServer) Start(e *echo.Echo) {
	if err := e.StartH2CServer(strings.Join([]string{":", s.Port}, ""), &s.HTTP2); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type StartServer struct {
	Server http.Server
}

func (s StartServer) Start(e *echo.Echo) {
	if err := e.StartServer(&s.Server); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
