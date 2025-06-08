package main

import (
	"fmt"
	"sync"
)

const (
	small  = "small"
	medium = "medium"
	big    = "big"
)

const (
	parking_levels         = 1
	parking_lots_per_level = 1
)

var globalSpotID int
var globalSpotIDMutex sync.Mutex

func getNextSpotID() int {
	globalSpotIDMutex.Lock()
	defer globalSpotIDMutex.Unlock()
	globalSpotID++
	return globalSpotID
}

// ---- Interfaces ----
type Car interface {
	getId() int
	getSize() string
}

type ParkingSpot interface {
	getLevel() int
	getId() int
	getSize() string
	tryOccupy(size string) bool
	setAvailability(bool)
}

// ---- Vehicles ----
type HomeCar struct {
	size string
	id   int
}

func (h *HomeCar) getId() int      { return h.id }
func (h *HomeCar) getSize() string { return h.size }

type Truck struct {
	size string
	id   int
}

func (t *Truck) getId() int      { return t.id }
func (t *Truck) getSize() string { return t.size }

type Motorcycle struct {
	size string
	id   int
}

func (m *Motorcycle) getId() int      { return m.id }
func (m *Motorcycle) getSize() string { return m.size }

// ---- Parking Spot ----
type GeneralSpot struct {
	level     int
	id        int
	size      string
	available bool
	mu        sync.Mutex // protect 'available'
}

func (g *GeneralSpot) getLevel() int   { return g.level }
func (g *GeneralSpot) getId() int      { return g.id }
func (g *GeneralSpot) getSize() string { return g.size }

// This is the KEY new function
func (g *GeneralSpot) tryOccupy(carSize string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.available && g.size == carSize {
		g.available = false
		return true
	}
	return false
}

func (g *GeneralSpot) setAvailability(b bool) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.available = b
}

// ---- Strategy Pattern ----
type CarRegisterStrategy interface {
	carRegister(car Car, levels map[int][]ParkingSpot, recordCarSpot map[int]int, mu *sync.Mutex) bool
}

type GeneralCarRegisterStrategy struct{}

func (g *GeneralCarRegisterStrategy) carRegister(car Car, levels map[int][]ParkingSpot, recordCarSpot map[int]int, mu *sync.Mutex) bool {
	for _, spots := range levels {
		for _, spot := range spots {
			if spot.tryOccupy(car.getSize()) {
				mu.Lock()
				recordCarSpot[car.getId()] = spot.getId()
				mu.Unlock()
				return true
			}
		}
	}
	return false
}

type RankingStrategy struct{}

func (r *RankingStrategy) carRegister(car Car, levels map[int][]ParkingSpot, recordCarSpot map[int]int, mu *sync.Mutex) bool {
	for _, spots := range levels {
		for _, spot := range spots {
			if spot.tryOccupy(spot.getSize()) { // Don't care about car size
				mu.Lock()
				recordCarSpot[car.getId()] = spot.getId()
				mu.Unlock()
				return true
			}
		}
	}
	return false
}

// ---- Parking Lot System ----
type ParkingLotSystem struct {
	levels        map[int][]ParkingSpot
	recordCarSpot map[int]int
	mu            sync.Mutex // Protect recordCarSpot
}

func createParkingLotSystem() *ParkingLotSystem {
	levels := make(map[int][]ParkingSpot)
	for i := 0; i < parking_levels; i++ {
		for j := 0; j < parking_lots_per_level; j++ {
			size := medium
			if j < parking_lots_per_level/3 {
				size = small
			} else if j < parking_lots_per_level/3*2 {
				size = medium
			} else {
				size = big
			}
			g := &GeneralSpot{
				level:     i,
				id:        getNextSpotID(),
				size:      size,
				available: true,
			}
			levels[i] = append(levels[i], g)
		}
	}
	return &ParkingLotSystem{levels: levels, recordCarSpot: make(map[int]int)}
}

func (p *ParkingLotSystem) carRegister(car Car) bool {
	strategy := &GeneralCarRegisterStrategy{}
	if car.getSize() == big || car.getSize() == medium {
		return strategy.carRegister(car, p.levels, p.recordCarSpot, &p.mu)
	} else {
		temp := strategy.carRegister(car, p.levels, p.recordCarSpot, &p.mu)
		if !temp {
			alternative := &RankingStrategy{}
			return alternative.carRegister(car, p.levels, p.recordCarSpot, &p.mu)
		}
		return temp
	}
}

func (p *ParkingLotSystem) carUnRegister(car Car) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	id, ok := p.recordCarSpot[car.getId()]
	if !ok {
		return false
	}

	for _, spots := range p.levels {
		for _, spot := range spots {
			if spot.getId() == id {
				spot.setAvailability(true)
				delete(p.recordCarSpot, car.getId())
				return true
			}
		}
	}
	return false
}

// ---- Factory ----
func createVehicle(vehicleType string, id int) Car {
	switch vehicleType {
	case "homeCar":
		return &HomeCar{id: id, size: medium}
	case "truck":
		return &Truck{id: id, size: big}
	case "motorcycle":
		return &Motorcycle{id: id, size: small}
	default:
		return &HomeCar{id: id, size: medium}
	}
}

// ---- Main ----
func main() {
	parkingLot := createParkingLotSystem()

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			car := createVehicle("motorcycle", id)
			if parkingLot.carRegister(car) {
				fmt.Printf("Car %d parked successfully.\n", id)
			} else {
				fmt.Printf("Car %d could not find a spot.\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All cars tried to park.")
}
