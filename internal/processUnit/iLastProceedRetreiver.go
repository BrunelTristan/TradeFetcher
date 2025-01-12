package processUnit

type ILastProceedRetriever interface {
	GetLastProceedTimestamp() int64
}
