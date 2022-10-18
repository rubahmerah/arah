package route

type HostHandler struct {
	Hostname string
	Config   *HostConfiguration
	Route    []RouteInterface
}
