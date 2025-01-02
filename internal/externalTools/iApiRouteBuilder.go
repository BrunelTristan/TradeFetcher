package externalTools

type IApiRouteBuilder interface {
	BuildRoute(routeComponents []string, parameters map[string]string) string
}
