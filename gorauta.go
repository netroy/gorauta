package gorauta

import "strings"

type Router struct {
	// map of all routes
	Routes map[string][]string
}

func (r *Router) register(path string, hosts []string) bool {
	// TODO: support regexp based routing
	r.Routes[path] = hosts
	return true
}

func (r *Router) HostFor(query string) string {
	// Try routing by exact match
	match := r.Routes[query]
	if match != nil && len(match) > 0 {
		return random(match)
	}
	// Try to route based on prefixes
	for path, hosts := range r.Routes {
		if strings.HasPrefix(query, path) && len(hosts) > 0 {
			return random(hosts)
		}
	}
	// Could not find a match
	return ""
}

func NewRouter(hosts map[string][]string, routes map[string][]string) *Router {
	router := new(Router)
	router.Routes = make(map[string][]string)
	for service, paths := range routes {
		for _, path := range paths {
			router.register(path, hosts[service])
		}
	}
	return router
}
