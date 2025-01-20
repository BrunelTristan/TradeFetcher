package threading

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/internal/processUnit"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewFetcherProcessorsWorker(t *testing.T) {
	fakeObject := NewFetcherProcessorsWorker(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestFetcherProcessorsWorkerRunWithFetcherError(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)
	processUnitMock := generatedMocks.NewMockIProcessUnit(mockController)

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 404})

	processUnitMock.
		EXPECT().
		ProcessTrades(gomock.Any()).
		Times(0)

	worker := NewFetcherProcessorsWorker(fetcherMock, []processUnit.IProcessUnit{processUnitMock})

	worker.Run()
}

func TestFetcherProcessorsWorkerRunCleanly(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)
	processUnitMock := generatedMocks.NewMockIProcessUnit(mockController)

	trades := []*trading.Trade{
		&trading.Trade{ExecutedTimestamp: 123456},
		&trading.Trade{ExecutedTimestamp: 123457},
		&trading.Trade{ExecutedTimestamp: 123458},
	}

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(trades, nil)

	processUnitMock.
		EXPECT().
		ProcessTrades(gomock.Eq(trades)).
		Times(1)

	worker := NewFetcherProcessorsWorker(fetcherMock, []processUnit.IProcessUnit{processUnitMock})

	worker.Run()
}
