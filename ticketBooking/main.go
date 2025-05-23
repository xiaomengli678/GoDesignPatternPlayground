package main

import "fmt"

type ticket struct {
	id   int
	name string
	age  int
}
type BookingStrategy interface {
	strategyBooking(t *ticket, queue []*ticket) []*ticket
}
type generalBookingStrategy struct{}

func (s *generalBookingStrategy) strategyBooking(t *ticket, queue []*ticket) []*ticket {
	return append(queue, t)
}

type goodBookingStrategy struct{}

func (g *goodBookingStrategy) strategyBooking(t *ticket, queue []*ticket) []*ticket {
	return append([]*ticket{t}, queue...)
}

type PlaneTicketSystem struct {
	strategy BookingStrategy
	queue    []*ticket
}

func (p *PlaneTicketSystem) book(t *ticket) {
	p.queue = p.strategy.strategyBooking(t, p.queue)
}
func (p *PlaneTicketSystem) showOnboardingStatus() {
	if len(p.queue) == 0 {
		fmt.Println("-----no tickets------")
		return
	}
	fmt.Println("-----show queue status------")
	for _, p := range p.queue {
		fmt.Println("name: " + p.name)
	}
}
func (p *PlaneTicketSystem) setStrategy(s BookingStrategy) {
	p.strategy = s
}

func main() {
	ps := PlaneTicketSystem{}
	t1 := &ticket{id: 1, name: "Lucy", age: 18}
	t2 := &ticket{id: 2, name: "Tom", age: 75}
	t3 := &ticket{id: 3, name: "Jake", age: 20}
	t4 := &ticket{id: 4, name: "Lilly", age: 9}
	tList := []*ticket{t1, t2, t3, t4}
	for _, t := range tList {
		if t.age < 10 || t.age > 60 {
			ps.setStrategy(&goodBookingStrategy{})
		} else {
			ps.setStrategy(&generalBookingStrategy{})
		}
		ps.book(t)
	}

	ps.showOnboardingStatus()

}
