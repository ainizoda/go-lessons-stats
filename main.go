package main

import (
	"fmt"

	"github.com/ainizoda/goLessons/pkg/bank/card"
	"github.com/ainizoda/goLessons/pkg/bank/types"
)

func main() {
	visa := card.Issue(types.USD, "black", "ginger")
	card.Deposit(&visa, 4000)
	fmt.Println(visa.Balance)
	card.Withdraw(&visa, 1999)
	fmt.Println(visa.Balance)
}
