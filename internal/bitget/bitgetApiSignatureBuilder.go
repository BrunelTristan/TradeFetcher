package bitget

import (
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
)

type BitgetApiSignatureBuilder struct {
	accountConfiguration *bitgetModel.AccountConfiguration
	crypter              externalTools.ICrypter
	encoder              externalTools.IEncoder
}

func NewBitgetApiSignatureBuilder(
	accountCfg *bitgetModel.AccountConfiguration,
	crypt externalTools.ICrypter,
	encode externalTools.IEncoder,
) externalTools.ISignatureBuilder {
	return &BitgetApiSignatureBuilder{
		accountConfiguration: accountCfg,
		crypter:              crypt,
		encoder:              encode,
	}
}

func (sb *BitgetApiSignatureBuilder) Sign(message []byte) []byte {
	return sb.encoder.Encode(sb.crypter.Crypt(string(message), sb.accountConfiguration.SecretKey))
}
