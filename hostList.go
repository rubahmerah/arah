package route

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var hostLists = &hostList{
	map[string]*host{},
}

type hostList struct {
	hosts map[string]*host
}

func (hl *hostList) register(hostname string, config *HostConfiguration) error {
	h, _ := hl.host(hostname)
	if h != nil {
		return errors.New("host already defined")
	}
	h = &host{
		hostname: hostname,
		echo:     echo.New(),
	}
	if config != nil {
		h.config = *config
	}
	h.config.HideBanner = true
	h.prep()
	hl.hosts[hostname] = h
	return nil
}

func (hl *hostList) host(hostname string) (*host, error) {
	h, isOk := hl.hosts[hostname]
	if !isOk {
		return nil, errors.New("host not found")
	}
	return h, nil
}
