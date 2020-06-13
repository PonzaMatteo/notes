# 01 - Some Examples
Before starting to dive into the language let's give a look to some simple program.

## Hello World
```go
package main

import "fmt"

func main() {
	fmt.Print("Hello 🌏")
}
```


## HTTP Server
```go
package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":12345"
	log.Println("start listening ", addr)
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(addr, nil)
	panic(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %s %s", r.Method, r.URL.String())
	_, err := w.Write([]byte("HELLO!"))
	if err != nil {
		log.Print("failed to send response")
	}
}
```


## HTTP Client
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	path := "hello"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	resp, err := http.Get("http://localhost:12345/" + path)
	if err != nil {
		log.Printf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		return
	}
	fmt.Print(string(content))
}
```


## Another HTTP Server
```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Food struct {
	Name string
	Icon string
}

var menu = []Food{
	{"Fries", "🍟"},
	{"Apple", "🍏"},
	{"Avocado", "🥑"},
	{"Pizza", "🍕"},
}

func main() {
	http.HandleFunc("/menu", Menu)
	err := http.ListenAndServe(":12345", nil)
	panic(err)
}

func Menu(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %s %s", r.Method, r.URL.String())
	err := json.NewEncoder(w).Encode(menu)
	if err != nil {
		log.Printf("failed to send response: %v", err)
	}
}
```


## Another HTTP Client
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Food struct {
	Name string
	Icon string
}

func main() {
	start := time.Now()
	defer func() {
		log.Println("total running time: ", time.Now().Sub(start).Milliseconds(), "ms")
	}()
	resp, err := http.Get("http://localhost:12345/menu")
	if err != nil {
		log.Printf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close() // ignoring error
	var food []Food
	err = json.NewDecoder(resp.Body).Decode(&food)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		return
	}
	fmt.Println(food)
}
```


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

func CodeBlock(ext, content string) string {
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
The draft of the language started in 2007 by Robert Griesemer, Rob Pike and Ken Thompson and became an open source project in 2009. What is the purpose of the project?
> Go was born out of frustration with existing languages and environments for the work we were doing at Google. Programming had become too difficult and the choice of languages was partly to blame. One had to choose either efficient compilation, efficient execution, or ease of programming; all three were not available in the same mainstream language. 

> Go addressed these issues by attempting to combine the **ease** of programming of an interpreted, dynamically typed language with the **efficiency** and **safety** of a statically typed, compiled language. It also aimed to be modern, with **support for networked and multicore computing**. Finally, working with Go is intended to be **fast**: it should take at most a few seconds to build a large executable on a single computer. 

[More Info](https://golang.org/doc/faq#history)

**TA**:
 - ~10 Years history -> ecosystem is not that rich? 
 - GO try to make it easy to build simple, reliable and efficient software
 
**Fun Fact**:
 (maybe not that fun..)  in official go faq website, there are 37 "Why" question, and 13 of them are "Why [...] not" questions... 

## Program structures

### Control Structures
The syntax of the language doesn't have anything surprising, maybe unusual feature is the ability to capture declarations also in the if/switch.
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

	switch rnd := rand.Int(); rnd % 2 {
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

### Composite types
- struct
- pointers
- data structures:
    - array
    - slice
    - map
    
#### Struct And Pointers
```go
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
```


#### Array and Slices
Arrays are fixed length, and it is part of the type. Usually they are not used directly, using `slices` instead.
```go
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
```
 

### Other Types
 - functions
 - channels
 - interfaces
 
#### Functions
Functions are first class values 
```go
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
```
 

- (Do Not Fear First Class Functions)[https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions]

#### Interfaces

Go takes an unusual approach to interfaces, they are **satisfied implicitly**. So we don't declare which interfaces as type implements, we just implement the methods.
Include example for:
- basic usage
- embedding
- type as couple
- writer / reader flexibility

# Object Oriented Programming

So far we learned how to declare new types and how to declare methods for them.
Go support the mechanisms of OOP in unusual and opinionated way:

## Encapsulation
There are two level of visibility:
- package level: the member name start with a lower capital letter
- exported: the member name is capitalized

```go
package oop

// the type is exported
type Foo struct {
	// Capitalized field: it is exported
	Exported int
	// field is not visible out of main package
	private int
}

func IsPublic()    {}
func isNotPublic() {}
```


```go
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
```


##  Inheritance
Go allow Inheritance though composition.
```go
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type handler struct {
	//embedding the http.ServeMux allow to use all its the methods
	//on a `handler` end extend it with new behaviour
	*http.ServeMux
}

func NewHandler() *handler {
	return &handler{http.DefaultServeMux}
}

//Method that allow to add as handler a function that return (value, error) and send the response\
//in json.
func (h *handler) HandleJson(patter string, serve func(r *http.Request) (interface{}, error)) {
	h.HandleFunc(patter, func(w http.ResponseWriter, r *http.Request) {
		resp, err := serve(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp = NewError(err)
		}
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
}

func main() {
	handler := NewHandler()
	//Note: HandleFunc is defined on *ServeMux
	handler.HandleFunc("/hello", hello)
	handler.HandleJson("/hello-json", helloJson)
	//Note: ListenAndServe accept an interface http.Handler that our handler inherit from the embedded *ServeMux
	err := http.ListenAndServe(":12345", handler)
	panic(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	message, err := GetMessage(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(NewError(err))
		return
	}
	_ = json.NewEncoder(w).Encode(message)
}

func helloJson(r *http.Request) (interface{}, error) {
	return GetMessage(r.RequestURI)
}

//Used for sending error as json
type Error struct {
	Error string
}
func NewError(err error) *Error {
	return &Error{Error: err.Error()}
}

//Simulate method that can fail
func GetMessage(str string) (map[string]string, error) {
	if rand.Float32() >= 0.5 {
		return nil, fmt.Errorf("failed to [...]")
	}
	return map[string]string{"msg": str}, nil
}
```

 