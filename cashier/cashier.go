package main

import "fmt"

type Cashier struct {
	bills    map[int]int
	strategy WithdrawStrategy
}

func NewCasher(strategy WithdrawStrategy) *Cashier {
	return &Cashier{
		bills:    make(map[int]int),
		strategy: strategy,
	}
}

func (c *Cashier) InsertBill(value int) error {
	bill, err := NewBill(value)
	if err != nil {
		return err
	}
	c.bills[bill.Value]++
	return nil
}

func (c *Cashier) ViewStatus() {
	fmt.Println("Cashier Inventory: ")
	for _, denom := range []int{1, 5, 20, 100} {
		count := c.bills[denom]
		fmt.Printf("$%d: %d bills\n", denom, count)
	}
}

func (c *Cashier) Withdraw(amount int) bool {
	copyBills := make(map[int]int)
	for k, v := range c.bills {
		copyBills[k] = v
	}

	res, ok := c.strategy.Withdraw(amount, copyBills)
	if !ok {
		fmt.Printf("Cannot withdraw $%d — insufficient funds\n", amount)
		return false
	}

	// apply withdraw
	for denom, count := range res {
		c.bills[denom] -= count
	}

	fmt.Printf("Withdrawn $%d using:\n", amount)
	for denom, count := range res {
		fmt.Printf("  $%d x %d\n", denom, count)
	}
	return true
}
