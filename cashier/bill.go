package main

import "fmt"

type Bill struct {
	Value int
}

func (b *Bill) String() string {
	return fmt.Sprintf("$%d bill", b.Value)
}

func NewBill(value int) (*Bill, error) {
	switch value {
	case 1, 5, 20, 100:
		return &Bill{Value: value}, nil
	default:
		return nil, fmt.Errorf("invalid bill denomination %d", value)
	}

}
