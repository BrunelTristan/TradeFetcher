package threading

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
	"tradeFetcher/internal/generatedMocks"
)

func TestNewPeriodicThreadOrchestrator(t *testing.T) {
	fakeObject := NewPeriodicThreadOrchestrator(nil, int64(0))

	assert.NotNil(t, fakeObject)
}

func TestPeriodicThreadOrchestratorEndNotStratedOrchestration(t *testing.T) {
	mockController := gomock.NewController(t)

	workerMock := generatedMocks.NewMockIThreadWorker(mockController)

	workerMock.
		EXPECT().
		Run().
		Times(0)

	worker := NewPeriodicThreadOrchestrator(workerMock, int64(1))

	worker.EndOrchestration()
}

func TestPeriodicThreadOrchestratorRunOnceAtStartup(t *testing.T) {
	mockController := gomock.NewController(t)

	workerMock := generatedMocks.NewMockIThreadWorker(mockController)

	workerMock.
		EXPECT().
		Run().
		Times(1)

	worker := NewPeriodicThreadOrchestrator(workerMock, 999999999)

	worker.Orchestrate()
	worker.EndOrchestration()
}

func TestPeriodicThreadOrchestratorRunMultipleTimes(t *testing.T) {
	mockController := gomock.NewController(t)

	workerMock := generatedMocks.NewMockIThreadWorker(mockController)

	workerMock.
		EXPECT().
		Run().
		Times(8)

	worker := NewPeriodicThreadOrchestrator(workerMock, int64(15))

	worker.Orchestrate()

	time.Sleep(110 * time.Millisecond)

	worker.EndOrchestration()
}
