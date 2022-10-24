package route

import (
	"errors"
	"log"
)

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

func Host(hostname string) (*host, error) {
	return hostLists.host(hostname)
}

func Name(routeName string, params ...interface{}) (string, error) {
	return routeLists.name(routeName, params...)
}

func Bind(s StartInterface) {
	host, err := hostLists.host(defaultHost)
	if err != nil {
		log.Fatal(errors.New("default host not found"))
	}
	echo := host.echo
	s.Start(echo)
}
