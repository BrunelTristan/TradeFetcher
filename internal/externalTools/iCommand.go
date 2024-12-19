package externalTools

type ICommand[T any] interface {
	Get(parameters *T) (interface{}, error)
}
