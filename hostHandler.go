package arah

// The host handler
type HostHandler struct {
	Hostname string
	Config   *HostConfiguration
	Route    []RouteInterface
}
