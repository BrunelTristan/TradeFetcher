package converter

type IStructConverter[T any, U any] interface {
	Convert(parameters *T) (*U, error)
}
