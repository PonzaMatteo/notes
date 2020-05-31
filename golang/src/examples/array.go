package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [3]int
	var b = [3]int{0, 10, 0}
	fmt.Println(a, b, a == b)
	a[1] = 10
	// if the elements are comparable the array is comparable to!
	fmt.Println(a, b, a == b)
	var clone = a
	clone[1] = 21
	fmt.Println(a, clone)

	//sort.Ints(a) -> cannot use 'a' (type [3]int) as type []int: which is a slice of integer

	// a slice give access to a subsequence (or all) elements of the so called underlying array
	var c []int = a[:]
	fmt.Println(c, len(c), cap(c)) // capacity is the size of the underlying array

	c[1] = 22 // `a` is the underlying array for `c`
	fmt.Println(a, c)
	//fmt.Println(c == a) -> c == a (mismatched types []int and [3]int)

	c = append(c, 1, 2, 3, 4, 5, 6, 7, 8) // add some elements
	sort.Ints(c)
	fmt.Println("c:", c, len(c), cap(c))

	// what is the relationship between d and c underlying array?
	d := append(c, 100)
	d[1] = 1234
	fmt.Println("c:", c, len(c), cap(c))
	fmt.Println("d:", d, len(d), cap(d)) // here it is the same!
	// fmt.Println(c == d) //Invalid operation: c == d (operator == is not defined on []int)

	d = append(c, 100, 200, 300)
	d[1] = 1234
	fmt.Println("c:", c, len(c), cap(c))
	fmt.Println("d:", d, len(d), cap(d)) // here is not!
	// no assumptions about what append return!

	// standard library is not very rich of methods...
	// IDE can help with snippets, for instance for a remove operation it suggest:
	fmt.Println("d:", d, len(d), cap(d))
	d = append(d[:4], d[5:]...)
	fmt.Println("d:", d, len(d), cap(d))
}
