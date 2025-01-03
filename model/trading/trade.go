package trading

type Trade struct {
	Pair              string
	ExecutedTimestamp int64
	Open              bool
	Price             float64
	Quantity          float64
	Fees              float64
}
