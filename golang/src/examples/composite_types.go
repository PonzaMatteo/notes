package main

import "fmt"

type drink struct {
	name string
	icon string
}

func (f *drink) AddIce() {
	fmt.Printf("adding ice to %s\n", f.name)
	f.name = "iced " + f.name
	f.icon += "\U0001F9CA"
}

func (f drink) AddStones() drink {
	fmt.Printf("adding stones to %s\n", f.name)
	f.name = "classy " + f.name
	f.icon += "\U0001F94C"
	return f
}

func main() {
	// define a new struct / variable
	var beer struct {
		name string
		icon string
	}
	// declaration initialize the variable to "all zero values"
	fmt.Println(beer)

	beer.name = "beer"
	beer.icon = "🍺"
	// using custom types
	whiskey := drink{name: "whiskey", icon: "🥃"}

	coffee := drink{
		name: "coffee",
		icon: "☕️",
	}
	fmt.Println(beer, whiskey, coffee)

	coffee.AddIce() // note: add ice wants a pointer, but the compiler help to easy the syntax and it is possible to call it on a struct
	newWhiskey := whiskey.AddStones()
	fmt.Println(coffee, whiskey, newWhiskey)

	//same syntax sugar works also in the other way
	var lastWhiskey *drink = &whiskey
	anotherOne := lastWhiskey.AddStones()
	lastWhiskey.AddIce()
	fmt.Println(whiskey, lastWhiskey, anotherOne)

	// remember that beer is not defined as drink?
	// beer.AddIce() does not compile!
	var lastOne drink = beer // The compiler check that the conversion is safe!
	lastOne.AddIce()         // btw: are you sure you want to do that?!?!?
	fmt.Println(beer, lastOne)
}
