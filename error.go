package arah

import "fmt"

type DefaultHostNotFound struct {
}

func (DefaultHostNotFound) Error() string {
	return "Default Host not found"
}

type HostNotFound struct {
	Hostname string
}

func (h HostNotFound) Error() string {
	return fmt.Sprintf("Host already defined: %s", h.Hostname)
}

type HostAlreadyDefined struct {
	Hostname string
}

func (d HostAlreadyDefined) Error() string {
	return fmt.Sprintf("Host already defined: %s", d.Hostname)
}
