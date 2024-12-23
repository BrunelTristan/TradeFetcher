package common

type IQuery[T any] interface {
	Get(parameters *T) (interface{}, error)
}
