package main

import (
	"fmt"
)

type order struct {
	ticket  int
	volume  int
	symbol  string
	account int
}

func main() {
	closed := order{
		ticket:  100000,
		volume:  100,
		symbol:  "EURUSD.ecn",
		account: 444499,
	}
	execute := IbDeposit(&closed.ticket, &closed.account, &closed.symbol, &closed.volume)
	fmt.Println(execute)
}
