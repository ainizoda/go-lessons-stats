package transfer

import "bank/pkg/bank/types"

const bonusPercent = 0.0050

func Bonus(amount types.Money) types.Money {
	bonus := types.Money(float64(amount) * bonusPercent)
	return bonus
}

func Total(amount types.Money) types.Money {
	total := amount + Bonus(amount)
	return total
}
