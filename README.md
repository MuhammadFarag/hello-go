# Hello Go

### Interesting concepts

* Unused variables result in compilation error, but unused constants do not.
* We might need to pay attention of the order of fields in a `struct` to optimize memory allocation due to **alignment**.
* I am already confused on which case to use for *variables*, *constants*, *structs* etc. Will update this when I figure it out.
* Following the above point: The fact that go encourages short variable names bugs me. In my point of view makes the code a tad harder to reason about.

### Variable declaration

```go
var a int // initialized to zero
a = 2
b := 3
```

Any variable in Go is initialized to a value, which is all bits set to zero by default. For numeric values that would be zero, for a string that would be an empty string.

The `:=` operator creates and assign a variable. `var a int` `var a int = 0`, `var a = 0` and `a := 0` have the same result. Note that Go has type inference, so you don't need to declare types explicitly.

Go doesn't have casting but instead it has conversion, which means you are creating a new variable and convert to it. You can even do that on declaration, so for an example you can say `a := float64(1.0)`

I searched for something similar to Scala's `val` the closest thing is `const`, which you can use as `const a = 1`

### Struct

Declaring a new type *person*

```go
type person struct {
	name string
	age  int
}
```

Instantiating an instance of *person*

```go
a := person{
	name: "Some name",
	age:  10,
}
```

Finally there is a notion of anonymous struct

```go
b := struct {
	description string
}{
	description: "This is an anonymous struct!",
}
```

Notes:

* we can perform conversion from one *struct* to the other if they have the same exact fields.
* anonymous *struct* doesn't require explicit conversion, if types are identical


### Pointers
* Everything in Go is pass by value.
* Pointers is for pass by reference.
* We use `&` similar to *C* to get the address of a var.
* We use `*` similar to *C* again to get the value that a pointer is pointing to.


### Constants
* Constants have a parallel type system, they could have a `type` or a `kind`!
* `iota` is an interesting concept, the simplest use for it is to create an enum like constants with incremental value.
* `iota` can be used for different increments and even with the shift operator  [iota](https://github.com/golang/go/wiki/Iota)

```go
const (
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)
```
Which can also be written as
```go
const (
	c0 = iota  // c0 == 0
	c1         // c1 == 1
	c2         // c2 == 2
)
```
