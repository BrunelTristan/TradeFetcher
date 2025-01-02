package externalTools

type IEncoder interface {
	Encode(message []byte) []byte
}
