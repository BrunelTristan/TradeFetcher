package bitget

import (
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/json"
)

type ApiQueryToStructDecorator[T any, J any] struct {
	decoratee         common.IQuery[T]
	parametersBuilder common.IQueryParametersBuilder[T]
	jsonConverter     json.IJsonConverter[J]
}

func NewApiQueryToStructDecorator[T any, J any](
	decor common.IQuery[T],
	paramsBuilder common.IQueryParametersBuilder[T],
	jConverter json.IJsonConverter[J],
) common.IQuery[T] {
	return &ApiQueryToStructDecorator[T, J]{
		decoratee:         decor,
		parametersBuilder: paramsBuilder,
		jsonConverter:     jConverter,
	}
}

func (c *ApiQueryToStructDecorator[T, J]) Get(parameters *T) (interface{}, error) {
	params, err := c.parametersBuilder.BuildQueryParameters()
	if err != nil {
		return nil, err
	}

	jsonGet, err := c.decoratee.Get(params)
	if err != nil {
		return nil, err
	}

	return c.jsonConverter.Import(jsonGet.(string))
}
