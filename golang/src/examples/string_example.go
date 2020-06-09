package main

import "fmt"

func main() {
	var str = "ehilà!😁!"

	fmt.Println("--- --- Loop with int and len() --- ---")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d: %q\n", i, str[i])
	}

	fmt.Println("--- --- Loop with range() --- ---")
	for i, c := range str {
		fmt.Printf("%d: %q\n", i, c)
	}
}
