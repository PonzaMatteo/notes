package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// each of the tree section is optional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var fruits = strings.Split("🍏,🍎,🍐,🥑", ",")
	for i := range fruits {
		fmt.Println(fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// no condition will be "infinite" loop
	for  {
		fmt.Print(".")
		if rand.Float64() < 0.5 {
			break
		}
	}
}
