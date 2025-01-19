package fetcher

import (
	"tradeFetcher/internal/processUnit"
	"tradeFetcher/model/trading"
)

type FilterByDateFetcherDecorator struct {
	decoratee           IFetcher
	filterDateRetriever processUnit.ILastProceedRetriever
}

func NewFilterByDateFetcherDecorator(decorateeFetcher IFetcher, lastProceedRetriever processUnit.ILastProceedRetriever) IFetcher {
	return &FilterByDateFetcherDecorator{
		decoratee:           decorateeFetcher,
		filterDateRetriever: lastProceedRetriever,
	}
}

func (f FilterByDateFetcherDecorator) FetchLastTrades() ([]*trading.Trade, error) {
	trades, err := f.decoratee.FetchLastTrades()

	if err != nil {
		return nil, err
	}

	filterTimestamp := f.filterDateRetriever.GetLastProceedTimestamp()
	filteredTrades := []*trading.Trade{}

	for _, trade := range trades {
		if trade.ExecutedTimestamp > filterTimestamp {
			filteredTrades = append(filteredTrades, trade)
		}
	}

	return filteredTrades, nil
}
