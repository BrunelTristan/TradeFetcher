package externalTools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApiRouteBuilder(t *testing.T) {
	object := NewApiRouteBuilder()

	assert.NotNil(t, object)
}

func TestBuildRouteOnEmptyMaps(t *testing.T) {
	routes := make([]string, 0)
	params := make(map[string]string)

	builder := NewApiRouteBuilder()

	route := builder.BuildRoute(routes, params)

	assert.NotNil(t, route)
}

// TODO test (and implement) management of slash (needed, or in excess)

func TestBuildRouteWithoutParams(t *testing.T) {
	routes := make([]string, 3)
	params := make(map[string]string)

	routes = append(routes, "/myPath")
	routes = append(routes, "/isHere")
	routes = append(routes, "/andThere")

	builder := NewApiRouteBuilder()

	route := builder.BuildRoute(routes, params)

	assert.NotNil(t, route)
	assert.Equal(t, "/myPath/isHere/andThere", route)
}

func TestBuildRouteWithoutRoute(t *testing.T) {
	routes := make([]string, 0)
	params := make(map[string]string)

	params["sdf"] = "ghj"

	builder := NewApiRouteBuilder()

	route := builder.BuildRoute(routes, params)

	assert.NotNil(t, route)
	assert.Equal(t, "?sdf=ghj", route)
}

func TestBuildRealRoute(t *testing.T) {
	routes := make([]string, 3)
	params := make(map[string]string)

	routes = append(routes, "/myPath")
	routes = append(routes, "/isNotHere")
	routes = append(routes, "/ButThere")

	params["p1"] = "high"
	params["a25"] = "plane"
	params["what"] = "theF"

	builder := NewApiRouteBuilder()

	route := builder.BuildRoute(routes, params)

	assert.NotNil(t, route)
	assert.Equal(t, "/myPath/isNotHere/ButThere?a25=plane&p1=high&what=theF", route)
}

// TODO manage special chars
