package main

import (
	"fmt"
)

func main() {
	manager1 := GetParkingLotManager(10)
	manager2 := GetParkingLotManager(1000) // ignored!

	manager1.Park()
	manager2.Park()
	manager1.Status()
	manager2.Status()

	fmt.Println("Same instance?", manager1 == manager2) // ✅ true
}
