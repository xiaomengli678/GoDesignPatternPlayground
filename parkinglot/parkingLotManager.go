package main

import (
	"fmt"
	"sync"
)

type ParkingLotManager struct {
	totalSpots int
	occupied   int
}

func (p *ParkingLotManager) Park() bool {
	if p.occupied >= p.totalSpots {
		fmt.Println("Parking Full ❌")
		return false
	}

	p.occupied += 1
	fmt.Println("Vehicle parked. ✅")
	return true
}

func (p *ParkingLotManager) Unpark() {
	if p.occupied > 0 {
		p.occupied -= 1
		fmt.Println("Vehicle unparked. ✅")
	} else {
		fmt.Println("Parking lot already empty. ❌")
	}
}

func (p *ParkingLotManager) Status() {
	fmt.Printf("Occupied: %d / %d\n", p.occupied, p.totalSpots)
}

var (
	managerInstance *ParkingLotManager
	once            sync.Once
)

func GetParkingLotManager(totalSpots int) *ParkingLotManager {
	once.Do(func() {
		fmt.Println("Creating parking lot manager...")
		managerInstance = &ParkingLotManager{
			totalSpots: totalSpots,
		}
	})
	return managerInstance
}
