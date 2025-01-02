package externalTools

type ICrypter interface {
	Crypt(message string, key string) []byte
}
