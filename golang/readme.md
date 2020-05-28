# 01 - Some Examples
Before starting to dive into the language let's give a look to some simple program.

## Hello World
```go
package main

import "fmt"

func main()  {
	fmt.Print("Hello 🌏")
}
```


## HTTP Server
```go
failed to read file to include: open ./src/example/01_http_server.go: no such file or directory```


## HTTP Client
```go
failed to read file to include: open ./src/example/01_http_client.go: no such file or directory```


## Another HTTP Server
```go
failed to read file to include: open ./src/example/02_http_server.go: no such file or directory```


## Another HTTP Client
```go
failed to read file to include: open ./src/example/02_http_client.go: no such file or directory```


## CLI Tool
Let's see a program that allow to process a markdown file and include some source code files as code blocks. The placeholder will be in the at the beginning of a new line with the format `!code(file_name)`
```go

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var fileName = flag.String("i", "", "name of the input file")
var srcDir = flag.String("d", "./", "directory where to look for the source code examples")

func main() {
	flag.Parse()
	input, err := Input(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	content, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	codePattern := regexp.MustCompile("!code\\((.*\\..*)\\)")
	filePattern := regexp.MustCompile("\\((.*)\\.(.*)\\)")
	includes := codePattern.FindAllString(string(content), -1)
	var substituteMap = make(map[string]string)
	for _, include := range includes {
		tokens := filePattern.FindStringSubmatch(include)
		fileName := fmt.Sprintf("%s.%s", tokens[1], tokens[2])
		content, err := ioutil.ReadFile(*srcDir + fileName)
		if err != nil {
			content = []byte("failed to read file to include: " + err.Error())
		}
		substituteMap[include] = CodeBlock(tokens[2], string(content))
	}


	var out = string(content)
	for key, value := range substituteMap {
		out = regexp.MustCompile(regexp.QuoteMeta(key)).ReplaceAllString(out, value)
	}
	fmt.Print(out)
}

func CodeBlock(ext , content string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("```%s\n", ext))
	builder.WriteString(content)
	builder.WriteString(fmt.Sprintf("```\n"))
	return builder.String()
}

func Input(fileName *string) (io.ReadCloser, error) {
	if fileName != nil && *fileName != "" {
		input, err := os.Open(*fileName)
		if err != nil {
			return nil, fmt.Errorf("failed to open input file: %v", err)
		}
		return input, nil
	}
	return os.Stdin, nil
}
```


# The Language

## Brief History 
The draft of the language started in 2007 Robert Griesemer, Rob Pike and Ken Thompson and became an open source project in 2009. What is the purpose of the project?
> Go was born out of frustration with existing languages and environments for the work we were doing at Google. Programming had become too difficult and the choice of languages was partly to blame. One had to choose either efficient compilation, efficient execution, or ease of programming; all three were not available in the same mainstream language. 

> Go addressed these issues by attempting to combine the **ease** of programming of an interpreted, dynamically typed language with the **efficiency** and **safety** of a statically typed, compiled language. It also aimed to be modern, with **support for networked and multicore computing**. Finally, working with Go is intended to be **fast**: it should take at most a few seconds to build a large executable on a single computer. 

[More Info](https://golang.org/doc/faq#history)

**TA**:
 - ~10 Years history -> ecosystem is not that rich? 
 - GO try to make it easy to build simple, reliable and efficient software
 
**Fun Fact**:
 (maybe not that fun..)  in official go faq website, there are 37 "Why" question, and 13 of them are "Why [...] not" questions... 
 
 ## Control Structures
 
The syntax of the language doesn't have anything surprising, maybe the only particular thing is the ability to capture declarations also in the if/switch.
```go
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

	switch rnd := rand.Int(); rnd%2 {
	case 0:
		fmt.Printf("%d is even", rnd)
	case 1:
		fmt.Printf("%d is odd", rnd)
	default:
		fmt.Printf("That's pretty odd?!")
	}
}
```

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// each of the tree section is optional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var fruits = []string{"🍏", "🍎", "🍐", "🥑"}
	for i := range fruits {
		fmt.Println(fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// no condition will be "infinite" loop
	for {
		fmt.Print(".")
		if rand.Float64() < 0.5 {
			break
		}
	}
	fruits = append(fruits[:1], fruits[2:]...)
}
```


## Types

### Basic types
- string
- bool.
- numeric types:
    - integer: int, rune, int(8|16|32|64)
    - unsigned: byte, u + one of the above
    - uintptr: [probably don't need to use](https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang)
    - floating point: float32, float64 
    - complex numbers: complex64, complex128.
    
Go does not provide implicit conversion between numeric types: [why?](https://golang.org/doc/faq#conversions)

`string` can contain arbitrary bytes, but in general literal strings almost always contain UTF-8 characters (Go source file must be written in UTF-8). `rune` is "UTF-8 code point", without going too much into the details we can think of a rune a character. [details](https://blog.golang.org/strings).

   