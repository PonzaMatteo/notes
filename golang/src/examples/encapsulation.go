package main

import (
	"examples/oop"
	"fmt"
)

func main() {
	var foo = oop.Foo{Exported: 1}
	// foo.private -> Unexported field 'private' usage
	fmt.Println(foo)
	oop.IsPublic()
	// ext.isNotPublic() -> Unexported function 'isNotPublic' usage
}

