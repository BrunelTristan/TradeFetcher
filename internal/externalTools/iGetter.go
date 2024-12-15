package externalTools

type IGetter interface {
	Get(parameters interface{}) interface{}
}
