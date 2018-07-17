# Hello Go

[TOC levels=4]: #
# Table of Contents
- [Interesting concepts](#interesting-concepts)
- [Code style](#code-style)
- [Commandline](#commandline)
    - [Arguments](#arguments)
    - [Flags](#flags)
- [Variable declaration](#variable-declaration)
    - [Interesting aspects of variable declaration:](#interesting-aspects-of-variable-declaration)
- [Type](#type)
- [Struct](#struct)
    - [Type embedding](#type-embedding)
- [Pointers](#pointers)
- [Constants](#constants)
- [Functions](#functions)
    - [Higher order functions](#higher-order-functions)
    - [Errors](#errors)
- [Collections](#collections)
    - [Arrays](#arrays)
    - [Slices](#slices)
    - [Maps](#maps)
- [File IO](#file-io)
    - [Reading files](#reading-files)
- [Switch](#switch)
- [Resources](#resources)

### Interesting concepts
* Unused variables result in compilation error, but unused constants do not.
* We might need to pay attention of the order of fields in a `struct` to optimize memory allocation due to **alignment**.
* Following the above point: The fact that go encourages short variable names bugs me. In my point of view makes the code a tad harder to reason about.

### Code style
* If an identifier name start with an upper case letter, it is `exported` which means it is visible outside of its package. Yes, visibility is determined by the case of the first letter.
* The letters of acronyms are rendered in the same case. e.g. `someAPI` not `someApi`
* File names should not include `-`, `_` or use camel case. Yes, just short lower case name. Of course it could be `shortname.go` or `shrtnm.go`. The reason why underscores are not a good idea, Go build tool actually uses underscores to distinguish between different architectures.

### Commandline

#### Arguments
The `main` function which acts as the entry point of the program doesn't take arguments. To read arguments from command line you will find them in `os.Args`

```go
import "os"

func main() {
	for i:=0; i< len(os.Args); i++ {
		println(os.Args[i])
	}
}
```

#### Flags

```go
f := flag.String("f","default value of this flag", "some interesting flag")
flag.Parse()
println(*f)
arguments :=
```

`flag` package gives a convenient function to reach commandline arguments without the flags using `flag.Args()`. Of course, you can declare as many flags as you want. You can also use other types such a boolean flag using `flag.Bool(...)` function. One thing to note, the call to `flag.Parse()` must proceed calling `flag.Args()`

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

#### Interesting aspects of variable declaration:

```go
var a, b, c = true, 1.5, "hello"
var j = 0
i, j := 0, 1 // reassign j and MUST declare at least one new variable
i, j = j, i //swap variables
```

### Type

`type` is not a type alias, which is different than Scala for an example. Which means the new type is not compatible, assignable or comparable to the original type. In my mind this is better than type aliasing because type aliasing, if overused, is in many cases redundant and confusing.
`type timeInSeconds int` and `type timeInMilliSeconds float` are not compatible. You need to cast from one to the other (or just for correctness to `int` and then to the other) to convert.

```go
type timeInSeconds int
type timeInMilliSeconds int

var seconds timeInSeconds = 10
var millis timeInMilliSeconds = timeInMilliSeconds(1000 * int(seconds))
```

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

We can instantiate person without naming the fields, given that we provide all fields in the correct order. If we don't provide any of the fields, i.e. use `{}` we end up with an object with all its fields set to its corresponding zero value.

```go
a1 := person {"name 1", 1}
a2 := person {}
```

There is also the built in function `new` which creates a pointer to a zero allocation of the type, thus the following two are identical
```go
a2 := &person {}
a3 := new(person)
```


Finally there is a notion of anonymous struct

```go
b := struct {
	description string
}{
	description: "This is an anonymous struct!",
}
```

#### Type embedding

Let's look at the following example using the person type we have defined earlier, we define a new type player.

```go
type player struct {
	person
	favouriteGame string
	int
}
```

We notice that we have two *anonymous* fields, one of type `person` and the other of type `int`. The question now, how do you access those? The field names are implicit and they are the same as the type.

```go
p := player{person{"jack", 23}, "game", 1}

fmt.Println(p)              // {{jack 23} game 1}
fmt.Println(p.person.name)  // jack
fmt.Println(p.int)          // 1
```

Go also give us another interesting feature, you can access the fields of the embedded type, `person` in our case, directly using `p.name` for example... just like that.

```go
fmt.Println(p.name)         // jack
```

If both the embedded and the embedding type has fields holding the same name, then you will need to use `p.name` for the embedding type and `p.person.name` for the embedded type. Let's redefine our type player and create an instance of it.

```go
type player struct {
	person
	favouriteGame string
	int
	name string
}

p := player{person{"jack", 23}, "game", 1, "Mo"}
```

As you can see we have defined `name` twice. In that case the one defined in the `player` struct will hide the one from `person` struct.

```go
fmt.Println(p)              // {{jack 23} game 1 Mo}
fmt.Println(p.person.name)  // jack
fmt.Println(p.name)         // Mo
fmt.Println(p.int)          // 1
```

Notes:

* we can perform conversion from one *struct* to the other if they have the same exact fields.
* Anonymous *struct* doesn't require explicit conversion, if types are identical.
* Field order is part of a type. Different field orders means different types, even if the fields are identical.
* If we are trying to access a field in a pointer to a struct, we don't need to explicitly use `*` to get to the contents of that stucts, we can use `.` directly. That is the following two lines of codes will print the exact same thing

```go
fmt.Println(b.description)
fmt.Println((&b).description)
```

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

### Functions

Functions may have multiple results, i.e. it might return more than one value separated by a comma. It is possible to have a name for each of the returned results for documentation purposes mostly.

```go
func threeTimes(a, b string) (r1 string, r2 string, r3 string) {
	r := fmt.Sprintf("%s %s", a, b)
	return r, r, r
}
```

It is possible to omit the arguments of the return statement if the named results if they have been assigned within the function, but that is not recommended because it makes code harder to reason about.

```go
func namedResult() (r string){
	r = "Hello Go"
	return
}
```
Go doesn't support tail recursion optimization. However, function stacks in go is dynamic which mitigate the risk of stack over flow, thus most of reasonably bound or terminating recursion calls are safe.

#### Higher order functions
Higher order functions are functions that take functions as their argument. Since functions are first class citizens in Go, they have a type and they can be passed around. a function type is a representation for its arguments and its return. In this case we are passing behaviour as a parameter.

```go
func consumeBehaviour(f func() string) {
	println(f())
}

consumeBehaviour(namedResult)
```

`consumeBehaviour` takes an argument `f` which takes no arguments and returns a string. The function `newResult` we defined earlier fits the requirement.

#### Errors

Function may return errors. Since Go has multiple results by convention the last result is an error indicator or an error. An error indicator is a boolean that will evaluate to true if there is no errors.

```go
func somethingMightGoWrong()(string, bool) {
	return "", false
}

_, ok := somethingMightGoWrong()

if !ok {
	// an error has occurred
}
```

An error result might be returned if we need more information on the error

```go
func somethingElseMightGoWrong()(string, error) {
	return "", fmt.Errorf("error: %s", "error description")
}
_, err := somethingElseMightGoWrong()

if err != nil {
    // an error has occurred
}
```

We can also check for error and execute accordingly.

```go
if _, err := somethingElseMightGoWrong(); err !=nil {
    // do something
}
```

The downside with that though is that the function results are only visible within the scope of the if statement, which might be ok in some cases.

We can define specific errors for comparison purposes. That is in case the client might want to take different actions based on different errors.

```go
var SpecificError = errors.New("Specific error")

func returnSpecificError() error{
	return SpecificError
}

if err:= returnSpecificError(); err == SpecificError {
		println(err.Error())
	}
```

### Collections
#### Arrays

We can initialize an array using either of the following

```go
var array [5]string

array[0] = "a"
array[1] = "b"
array[2] = "c"
array[3] = "d"
array[3] = "e"
```

or
```go
array := [5]string{"a", "b", "c", "d", "e"}
```

In the last example, we can replace the length of the array `5` with `...` which will result in automatically deducting the size from the number of elements.

We can also skip indexes. The skipped indexes will be set to the zero value of the type, in the following example an empty string. The type of the following array is `[3]string` since the index ` is implicitly set to empty string.
```go
a := [...]string{0:"a", 2:"b"}
```

We can use the traditional for loop to loop through the array

```go
for i := 0; i < len(array); i++ {
	fmt.Println(i, array[i])
}
```

of the more Go'ish alternative
```go
for i, s := range array {
    fmt.Println(i, s)
}
```

##### Notes
- The size of the array is part of the type. i.e. the type is `[5]int` not an array of `int`. Thus `[5]int` and `[3]int` are not compatible types and variable of one type can't receive a value of the other.
- Arrays are rarely used because the cases of fixed size arrays are very limited. Slices are a more powerful alternative.

#### Slices

A *slice* is a data structure packed by an array (for simplicity).

Declaring a slice:

```go
var zeroSlice []int
```

```go
slice := make([]int, 5)
```

```go
sliceWithCapacity := make([]int, 5, 8)
```

The third argument in the above example is *capacity*. When you declare a slice it has a length but also has a capacity, which is adjacent memory locations that are not utilized. This gives the slice some of the characteristics of dynamic data structures.

```go
sliceWithCapacity = append(sliceWithCapacity, 10, 20, 30, 40)
```

The above call appends the elements to the slice. If you are keeping count, the number of elements exceed the capacity, what the slice does is that it doubles the capacity and create a new packing array for those values. It keeps doubling until the capacity exceeds a 1000 elements and in that case, it increases the capacity by 25%.

You can take a "slice" of a slice, or in other words a subset of a slice using the syntax `subSlice := slice[4:5]`. Note that those two share the same packing array. Which means changes to one, affect the other. The only thing that can prevent this ripple effect is limiting the sub-slice size to minimum, i.e. the length and capacity are equal, and then append to the sub-slice. To do that we use `subSlice := slice[4:5:5]`. After the append, the packing array was copied and thus any changes to the sub-slice will not affect the source slice.

#### Maps

The syntax to declare a zero map is
```go
map1 := make(map[string]int)
```
We can also initialize on declaration
```go
map1 := map[string]int{
	"Google":   4,
	"LinkedIn": 5,
}
```

To add an item to a map we use `map1["Samsung"] = 3`, to delete an item from a map `delete(map1, "Samsung")`. Finding an item in a map is more interesting, as the result is two variables. The first is the zero value of the value and the second is a boolean that indicates whether or not the element was found.
```go
v, found := map1["Toshiba"]
```
You may just use `v := map1["Toshiba"]`. But, you wouldn't know if the value was actually `0` or not found.

By convention the second result is not called found but rather `ok` and we usually end up with the following pattern:

```go
if n, ok := map1["Toshiba"]; !ok { ... }
```

Finally iterating over a map using for
```go
for k, v := range map2 {
	fmt.Println(k, v)
}
```

##### Sets

This is intentionally a subset of Maps. There is no sets in Go, but since map keys can't duplicate, maps can be used to represent sets. For an example to represent a set of strings we can use `map[string]bool` This is a "set of strings"

### File IO
#### Reading files
##### Scanning files

```go
file, err := os.Open(fileNameAndPath)
if err != nil {
	println("Failed to open file: ", fileName, ". Error: ", err)
} else {
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		println(i, ": ", scanner.Text())
	}
}
```

##### Reading files to memory

```go
file, err := ioutil.ReadFile(fileName)
if err != nil {
	println("Failed to open file: ", fileName, ". Error: ", err)
}
for i, line := range strings.Split(string(file), "\n") {
	println(i, ": ", line)
}
```

### Switch
No fallthrough between cases by default, even though you can override that using `fallthrough` keyword which has to be the last statement in the case.
I am surprised that the following code compiled, or at least showed no warnings, given that the default statement is not reachable. Even without the fall through where `println("ola") was not reachable, that was no problem at all.

```go
switch "hi" {
case "hello", "hi":
	println("hi")
	fallthrough
case "ola":
	println("ola")
default:
	println("The default")
}
```

Switch can get more interesting:
```go
x:= "hi"
switch {
case strings.Contains("hi there", x):
	println("contains hi")
default:
	println("default")
}
```

---
## Resources
1. [Ultimate Go Programming](https://www.safaribooksonline.com/library/view/ultimate-go-programming/9780134757476/)
2. [The Go Programming Language](https://www.gopl.io/)
3. [Go Documentation](https://golang.org/doc/)
4. [A tour of Go](https://tour.golang.org/list)
5. [Go wiki: Switch](https://github.com/golang/go/wiki/Switch)
6. [The Go Programming Language Specification](https://golang.org/ref/spec)
7. [Go by Example](https://gobyexample.com/)
