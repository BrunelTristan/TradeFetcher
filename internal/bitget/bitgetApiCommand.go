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

func NewBitgetApiCommand() externalTools.ICommand[bitgetModel.ApiCommandParameters] {
	return &BitgetApiCommand{}
}

func (c *BitgetApiCommand) Get(parameters *bitgetModel.ApiCommandParameters) (interface{}, error) {
	if parameters == nil {
		return nil, &customError.RestApiError{HttpCode: 999}
	}

	var fullUrlBuilder strings.Builder

	fullUrlBuilder.WriteString("https://api.bitget.com")
	fullUrlBuilder.WriteString(parameters.Route)

	response, err := http.Get(fullUrlBuilder.String())

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	return string(responseData), err
}
