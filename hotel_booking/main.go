package main

import "fmt"

type Room struct {
	id        int
	type_     string
	price     int
	available bool
}

const (
	luxury  = "luxury"
	general = "general"
)

type SearchStrategy interface {
	searchRooms(record map[int]*Room) map[int]*Room
}

type searchForType struct {
	type_ string
}

func (s *searchForType) searchRooms(record map[int]*Room) map[int]*Room {
	ans := make(map[int]*Room)
	for id_, r := range record {
		if r.type_ == s.type_ {
			ans[id_] = r
		}
	}
	return ans
}

type searchForPrice struct {
	price int
}

func (s *searchForPrice) searchRooms(record map[int]*Room) map[int]*Room {
	ans := make(map[int]*Room)
	for id_, r := range record {
		if r.price <= s.price {
			ans[id_] = r
		}
	}
	return ans
}

type searchForBoth struct {
	strategies []SearchStrategy
}

func (s *searchForBoth) searchRooms(record map[int]*Room) map[int]*Room {
	ans := make(map[int]*Room)
	initial := true
	for _, strate := range s.strategies {
		temp := strate.searchRooms(record)
		if initial {
			initial = false
			for k, v := range temp {
				ans[k] = v
			}
		} else {
			for k, _ := range ans {
				if _, ok := temp[k]; !ok {
					delete(ans, k)
				}
			}
		}
	}
	return ans
}

type hotelSystem struct {
	searchStrategy SearchStrategy
	allTheRooms    map[int]*Room
}

func (h *hotelSystem) searchForRooms() map[int]*Room {
	return h.searchStrategy.searchRooms(h.allTheRooms)
}
func (h *hotelSystem) setStrategy(price int, type_ string) {
	s1 := searchForType{type_: type_}
	s2 := searchForPrice{price: price}
	s3 := searchForBoth{strategies: []SearchStrategy{&s1, &s2}}
	h.searchStrategy = &s3

}
func createHotel() *hotelSystem {
	allTheRooms := map[int]*Room{}
	for i := 0; i < 3; i++ {
		allTheRooms[i] = &Room{id: i, type_: general, price: 100, available: true}
	}
	for i := 3; i < 6; i++ {
		allTheRooms[i] = &Room{id: i, type_: luxury, price: 200, available: true}
	}
	return &hotelSystem{allTheRooms: allTheRooms}
}

func main() {
	ht := createHotel()
	ht.setStrategy(250, luxury)
	for _, r := range ht.searchForRooms() {
		fmt.Println(r.id, r.price, r.type_)
	}
	fmt.Println("----------")
	ht.setStrategy(100, general)
	for _, r := range ht.searchForRooms() {
		fmt.Println(r.id, r.price, r.type_)
	}
}
