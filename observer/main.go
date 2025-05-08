package main

import "fmt"

type Observer interface {
	Update(status string)
	String() string
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

type PhoneApp struct {
	id     string
	status string
}

func (d *PhoneApp) Update(status string) {
	d.status = status
}
func (d *PhoneApp) String() string {
	return d.id + " " + d.status
}

type Tablet struct {
	id     string
	status string
}

func (t *Tablet) Update(status string) {
	t.status = status
}
func (t *Tablet) String() string {
	return t.id + " " + t.status
}

func main() {
	parkingLot := &ParkingLotWithObserver{}
	ob1 := &PhoneApp{id: "1", status: "new"}
	ob2 := &Tablet{id: "2", status: "newer"}

	parkingLot.RegisterObserver(ob1)
	parkingLot.RegisterObserver(ob2)
	fmt.Println(parkingLot.observers)
	parkingLot.SetStatus("A")
	parkingLot.RemoveObserver(ob2)
	fmt.Println(parkingLot.observers)
}
