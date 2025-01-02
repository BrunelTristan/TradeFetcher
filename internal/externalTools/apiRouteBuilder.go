package externalTools

import (
	"sort"
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

	sortedKeys := make([]string, len(parameters))
	i := 0
	for key, _ := range parameters {
		sortedKeys[i] = key
		i++
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		if firstParam {
			builder.WriteString("?")
			firstParam = false
		} else {
			builder.WriteString("&")
		}
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(parameters[key])
	}

	return builder.String()
}
