package main

import (
	"examples/pkg"
	"fmt"
)

func main() {
	var foo = pkg.Foo{Exported: 1}
	// foo.private -> Unexported field 'private' usage
	fmt.Println(foo)
	pkg.IsPublic()
	// ext.isNotPublic() -> Unexported function 'isNotPublic' usage
}
