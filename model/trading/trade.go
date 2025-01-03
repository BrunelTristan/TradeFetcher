package trading

type Trade struct {
	Pair              string
	ExecutedTimestamp int64
	Price             float64
	Quantity          float64
	Fees              float64
}
