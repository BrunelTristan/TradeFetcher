package common

type NilQueryParametersBuilder[T any] struct {
}

func NewNilQueryParametersBuilder[T any]() IQueryParametersBuilder[T] {
	return &NilQueryParametersBuilder[T]{}
}

func (b *NilQueryParametersBuilder[T]) BuildQueryParameters() (*T, error) {
	return nil, nil
}
