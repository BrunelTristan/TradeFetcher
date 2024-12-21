package bitget

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
)

type BitgetApiCommand struct {
	accountConfiguration *bitgetModel.AccountConfiguration
	signatureBuilder     externalTools.ISignatureBuilder
}

func NewBitgetApiCommand(
	accountCfg *bitgetModel.AccountConfiguration,
	signBuilder externalTools.ISignatureBuilder,
) externalTools.ICommand[bitgetModel.ApiCommandParameters] {
	return &BitgetApiCommand{
		accountConfiguration: accountCfg,
		signatureBuilder:     signBuilder,
	}
}

func (c *BitgetApiCommand) Get(parameters *bitgetModel.ApiCommandParameters) (interface{}, error) {
	if parameters == nil {
		return nil, &customError.RestApiError{HttpCode: 999}
	}

	const GET_VERB = "GET"

	var fullUrlBuilder strings.Builder
	var fullMessageToSignBuilder strings.Builder

	fullUrlBuilder.WriteString("https://api.bitget.com")
	fullUrlBuilder.WriteString(parameters.Route)

	request, err := http.NewRequest(GET_VERB, fullUrlBuilder.String(), nil)
	if err != nil {
		return nil, err
	}

	timestamp := strconv.FormatInt(int64(time.Now().UnixNano()/1000000), 10)

	fullMessageToSignBuilder.WriteString(timestamp)
	fullMessageToSignBuilder.WriteString(GET_VERB)
	fullMessageToSignBuilder.WriteString(parameters.Route)

	signature := c.signatureBuilder.Sign([]byte(fullMessageToSignBuilder.String()))

	request.Header.Set("ACCESS-KEY", c.accountConfiguration.ApiKey)
	request.Header.Set("ACCESS-SIGN", string(signature))
	request.Header.Set("ACCESS-TIMESTAMP", timestamp)
	request.Header.Set("ACCESS-PASSPHRASE", c.accountConfiguration.PassPhrase)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	return string(responseData), err
}
