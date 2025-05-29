package main

import "fmt"

type Room interface {
	calculatePrice() int
}

type BasicRoom struct {
	price int
}

func (b *BasicRoom) calculatePrice() int {
	return b.price
}
func (b *BasicRoom) setPrice(p int) {
	b.price = p
}

type RoomCoupon struct {
	baseRoom Room
}

func (r *RoomCoupon) calculatePrice() int {
	return int(float32(r.baseRoom.calculatePrice()) * 0.8)
}

type BigRoomWithOceanView struct {
	baseRoom Room
}

func (r *BigRoomWithOceanView) calculatePrice() int {
	return r.baseRoom.calculatePrice() + 300
}

func main() {
	room := &BasicRoom{}
	room.setPrice(50)
	room0 := &BigRoomWithOceanView{baseRoom: room}
	room1 := &RoomCoupon{baseRoom: room0}
	// room2 := &BigRoomWithOceanView{baseRoom: room1}
	fmt.Println(room1.calculatePrice())
	// 300 + 50 * 0.8 = 340
	// 350 * 0.8 = 280
}
