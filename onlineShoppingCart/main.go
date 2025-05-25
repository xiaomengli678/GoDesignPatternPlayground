package main

import "fmt"

type product struct {
	name  string
	stock int
	price int
}

const (
	discount = "discount"
	general  = "general"
)

type checkingoutStrategy interface {
	calculateMoney(cart map[*product]int) (float32, bool)
}

type generalCheckingout struct {
}

func (g *generalCheckingout) calculateMoney(cart map[*product]int) (float32, bool) {
	ans := 0
	for p, c := range cart {
		if p.stock >= c {
			ans += c * p.price
		} else {
			return 0, false
		}
	}
	return float32(ans), true
}

type couponCheckingout struct {
	couponRate float32
}

func (c *couponCheckingout) calculateMoney(cart map[*product]int) (float32, bool) {
	ans := 0
	for p, c := range cart {
		if p.stock >= c {
			ans += c * p.price
		} else {
			return 0, false
		}
	}
	return float32(ans) * c.couponRate, true
}

type shoppingCartSystem struct {
	record      map[*product]int
	strategyMap map[string]checkingoutStrategy
}

func (s *shoppingCartSystem) addProduct(p *product) {
	s.record[p] += 1
}
func (s *shoppingCartSystem) removeProduct(p *product) {
	if _, ok := s.record[p]; ok {
		s.record[p] -= 1
	}
}
func (s *shoppingCartSystem) checkout(primeMember bool) float32 {
	ans := float32(0)
	yes := false
	if primeMember {
		ans, yes = s.strategyMap[discount].calculateMoney(s.record)
	} else {
		ans, yes = s.strategyMap[general].calculateMoney(s.record)
	}
	if yes {
		s.adjustStock()
		return ans
	} else {
		return ans
	}
}
func (s *shoppingCartSystem) adjustStock() {
	for p, c := range s.record {
		p.stock -= c
	}
}

func main() {
	g := &generalCheckingout{}
	c := &couponCheckingout{couponRate: 0.8}
	s_map := map[string]checkingoutStrategy{
		discount: c,
		general:  g,
	}
	sc := shoppingCartSystem{record: map[*product]int{}, strategyMap: s_map}
	p1 := &product{name: "nike shoes", stock: 10, price: 100}
	p2 := &product{name: "under armer", stock: 5, price: 50}
	sc.addProduct(p1)
	sc.addProduct(p2)
	fmt.Println(sc.checkout(true))
	fmt.Println(p1.stock, p2.stock)

}
