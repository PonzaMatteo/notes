package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// normal if
	rnd := rand.Float64()
	if rnd < 0.01 { // () brackets are not required
		fmt.Println("Woow!")
	} else { // but {} are!
		fmt.Println("....")
	}

	// capturing variable
	if rnd := rand.Int(); rnd%2 == 1 {
		fmt.Printf("%d is odd", rnd)
	} else {
		fmt.Printf("%d is even", rnd)
	}
	// note that a declaration would not compile since the type of rnd id float64
	// rnd := rand.Int()

	switch rnd := rand.Int(); rnd % 2 {
	case 0:
		fmt.Printf("%d is even", rnd)
	case 1:
		fmt.Printf("%d is odd", rnd)
	default:
		fmt.Printf("That's pretty odd?!")
	}
}
