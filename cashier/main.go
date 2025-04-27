package main

import "fmt"

func main() {
	cashier1 := NewCasher(&BigToSmall{})
	for _, v := range []int{100, 100, 20, 20, 5, 5, 1, 1, 1} {
		cashier1.InsertBill(v)
	}

	cashier1.ViewStatus()
	cashier1.Withdraw(126)
	cashier1.ViewStatus()
	fmt.Println("Cashier Inventory-----------: ")
	cashier2 := NewCasher(&SmallToBig{})
	for _, v := range []int{100, 100, 20, 20, 5, 5, 1, 1, 1} {
		cashier2.InsertBill(v)
	}

	cashier2.ViewStatus()
	cashier2.Withdraw(126)
	cashier2.ViewStatus()
}
