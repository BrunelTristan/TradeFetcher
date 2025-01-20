package threading

import (
	"time"
)

type PeriodicThreadOrchestrator struct {
	worker         IThreadWorker
	periodicity    int64
	runningChannel chan struct{}
}

func NewPeriodicThreadOrchestrator(thread IThreadWorker, msPeriod int64) IThreadOrchestrator {
	return &PeriodicThreadOrchestrator{
		worker:         thread,
		periodicity:    msPeriod,
		runningChannel: make(chan struct{}),
	}
}

func (o *PeriodicThreadOrchestrator) Orchestrate() {
	o.worker.Run()

	ticker := time.NewTicker(time.Duration(o.periodicity) * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				o.worker.Run()
			case <-o.runningChannel:
				ticker.Stop()
				return
			}
		}
	}()
}

func (o *PeriodicThreadOrchestrator) EndOrchestration() {
	close(o.runningChannel)
}
