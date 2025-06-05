package main

import (
	"fmt"
	"sync"
)

type ChangeStrategy interface {
	GiveChange(cashier *Cashier, amount int) (map[int]int, error)
}

type Cashier struct {
	bills    map[int]int
	mu       sync.Mutex
	strategy ChangeStrategy
}

var cashierInstance *Cashier
var once sync.Once

func GetCashierInstance() *Cashier {
	once.Do(func() {
		cashierInstance = &Cashier{
			bills: make(map[int]int),
		}
	})
	return cashierInstance
}

func (c *Cashier) AcceptBill(money int, count int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bills[money] += count
}

func (c *Cashier) ShowBills() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Current bills in cashier:", c.bills)
}

func (c *Cashier) SetStrategy(strategy ChangeStrategy) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.strategy = strategy
}

func (c *Cashier) GiveChange(amount int) (map[int]int, error) {
	c.mu.Lock()
	strategy := c.strategy
	c.mu.Unlock()

	if strategy == nil {
		return nil, fmt.Errorf("no strategy set")
	}
	return strategy.GiveChange(c, amount)

}

type GreedyStrategy struct{}

func (g *GreedyStrategy) GiveChange(cashier *Cashier, amount int) (map[int]int, error) {
	cashier.mu.Lock()
	defer cashier.mu.Unlock()

	change := make(map[int]int)
	denominations := []int{100, 20, 10, 5, 1}

	for _, bill := range denominations {
		for amount >= bill && cashier.bills[bill] > 0 {
			amount -= bill
			cashier.bills[bill] -= 1
			change[bill] += 1
		}
	}

	if amount != 0 {
		for bill, count := range change {
			cashier.bills[bill] += count
		}
		return nil, fmt.Errorf("unable to give exact change")
	}
	return change, nil
}

func main() {
	cashier := GetCashierInstance()

	cashier.AcceptBill(100, 2)
	cashier.AcceptBill(20, 5)
	cashier.AcceptBill(10, 5)
	cashier.AcceptBill(5, 5)
	cashier.AcceptBill(1, 10)

	cashier.ShowBills()
	// Set the strategy
	cashier.SetStrategy(&GreedyStrategy{})

	change, err := cashier.GiveChange(136)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Change given: ", change)
	}
	cashier.ShowBills()
}
