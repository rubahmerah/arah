package arah

import (
	"github.com/labstack/echo/v4"
)

var hostLists = &hostList{
	map[string]*host{},
}

// List host
type hostList struct {
	hosts map[string]*host
}

// Register host
func (hl *hostList) register(hostname string, config *HostConfiguration) error {
	h, _ := hl.host(hostname)
	if h != nil {
		return HostAlreadyDefined{hostname}
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

// Get host by specified name
func (hl *hostList) host(hostname string) (*host, error) {
	h, isOk := hl.hosts[hostname]
	if !isOk {
		return nil, HostAlreadyDefined{hostname}
	}
	return h, nil
}
