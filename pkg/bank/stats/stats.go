package stats

import "github.com/ainizoda/goLessons/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, payment := range payments {
		sum += payment.Amount
	}
	return types.Money(int(sum) / len(payments))
}

func TotalInCategory(payments []types.Payment, category types.PaymentCategory) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Category == category {
			total += payment.Amount
		}
	}
	return total
}
