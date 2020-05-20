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
 
The syntax of the language doesn't have anything surprising, maybe the only particular thing is the ability to capture declarations also in the if.
!code(examples/selection.go)
!code(examples/loops.go)