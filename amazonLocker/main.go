package main

import "fmt"

var Size []string = []string{"Small", "Medium", "Large"}
var Status []string = []string{"Available", "Occupied"}

const availableLocketsPerSize = 3

type Package struct {
	id_      string
	size     string
	lockerId string
}

type Locker struct {
	id_    string
	size   string
	status string
}

type PackageSystemInterface interface {
	acceptPackage(packageSize string, packageId string) string
	pickupPackage(packageSize string, packageId string) bool
	setStrategy(ps PickupStrategy)
}
type PickupStrategy interface {
	pickLockers(lockers map[string]([]*Locker), packageSize string, packageId string) string
}
type PackageSystem struct {
	pickupStrategy          PickupStrategy
	lockers                 map[string]([]*Locker)
	recordPackageIdLockerId map[string]string
}

func createNewSystem() *PackageSystem {
	lockers := map[string]([]*Locker){} // key --> size, value --> list of lockers
	for _, s := range Size {
		lockers[s] = []*Locker{}
		for i := 0; i < availableLocketsPerSize; i++ {
			lockers[s] = append(lockers[s], &Locker{size: s, id_: string(i) + "locker", status: Status[0]})
		}
	}
	return &PackageSystem{lockers: lockers, recordPackageIdLockerId: map[string]string{}}
}
func (p *PackageSystem) setStrategy(ps PickupStrategy) {
	p.pickupStrategy = ps
}
func (p *PackageSystem) acceptPackage(packageSize string, packageId string) string {
	lockerId := p.pickupStrategy.pickLockers(p.lockers, packageSize, packageId)
	if lockerId != "" {
		p.recordPackageIdLockerId[packageId] = lockerId
	}
	return lockerId
}
func (p *PackageSystem) pickupPackage(packageSize string, packageId string) bool {
	lockerId := p.recordPackageIdLockerId[packageId]
	for _, l := range p.lockers {
		for _, locker := range l {
			if lockerId == locker.id_ {
				locker.status = Status[0]
				return true
			}
		}
	}
	return false
}

type smallStragegy struct{}

func (s *smallStragegy) pickLockers(lockers map[string]([]*Locker), packageSize string, packageId string) string {
	for _, l := range lockers["Small"] {
		if l.status == Status[0] {
			l.status = Status[1]
			return l.id_
		}
	}
	for _, l := range lockers["Medium"] {
		if l.status == Status[0] {
			l.status = Status[1]
			return l.id_
		}
	}
	for _, l := range lockers["Large"] {
		if l.status == Status[0] {
			l.status = Status[1]
			return l.id_
		}
	}
	return ""
}

type mediumStrategy struct{}

func (m *mediumStrategy) pickLockers(lockers map[string]([]*Locker), packageSize string, packageId string) string {
	for _, l := range lockers["Medium"] {
		if l.status == Status[0] {
			l.status = Status[1]
			return l.id_
		}
	}
	for _, l := range lockers["Large"] {
		if l.status == Status[0] {
			l.status = Status[1]
			return l.id_
		}
	}
	return ""
}

type LargeStrategy struct{}

func (l *LargeStrategy) pickLockers(lockers map[string]([]*Locker), packageSize string, packageId string) string {
	for _, l := range lockers["Large"] {
		if l.status == Status[0] {
			l.status = Status[1]
			return l.id_
		}
	}
	return ""
}

func pickStrategy(ps PackageSystemInterface, s string) PickupStrategy {
	s1 := &smallStragegy{}
	s2 := &mediumStrategy{}
	s3 := &LargeStrategy{}

	switch s {
	case "Small":
		return s1
	case "Medium":
		return s2
	default:
		return s3
	}

}

func main() {
	system_ := createNewSystem()
	p1 := Package{id_: "1", size: "Medium"}
	system_.setStrategy(pickStrategy(system_, p1.size))
	fmt.Println(system_.acceptPackage(p1.size, p1.id_))
}
