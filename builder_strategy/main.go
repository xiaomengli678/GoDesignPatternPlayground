package main

import "fmt"

type PizzaInterface interface {
	GetDescription() string
	GetPrice() int
}

type Pizza struct {
	Size     string
	Toppings []string
}

type PizzaBuilder struct {
	pizza *Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{pizza: &Pizza{}}
}
func (b *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	b.pizza.Size = size
	return b
}
func (b *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	b.pizza.Toppings = append(b.pizza.Toppings, topping)
	return b
}
func (b *PizzaBuilder) Build() *Pizza {
	return b.pizza
}

func (p *Pizza) GetDescription() string {
	return fmt.Sprintf("%s pizza with toppings: %v", p.Size, p.Toppings)
}
func (p *Pizza) GetPrice() int {
	price := 0
	switch p.Size {
	case "Small":
		price = 8
	case "Large":
		price = 12
	}
	return price
}

type ExtraCheeseDecorator struct {
	pizza PizzaInterface
}

func (e *ExtraCheeseDecorator) GetDescription() string {
	return e.pizza.GetDescription() + ", extra cheese"
}
func (e *ExtraCheeseDecorator) GetPrice() int {
	return e.pizza.GetPrice() + 40
}

type ExtraPepperDecorator struct {
	pizza PizzaInterface
}

func (e *ExtraPepperDecorator) GetDescription() string {
	return e.pizza.GetDescription() + ", extra pepper"
}
func (e *ExtraPepperDecorator) GetPrice() int {
	return e.pizza.GetPrice() + 80
}

func main() {
	basicPizza := NewPizzaBuilder().
		SetSize("Large").
		AddTopping("Pepper").
		AddTopping("Cheese").
		Build()

	pizzaWithCheese := &ExtraCheeseDecorator{pizza: basicPizza}
	pizzaWithCheeseAndPepper := &ExtraPepperDecorator{pizza: pizzaWithCheese}

	fmt.Println(pizzaWithCheeseAndPepper.GetDescription())
	fmt.Println(pizzaWithCheeseAndPepper.GetPrice())

	basicPizza.Size = "Small"

	fmt.Println(pizzaWithCheeseAndPepper.GetDescription())
	fmt.Println(pizzaWithCheeseAndPepper.GetPrice())
}

// 12 + 40 + 80 = 132
// 8 + 40 + 80 = 128
