package types

type Money int64

// type PaymentCategory string

// type PaymentStatus string

// const (
// 	StatusOk         PaymentStatus = "OK"
// 	StatusFail       PaymentStatus = "FAIL"
// 	StatusInProgress PaymentStatus = "INPROGRESS"
// )

type Payment struct {
	ID     int
	Amount Money
	// Category PaymentCategory
	// Status   PaymentStatus
}

// type Phone string

// type Account struct {
// 	ID      int
// 	Amount  Money
// 	Balance Money
// }

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID         int64
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
