package main

import "fmt"

type ParkingSpot interface {
	GetDescription() string
	GetPrice(hours int) int
}

const (
	EVChargingHourlyRate   = 10
	PremiumLocationFlatFee = 1000
)

type PricingStrategy interface {
	CalculatePrice(hours int) int
}

type HourlyPricingStrategy struct {
	rate int
}

func (h *HourlyPricingStrategy) CalculatePrice(hours int) int {
	return h.rate * hours
}

type LumpSumPricingStragegy struct {
	lumpSum int
}

func (l *LumpSumPricingStragegy) CalculatePrice(hours int) int {
	return l.lumpSum
}

type BasicSpot struct {
	description     string
	pricingStrategy PricingStrategy
}

func (b *BasicSpot) GetDescription() string {
	return b.description
}
func (b *BasicSpot) GetPrice(hours int) int {
	return b.pricingStrategy.CalculatePrice(hours)
}

type EVChargingDecorator struct {
	spot ParkingSpot
}

func (e *EVChargingDecorator) GetDescription() string {
	return e.spot.GetDescription() + ", EV Charging"
}
func (e *EVChargingDecorator) GetPrice(hours int) int {
	return e.spot.GetPrice(hours) + EVChargingHourlyRate*hours
}

type PremiumLocationDecorator struct {
	spot ParkingSpot
}

func (p *PremiumLocationDecorator) GetDescription() string {
	return p.spot.GetDescription() + ", Premium Location"
}
func (p *PremiumLocationDecorator) GetPrice(hours int) int {
	return p.spot.GetPrice(hours) + PremiumLocationFlatFee
}

func main() {
	hourlyPricingStrategy := &HourlyPricingStrategy{rate: 50}
	lumpSumStrategy := &LumpSumPricingStragegy{lumpSum: 300}

	hourlySpot := &BasicSpot{
		description:     "Hourly Spot",
		pricingStrategy: hourlyPricingStrategy,
	}

	lumpSumSpot := &BasicSpot{
		description:     "lumpsum Spot",
		pricingStrategy: lumpSumStrategy,
	}

	evHourlySpot := &EVChargingDecorator{spot: hourlySpot}
	premiumLumpSumSpot := &PremiumLocationDecorator{spot: lumpSumSpot}

	fmt.Println(evHourlySpot.GetDescription(), evHourlySpot.GetPrice(3))
	fmt.Println(premiumLumpSumSpot.GetDescription(), premiumLumpSumSpot.GetPrice(3))
}

// 50 * 3 + 10 * 3 = 180
// 300 + 1000 = 1300"
