# Object Oriented Programming

So far we learned how to declare new types and how to declare methods for them.
Go support the mechanisms of OOP in unusual and opinionated way:

## Encapsulation
There are two level of visibility:
- package level: the member name start with a lower capital letter
- exported: the member name is capitalized

!code(examples/oop/encapsulation.go)

!code(examples/encapsulation.go)

##  Inheritance
Go allow Inheritance though composition.
!code(examples/inheritance/inheritance.go)
 