package main

import "fmt"

type ParkingLot struct {
	id   int
	size string
}

type Observer interface {
	Notify(status string, p *ParkingLot)
}

type BigParkingSystem struct {
	findSpotStrategy map[string]findSpotStrategy
	available_record map[int]*ParkingLot
	occupied_record  map[int]*ParkingLot
	observers        []Observer
}

func (bps *BigParkingSystem) AddObserver(o Observer) {
	bps.observers = append(bps.observers, o)
}
func (bps *BigParkingSystem) AddParkingLotForActiveUse(p *ParkingLot) bool {
	if _, ok := bps.occupied_record[p.id]; !ok {
		bps.occupied_record[p.id] = p
		delete(bps.available_record, p.id)
		bps.NotifyDevices("occupied", p)
		return true
	}
	return false
}
func (bps *BigParkingSystem) ReleaseParkingLotForAvailableAgain(parkingId int) bool {
	if val, ok := bps.occupied_record[parkingId]; ok {
		bps.available_record[parkingId] = val
		delete(bps.occupied_record, parkingId)
		bps.NotifyDevices("available", val)
		return true
	}
	return false
}
func (bps *BigParkingSystem) NotifyDevices(status string, p *ParkingLot) {
	for _, ob := range bps.observers {
		ob.Notify(status, p)
	}
}

type Phone struct{}

func (p *Phone) Notify(status string, pl *ParkingLot) {
	fmt.Printf("phone app is updating with status: %s, parking lot number %d, parking lot size %s \n", status, pl.id, pl.size)
}

type findSpotStrategy interface {
	findMeSpot(size string, available_record map[int]*ParkingLot) (*ParkingLot, bool)
}
type findSpotNormalStrategy struct{}

func (fs *findSpotNormalStrategy) findMeSpot(size string, available_record map[int]*ParkingLot) (*ParkingLot, bool) {
	for _, v := range available_record {
		if v.size == size {
			return v, true
		}
	}
	return nil, false
}

type findSpotLargeOrEqualNeededSizeStrategy struct{}

func (fs *findSpotLargeOrEqualNeededSizeStrategy) findMeSpot(size string, available_record map[int]*ParkingLot) (*ParkingLot, bool) {
	for _, v := range available_record {
		if v.size == "medium" && size == "small" {
			return v, true
		} else if v.size == "large" && size == "medium" {
			return v, true
		}
	}
	return nil, false
}
func main() {
	phone := &Phone{}
	bps := &BigParkingSystem{
		findSpotStrategy: map[string]findSpotStrategy{"normal": &findSpotNormalStrategy{}, "LargeOrEqual": &findSpotLargeOrEqualNeededSizeStrategy{}},
		available_record: make(map[int]*ParkingLot),
		occupied_record:  make(map[int]*ParkingLot),
	}
	bps.AddObserver(phone)
	bps.AddParkingLotForActiveUse(&ParkingLot{id: 3, size: "medium"})
	bps.AddParkingLotForActiveUse(&ParkingLot{id: 4, size: "large"})
	bps.ReleaseParkingLotForAvailableAgain(4)
}
