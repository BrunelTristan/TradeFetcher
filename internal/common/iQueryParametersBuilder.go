package common

type IQueryParametersBuilder[T any] interface {
	BuildQueryParameters() (*T, error)
}
