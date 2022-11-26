package arah

import (
	"errors"
)

var (
	ErrorDefaultHostNotFound = errors.New("default host not found")
	ErrorHostNotFound        = errors.New("host not found")
	ErrorHostAlreadyDefined  = errors.New("host already defined")
	ErrorRouteNotFound       = errors.New("route not found")
)

func IsDefaultHostNotFound(err error) bool {
	return err == ErrorDefaultHostNotFound
}

func IsHostNotFound(err error) bool {
	return err == ErrorHostNotFound
}

func IsHostAlreadyDefined(err error) bool {
	return err == ErrorHostAlreadyDefined
}

func IsRouteNotFound(err error) bool {
	return err == ErrorRouteNotFound
}
