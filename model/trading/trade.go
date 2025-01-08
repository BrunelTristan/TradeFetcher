package trading

type TransactionType int

const (
	OPENING TransactionType = iota
	CLOSE
	FUNDING
)

type Trade struct {
	Pair              string
	ExecutedTimestamp int64
	TransactionType   TransactionType
	// TODO remove Open
	//Open     bool
	Long     bool
	Price    float64
	Quantity float64
	Fees     float64
}
