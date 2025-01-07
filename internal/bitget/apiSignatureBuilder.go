package bitget

import (
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
)

type ApiSignatureBuilder struct {
	accountConfiguration *bitgetModel.AccountConfiguration
	crypter              externalTools.ICrypter
	encoder              externalTools.IEncoder
}

func NewApiSignatureBuilder(
	accountCfg *bitgetModel.AccountConfiguration,
	crypt externalTools.ICrypter,
	encode externalTools.IEncoder,
) externalTools.ISignatureBuilder {
	return &ApiSignatureBuilder{
		accountConfiguration: accountCfg,
		crypter:              crypt,
		encoder:              encode,
	}
}

func (sb *ApiSignatureBuilder) Sign(message []byte) []byte {
	return sb.encoder.Encode(sb.crypter.Crypt(string(message), sb.accountConfiguration.SecretKey))
}
