package main

import (
	"fmt"
)

func main() {
	// functions / closures can be stored in variables
	greet := func(name string) {
		fmt.Println("Hello", name, "!")
	}
	greet("🌍")
	fmt.Printf("%T\n", greet)

	// deferring the call to an anonymous function
	defer func() {
		fmt.Println("This will be executed before returning")
	}()

	//
	fmt.Printf("%+v\n", CanIHaveACoffeePlease(Cappuccino()))
	fmt.Printf("%+v\n", CanIHaveACoffeePlease(Espresso(), Lungo))
	fmt.Printf("%+v\n", CanIHaveACoffeePlease(Espresso(), Double, InBigCup(), Iced()))

}

func CanIHaveACoffeePlease(coffee Coffee, options ...CoffeeOption) Coffee {
	for _, option := range options {
		coffee = option(coffee)
	}
	return coffee
}

type (
	CoffeeOption func(coffee Coffee) Coffee
	Coffee       struct {
		Cup    string
		Coffee float32
		Extra  []string
		Ice    bool
	}
)

func Espresso() Coffee {
	var coffee Coffee
	coffee.Coffee = 30.0
	coffee.Cup = "half cup"
	return coffee
}
func Lungo(coffee Coffee) Coffee {
	coffee.Coffee *= 1.5
	return coffee
}
func Double(coffee Coffee) Coffee {
	coffee.Coffee *= 2.0
	return coffee
}
func Cappuccino() Coffee {
	return CanIHaveACoffeePlease(Espresso(), WithMilk(125), InBigCup())
}
func Iced() CoffeeOption {
	return func(coffee Coffee) Coffee {
		coffee.Ice = true
		return coffee
	}
}
func WithMilk(ml float32) CoffeeOption {
	return func(coffee Coffee) Coffee {
		coffee.Extra = append(coffee.Extra, fmt.Sprintf("%v of Milk", ml))
		return coffee
	}
}
func InBigCup() CoffeeOption {
	return func(coffee Coffee) Coffee {
		coffee.Cup = "big cup"
		return coffee
	}
}
