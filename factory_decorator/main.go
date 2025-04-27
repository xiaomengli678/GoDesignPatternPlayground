package main

import "fmt"

type ParkingSpot interface {
	GetType() string
	GetPrice(hours int) int
}

type CompactSpot struct {
	basePrice int
}

func (c *CompactSpot) GetType() string {
	return "Compact Spot"
}
func (c *CompactSpot) GetPrice(hours int) int {
	return c.basePrice * hours
}

type LargeSpot struct {
	basePrice int
}

func (l *LargeSpot) GetType() string {
	return "Large Spot"
}
func (l *LargeSpot) GetPrice(hours int) int {
	return l.basePrice * hours
}
func (l *LargeSpot) GetRecommend() string {
	return "larggggggggggggggge"
}

type ParkingSpotFactory struct{}

func (f *ParkingSpotFactory) CreateParkingSpot(spotType string, basePrice int) ParkingSpot {
	switch spotType {
	case "Compact":
		return &CompactSpot{basePrice: basePrice}
	case "Large":
		return &LargeSpot{basePrice: basePrice}
	default:
		return nil
	}
}

type EVChargingDecorator struct {
	spot         ParkingSpot
	bighourprice int
}

func (e *EVChargingDecorator) GetType() string {
	return e.spot.GetType() + " with EV charging"
}
func (e *EVChargingDecorator) GetPrice(hours int) int {
	return e.spot.GetPrice(hours) + e.bighourprice*hours
}

type PremiumLocationDecorator struct {
	spot     ParkingSpot
	bigprice int
}

func (p *PremiumLocationDecorator) GetType() string {
	return p.spot.GetType() + " with Premium Location"
}
func (p *PremiumLocationDecorator) GetPrice(hours int) int {
	return p.spot.GetPrice(hours) + p.bigprice
}

func main() {
	factory := &ParkingSpotFactory{}

	compactSpot := factory.CreateParkingSpot("Compact", 5)
	largeSpot := factory.CreateParkingSpot("Large", 10)
	if compactSpot == nil || largeSpot == nil {
		return
	}

	compactWithTwo := &PremiumLocationDecorator{
		spot: &EVChargingDecorator{
			spot:         compactSpot,
			bighourprice: 1},
		bigprice: 10}
	largeWithPremium := &PremiumLocationDecorator{spot: largeSpot,
		bigprice: 10}

	fmt.Println(compactWithTwo.GetType(), compactWithTwo.GetPrice(1))
	fmt.Println(largeWithPremium.GetType(), largeWithPremium.GetPrice(1))

	if large, ok := largeSpot.(*LargeSpot); ok {
		fmt.Println(large.GetRecommend())
	} else {
		fmt.Println("largeSpot is not a LargeSpot")
	}
}

// 5  * 1 + 1 + 10
// 10 * 1 + 10
