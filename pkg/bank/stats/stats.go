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

func FilterByCategory(payments []types.Payment, category types.PaymentCategory) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

func CategoriesTotal(payments []types.Payment) map[types.PaymentCategory]types.Money {
	categories := map[types.PaymentCategory]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}

func CategoriesAvg(payments []types.Payment) map[types.PaymentCategory]types.Money {
	categories := map[types.PaymentCategory]types.Money{}
	count := map[types.PaymentCategory]int{}

	for _, payment := range payments {
		count[payment.Category]++
	}
	for _, payment := range payments {
		categories[payment.Category] += types.Money(int(payment.Amount) / count[payment.Category])
	}
	return categories
}
