package bitget

type IApiResponse[T any] interface {
	GetCode() string
	GetMessage() string
	GetList() []*T
}
