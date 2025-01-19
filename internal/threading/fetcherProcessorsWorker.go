package threading

import (
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/processUnit"
)

type FetcherProcessorsWorker struct {
	fetcher    fetcher.IFetcher
	processors []processUnit.IProcessUnit
}

func NewFetcherProcessorsWorker(fetch fetcher.IFetcher, processUnits []processUnit.IProcessUnit) IThreadWorker {
	return &FetcherProcessorsWorker{
		fetcher:    fetch,
		processors: processUnits,
	}
}

func (w *FetcherProcessorsWorker) Run() {
	trades, err := w.fetcher.FetchLastTrades()

	if err == nil {
		for _, processor := range w.processors {
			_ = processor.ProcessTrades(trades)
		}
	}
}
