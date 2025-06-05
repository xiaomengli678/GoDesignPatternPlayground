package main

import "fmt"

const (
	general = "general"
	luxury  = "luxury"
)
const (
	gymClass = "gymClass"
	bigTv    = "bigTV"
)

type Room struct {
	RoomType          string
	oceanView         bool
	extraRoomServices []string
}

type RoomBuilder interface {
	SetRoomType(roomType string) RoomBuilder
	SetOceanView(choice bool) RoomBuilder
	AddExtraService(service string) RoomBuilder
	Build() *Room
}

type ConcreteRoomBuilder struct {
	roomType          string
	oceanView         bool
	extraRoomServices []string
}

func createConcreteRoom() RoomBuilder {
	return &ConcreteRoomBuilder{roomType: general, oceanView: false, extraRoomServices: []string{}}
}
func (c *ConcreteRoomBuilder) SetRoomType(roomType string) RoomBuilder {
	c.roomType = roomType
	return c
}
func (c *ConcreteRoomBuilder) SetOceanView(choice bool) RoomBuilder {
	c.oceanView = choice
	return c
}
func (c *ConcreteRoomBuilder) AddExtraService(service string) RoomBuilder {
	c.extraRoomServices = append(c.extraRoomServices, service)
	return c
}
func (c *ConcreteRoomBuilder) Build() *Room {
	return &Room{RoomType: c.roomType, oceanView: c.oceanView, extraRoomServices: c.extraRoomServices}
}

func selectGoodRoom(roomType string) RoomBuilder {
	switch roomType {
	case general:
		room := createConcreteRoom().
			SetRoomType(general)
		return room
	case luxury:
		room := createConcreteRoom().
			SetRoomType(luxury)
		return room
	}
	return nil
}

func main() {
	room := selectGoodRoom(luxury).
		SetOceanView(true).
		AddExtraService(gymClass).
		Build()
	fmt.Printf("%+v\n", room)
}
