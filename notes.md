In Go, functions and methods are related but serve different purposes and have different definitions.

Functions: A function is a piece of code that is called by name. It can be passed data to operate on (i.e., the parameters) and can optionally return data (the return value(s)). It is independent and can be defined outside of a type, or even outside of a package. For instance:

```
func add(a int, b int) int {
    return a + b
}
```

In this example, add is a function that takes two integers as parameters and returns their sum.

Methods: A method in Go is a function that is defined on a specific type, known as the receiver type. A method is invoked on a variable of the receiver type. For instance:

```
type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
```

Here, area is a method defined on the type Circle. It can be invoked using a variable of type Circle. The receiver (c Circle) allows the method to access the fields of the type it's defined on.

The main difference between a function and a method in Go is that a method has a receiver type, allowing it to work with the type it's defined on, whereas a function does not have a receiver and is independent. This concept of methods allows Go to have something akin to object-oriented programming, as methods can access and manipulate the state of the type they are associated with.