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
}

func NewBitgetApiCommand(
	accountCfg *bitgetModel.AccountConfiguration,
	signBuilder interface{},
) externalTools.ICommand[bitgetModel.ApiCommandParameters] {
	return &BitgetApiCommand{}
}

func (c *BitgetApiCommand) Get(parameters *bitgetModel.ApiCommandParameters) (interface{}, error) {
	if parameters == nil {
		return nil, &customError.RestApiError{HttpCode: 999}
	}

	var fullUrlBuilder strings.Builder

	fullUrlBuilder.WriteString("https://api.bitget.com")
	fullUrlBuilder.WriteString(parameters.Route)

	request, err := http.NewRequest("GET", fullUrlBuilder.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	return string(responseData), err
}
