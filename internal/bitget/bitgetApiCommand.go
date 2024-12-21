package bitget

import (
	"io"
	"net/http"
	"strings"
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
)

type BitgetApiCommand struct {
	signatureBuilder externalTools.ISignatureBuilder
}

func NewBitgetApiCommand(
	accountCfg *bitgetModel.AccountConfiguration,
	signBuilder externalTools.ISignatureBuilder,
) externalTools.ICommand[bitgetModel.ApiCommandParameters] {
	return &BitgetApiCommand{
		signatureBuilder: signBuilder,
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

	fullMessageToSignBuilder.WriteString(GET_VERB)
	fullMessageToSignBuilder.WriteString(parameters.Route)

	request, err := http.NewRequest(GET_VERB, fullUrlBuilder.String(), nil)
	if err != nil {
		return nil, err
	}

	signature := c.signatureBuilder.Sign([]byte(fullMessageToSignBuilder.String()))

	request.Header.Set("ACCESS-SIGN", string(signature))

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	return string(responseData), err
}
