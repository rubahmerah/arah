package route

import (
	"log"
)

func Register(routeHandler ...HostHandler) error {
	for _, rh := range routeHandler {
		err := hostLists.register(rh.Hostname, rh.Config)
		if err != nil {
			return err
		}
		host, err := hostLists.host(rh.Hostname)
		if err != nil {
			return err
		}
		for _, r := range rh.Route {
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
	host, err := hostLists.host(mainHost)
	if err != nil {
		log.Fatal(err)
	}
	echo := host.echo
	s.Start(echo)
}
