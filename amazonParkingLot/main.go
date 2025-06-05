package main

// available sizes
const (
	small  = "small"
	big    = "big"
	medium = "medium"
)
const (
	parking_levels         = 3
	parking_lots_per_level = 3
)

var globalSpotID int = 0

func getNextSpotID() int {
	globalSpotID++
	return globalSpotID
}

type Car interface {
	getId() int
	getSize() string
}
type ParkingSpot interface {
	getLevel() int
	getId() int
	getSize() string
	getAvailability() bool
	setAvailability(bool)
}

type homeCar struct {
	size string
	id   int
}

func (h *homeCar) getId() int {
	return h.id
}
func (h *homeCar) getSize() string {
	return h.size
}

type Truck struct {
	size string
	id   int
}

func (t *Truck) getId() int {
	return t.id
}
func (t *Truck) getSize() string {
	return t.size
}

type Motorcycle struct {
	size string
	id   int
}

func (m *Motorcycle) getId() int {
	return m.id
}
func (m *Motorcycle) getSize() string {
	return m.size
}

type GeneralSpot struct {
	level     int
	id        int
	size      string
	available bool
}

func (g *GeneralSpot) getLevel() int          { return g.level }
func (g *GeneralSpot) getId() int             { return g.id }
func (g *GeneralSpot) getSize() string        { return g.size }
func (g *GeneralSpot) getAvailability() bool  { return g.available }
func (g *GeneralSpot) setAvailability(b bool) { g.available = b }

type carRegisterStrategy interface {
	carRegister(car Car, levels map[int][]ParkingSpot, record_car_spot map[int]int) bool
}
type generalCarRegisterStrategy struct{}

func (g *generalCarRegisterStrategy) carRegister(car Car, levels map[int][]ParkingSpot, record_car_spot map[int]int) bool {
	for i := 0; i < parking_levels; i++ {
		for j := 0; j < parking_lots_per_level; j++ {
			if levels[i][j].getAvailability() && levels[i][j].getSize() == car.getSize() {
				record_car_spot[car.getId()] = levels[i][j].getId()
				levels[i][j].setAvailability(false)
				return true
			}
		}
	}
	return false
}

type rankingStrategy struct{}

func (r *rankingStrategy) carRegister(car Car, levels map[int][]ParkingSpot, record_car_spot map[int]int) bool {
	for i := 0; i < parking_levels; i++ {
		for j := 0; j < parking_lots_per_level; j++ {
			if levels[i][j].getAvailability() {
				record_car_spot[car.getId()] = levels[i][j].getId()
				levels[i][j].setAvailability(false)
				return true
			}
		}
	}
	return false
}

type ParkingLotSystem struct {
	levels          map[int][]ParkingSpot
	record_car_spot map[int]int
	strategy        carRegisterStrategy
}

func createAnActualParkingSystem() *ParkingLotSystem {
	levels := make(map[int][]ParkingSpot)
	for i := 0; i < parking_levels; i++ {
		for j := 0; j < parking_lots_per_level; j++ {
			if j < parking_lots_per_level/3 {
				g := &GeneralSpot{level: i, id: getNextSpotID(), size: small, available: true}
				levels[i] = append(levels[i], g)
			} else if j < parking_lots_per_level/3*2 {
				g := &GeneralSpot{level: i, id: getNextSpotID(), size: medium, available: true}
				levels[i] = append(levels[i], g)
			} else {
				g := &GeneralSpot{level: i, id: getNextSpotID(), size: big, available: true}
				levels[i] = append(levels[i], g)
			}
		}
	}

	p := &ParkingLotSystem{levels: levels, record_car_spot: make(map[int]int)}
	return p
}
func (p *ParkingLotSystem) carRegister(car Car) bool {
	p.strategy = &generalCarRegisterStrategy{}
	if car.getSize() == big || car.getSize() == medium {
		return p.strategy.carRegister(car, p.levels, p.record_car_spot)
	} else {
		temp := p.strategy.carRegister(car, p.levels, p.record_car_spot)
		if !temp {
			p.strategy = &rankingStrategy{}
			return p.strategy.carRegister(car, p.levels, p.record_car_spot)
		}
		return true
	}
	return false
}
func (p *ParkingLotSystem) carUnRegister(car Car) bool {
	id, ok := p.record_car_spot[car.getId()]
	if ok {
		for i := 0; i < parking_levels; i++ {
			for j := 0; j < parking_lots_per_level; j++ {
				if p.levels[i][j].getId() == id {
					p.levels[i][j].setAvailability(true)
					delete(p.record_car_spot, car.getId())

					return true
				}
			}
		}
		return false
	}
	return false
}

func createVechicle(s string, id_ int) Car {
	switch s {
	case "homeCar":
		return &homeCar{id: id_, size: medium}
	case "truck":
		return &Truck{id: id_, size: big}
	case "motorcycle":
		return &Motorcycle{id: id_, size: small}
	}
	return &homeCar{id: id_, size: medium}
}

func main() {

}
