package arah

var routeLists = &routeList{
	map[string]*host{},
}

// List route
type routeList struct {
	routes map[string]*host
}

func (rl *routeList) name(route string, params ...interface{}) (string, error) {
	r, isOk := rl.routes[route]
	if !isOk {
		return "", ErrorRouteNotFound
	}
	return r.echo.Reverse(route, params...), nil
}
