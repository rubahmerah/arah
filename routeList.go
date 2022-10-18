package route

import "errors"

var routeLists = &routeList{
	map[string]*host{},
}

type routeList struct {
	routes map[string]*host
}

func (rl *routeList) name(route string, params ...interface{}) (string, error) {
	r, isOk := rl.routes[route]
	if !isOk {
		return "", errors.New("Route not found")
	}
	return r.echo.Reverse(route, params...), nil
}
