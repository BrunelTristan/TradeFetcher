package json

type IJsonConverter[T any] interface {
	Import(jsonText string) (*T, error)
	Export(object *T) (string, error)
}
