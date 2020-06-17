# Patterns

## Error is just a value

Although Go offer mechanism similar to try-catch (using panic and recover) the common way to handle errors and failures is simply writing function that return a couple (value, error).
The type error is a builtin interface defined as follows:
```go
package builtin
type error interface {
    Error() string
}
``` 
So we can also define custom error type just making sure they satisfy the interface defined above. 

As we have seen It is common practice for function that can fail to return a couple of values. It is common practice dealing with the error first, and then proceeding with the correct flow of execution. We can find many of such examples in the standard library:
```go
package io
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

A common pattern is, just after calling the function, to check if the error is `nil` before using the result, or just for checking  that the operation was successful.
```go
package demo
import (
	"io"
    "fmt"
    "os"
)
func foo() error {
    // do something
    var out io.Writer = os.Stdout
    _, err := fmt.Fprintf(out, "hello")
    if err != nil {
        return fmt.Errorf("failed to print message: %v",err)
    }
    return nil
}
```
This has the benefit of making the flow of execution and the error handling trivial, the code is readable and nothing is "unexpected".

As pointed out in the article (Errors are values)[https://blog.golang.org/errors-are-values] this is by far the only pattern we can use. Since errors values we can store them and process them later, we can make them part of the "computation". We can treat them as we would use any other variable.

!code(examples/errors/writer.go)

Another point is that although the definition of the interface is really simple, we can define new "error" data structures that offer more "sophisticated" behaviour, we are not limited to use the one in the standard library.

     