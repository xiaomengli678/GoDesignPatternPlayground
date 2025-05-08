package main

import "fmt"

type Observer interface {
	Update(status string)
}

type ParkingLotWithObserver struct {
	observers []Observer
	status    string
}

func (p *ParkingLotWithObserver) RegisterObserver(observer Observer) {
	p.observers = append(p.observers, observer)
}

func (p *ParkingLotWithObserver) RemoveObserver(observer Observer) {
	for i, o := range p.observers {
		if o == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *ParkingLotWithObserver) NotifyObservers() {
	for _, o := range p.observers {
		o.Update(p.status)
	}
}

func (p *ParkingLotWithObserver) SetStatus(status string) {
	p.status = status
	p.NotifyObservers()
}

type DisplayBoard struct {
	id     string
	status string
}

func (d *DisplayBoard) Update(status string) {
	d.status = status
	fmt.Printf("DisplayBoard %s status updated to '%s'\n", d.id, d.status)
}

type PhoneApp struct {
	id     string
	status string
}

func (d *PhoneApp) Update(status string) {
	d.status = status
	fmt.Printf("PhoneApp %s status updated to '%s'\n", d.id, d.status)
}

func main() {
	parkingLot := &ParkingLotWithObserver{}
	ob1 := &DisplayBoard{id: "1"}
	ob2 := &PhoneApp{id: "2"}

	parkingLot.RegisterObserver(ob1)
	parkingLot.RegisterObserver(ob2)

	parkingLot.SetStatus("status1")

	parkingLot.SetStatus("status____________")
	fmt.Println(len(parkingLot.observers))
	parkingLot.RemoveObserver(ob1)
	fmt.Println(len(parkingLot.observers))

}
