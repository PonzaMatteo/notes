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
!code(examples/selection.go)
!code(examples/loops.go)

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

   