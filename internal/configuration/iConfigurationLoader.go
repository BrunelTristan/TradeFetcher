package configuration

type IConfigurationLoader[T any] interface {
	Load() (*T, error)
}
