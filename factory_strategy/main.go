package main

import "fmt"

type ParkingSpot interface {
	GetType() string
	GetPrice(hours int) int
}

type PricingStrategy interface {
	Calculate(hours int) int
}

type CompactSpot struct {
	pricingStrategy PricingStrategy
}

func (c *CompactSpot) GetType() string {
	return "Compact"
}
func (c *CompactSpot) GetPrice(hours int) int {
	return c.pricingStrategy.Calculate(hours)
}

type LargeSpot struct {
	pricingStrategy PricingStrategy
}

func (l *LargeSpot) GetType() string {
	return "Large"
}
func (l *LargeSpot) GetPrice(hours int) int {
	return l.pricingStrategy.Calculate(hours)
}

type HourlyPricingStragegy struct {
	ratePerHour int
}

func (h *HourlyPricingStragegy) Calculate(hours int) int {
	return h.ratePerHour * hours
}

type LumpSumStragegy struct {
	lumpSum int
}

func (l *LumpSumStragegy) Calculate(hours int) int {
	return l.lumpSum
}

type ParkingSpotFactory struct{}

func (f *ParkingSpotFactory) CreateParkingSpot(spotType string, stragegy PricingStrategy) ParkingSpot {
	switch spotType {
	case "Compact":
		return &CompactSpot{pricingStrategy: stragegy}
	case "Large":
		return &LargeSpot{pricingStrategy: stragegy}
	default:
		return nil
	}
}

func main() {
	factory := &ParkingSpotFactory{}

	hourlystrategy := &HourlyPricingStragegy{ratePerHour: 5}
	lumpsumstrategy := &LumpSumStragegy{lumpSum: 1000}

	compactSpot := factory.CreateParkingSpot("Compact", hourlystrategy)
	bigSpot := factory.CreateParkingSpot("Large", lumpsumstrategy)

	fmt.Println(compactSpot.GetType(), compactSpot.GetPrice(3))
	fmt.Println(bigSpot.GetType(), bigSpot.GetPrice(3))
}
