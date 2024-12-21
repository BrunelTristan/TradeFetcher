package externalTools

type ISignatureBuilder interface {
	Sign(message []byte) []byte
}
