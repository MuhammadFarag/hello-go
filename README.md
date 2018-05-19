# Hello Go

## Interesting concepts

* Unused variables result in compilation error.

## Variable declaration

```go
var a int // initialized to zero
a = 2
b := 3
```

Any variable in Go is initialized to a value, which is all bits set to zero by default. For numeric values that would be zero, for a string that would be an empty string.

The `:=` operator creates and assign a variable. `var a int` `var a int = 0`, `var a = 0` and `a := 0` have the same result. Note that Go has type inference, so you don't need to declare types explicitly.

Go doesn't have casting but instead it has conversion, which means you are creating a new variable and convert to it. You can even do that on declaration, so for an example you can say `a := float64(1.0)`