package arah

import (
	"log"
)

// Register the Host Handler
func Register(hostHandler ...HostHandler) error {
	for _, hh := range hostHandler {
		hostname := hh.Hostname
		if hostname == "" {
			hostname = defaultHost
		}
		err := hostLists.register(hostname, hh.Config)
		if err != nil {
			return err
		}
		host, err := hostLists.host(hostname)
		if err != nil {
			return err
		}
		for _, r := range hh.Route {
			routeSingle := routePathSingle{
				route: &route{
					host: host,
					echo: host.echo,
				},
			}
			r.Create(routeSingle)
		}
	}
	return nil
}

// Return Host with specified name
func Host(hostname string) (*host, error) {
	return hostLists.host(hostname)
}

// Return Route Path with specified name
func Name(routeName string, params ...interface{}) (string, error) {
	return routeLists.name(routeName, params...)
}

// Bind Main Host into http and serve it
func Bind(s StartInterface) {
	host, err := hostLists.host(defaultHost)
	if err != nil {
		log.Fatal(DefaultHostNotFound{})
	}
	echo := host.echo
	s.Start(echo)
}
