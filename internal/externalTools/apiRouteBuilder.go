package externalTools

import (
	"strings"
)

type ApiRouteBuilder struct {
}

func NewApiRouteBuilder() IApiRouteBuilder {
	return &ApiRouteBuilder{}
}

func (b *ApiRouteBuilder) BuildRoute(
	routeComponents []string,
	parameters map[string]string,
) string {
	var builder strings.Builder
	firstParam := true

	for _, route := range routeComponents {
		builder.WriteString(route)
	}

	for key, value := range parameters {
		if firstParam {
			builder.WriteString("?")
			firstParam = false
		} else {
			builder.WriteString("&")
		}
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(value)
	}

	return builder.String()
}
