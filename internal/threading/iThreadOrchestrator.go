package threading

type IThreadOrchestrator interface {
	Orchestrate()
	EndOrchestration()
}
