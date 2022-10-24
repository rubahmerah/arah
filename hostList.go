package route

import (
	"errors"
	"strings"

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
	hostname = strings.ToLower(hostname)
	h, _ := hl.host(hostname)
	if h != nil {
		return errors.New("host already defined: " + hostname)
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
	h, isOk := hl.hosts[strings.ToLower(hostname)]
	if !isOk {
		return nil, errors.New("host not found")
	}
	return h, nil
}
