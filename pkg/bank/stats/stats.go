package stats

import "github.com/ainizoda/go-lessons-types/v2/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	return types.Money(int(sum) / len(payments))
}

func TotalInCategory(payments []types.Payment, category types.PaymentCategory) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			total += payment.Amount
		}
	}
	return total
}
