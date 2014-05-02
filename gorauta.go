package gorauta

import "strings"

type RouteMap map[string][]string

type Router struct {
  // map of all routes
  routes RouteMap
}

func (r *Router) register(path string, hosts []string) bool {
  // TODO: support regexp based routing
  r.routes[path] = hosts
  return true
}

func (r *Router) HostFor(query string) string {
  // Try routing by exact match
  match := r.routes[query]
  if match != nil && len(match) > 0 {
    return match[0]
  }
  // Try to route based on prefixes
  for path, hosts := range r.routes {
    if strings.HasPrefix(query, path) && len(hosts) > 0 {
      return hosts[0]
    }
  }
  // Could not find a match
  return ""
}

func NewRouter(hosts RouteMap, routes RouteMap) *Router {
  router := new(Router)
  router.routes = make(RouteMap)
  for service, paths := range routes {
    for _, path := range paths {
      router.register(path, hosts[service])
    }
  }
  return router
}
