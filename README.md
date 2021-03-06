# Hello Go

[TOC levels=4]: #
# Table of Contents
- [Interesting concepts](#interesting-concepts)
- [Code style](#code-style)
- [Commandline](#commandline)
    - [Arguments](#arguments)
    - [Flags](#flags)
    - [Custom flags](#custom-flags)
- [Variable declaration](#variable-declaration)
    - [Interesting aspects of variable declaration:](#interesting-aspects-of-variable-declaration)
- [Type](#type)
- [Struct](#struct)
    - [Type embedding](#type-embedding)
- [Pointers](#pointers)
- [Constants](#constants)
- [Functions](#functions)
    - [Higher order functions](#higher-order-functions)
    - [Variadic function](#variadic-function)
    - [Defer](#defer)
    - [Errors](#errors)
    - [Panic and recover](#panic-and-recover)
- [Methods](#methods)
- [Pimp My Library](#pimp-my-library)
- [Interfaces](#interfaces)
    - [any or object](#any-or-object)
- [Collections](#collections)
    - [Arrays](#arrays)
    - [Slices](#slices)
    - [Maps](#maps)
- [Go Routines](#go-routines)
- [Channels](#channels)
    - [Unbuffered Channels](#unbuffered-channels)
    - [Notes](#notes)
- [File IO](#file-io)
    - [Reading files](#reading-files)
- [Switch](#switch)
- [Resources](#resources)

### Interesting concepts
* Unused variables result in compilation error, but unused constants do not show this behaviour.
* We need to pay attention to the order of fields in a `struct` to optimize memory allocation due to **alignment**.
* Following is the explanation of above point: The fact that go encourages short variable names bugs me. I think it makes the code a tad harder to reason about.

### Code style
* If an identifier name starts with upper case letter, it is `exported` which means it's visible outside of the package. Yes, visibility is determined by the case of the identifier's first letter.
* The letters of acronyms are rendered in the same case. e.g. `someAPI` not `someApi`
* File names should not include `-`, `_` or use camel case. Yes, just short lower case name. Of course it could be `shortname.go` or `shrtnm.go`. The reason why underscores are not a good idea is that Go build tool actually uses underscores to distinguish between different architectures.

### Commandline

#### Arguments
The `main` function which acts as the entry point of program doesn't take arguments. To read arguments from command line you will find them in `os.Args`

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

`flag` package gives a convenient function to reach commandline arguments without the flags using `flag.Args()`. Of course, you can declare as many flags as you want. You can also use other types, such a boolean flag using `flag.Bool(...)` function. One thing to note is that, the call to `flag.Parse()` must proceed calling `flag.Args()`

#### Custom flags
To create a custom flag we need to satisfy the `flag.Value` interface. I am using a simple example custom parsing a float. Note that we are reinventing `flag.Float64()` for demonstration purposes.

```go
type floatFlag struct {
	 float64
}

func (m *floatFlag) String() string {
	return fmt.Sprintf("%v", m.float64)
}

func (m *floatFlag) Set(s string) error {
	m.float64, _ = strconv.ParseFloat(s, 32)
	return nil
}
```

We have defined a new type `floatFlag` that satisfies the `flag.Value` interface. Thus it has the `String` and `Set` methods. This is a simple implementation without error handling. Now we need the following in our main to use this custom flag.

```go
f:= floatFlag{9} // set default value
	flag.CommandLine.Var(&f, "flag-name", "usage of this test flag")
	flag.Parse()
	fmt.Println(f.float64) // will give us the value passed to commandline 
```

Now if we call our app with `-flag-name 2.5`, the value of `f.float64` will be `2.5`


### Variable declaration

```go
var a int // initialized to zero
a = 2
b := 3
```

Any variable in Go is initialized to a value, which is all bits set to zero by default. For numeric values that would be zero, for a string that would be an empty string.

The `:=` operator creates and assigns a variable. `var a int` `var a int = 0`, `var a = 0` and `a := 0` have the same result. Note that Go has type inference, so you don't need to declare types explicitly.

Go doesn't have casting but instead it has conversion, which means you are creating a new variable and convert to it. You can even do that on declaration, so for an example you can say `a := float64(1.0)`

I have also searched for something similar to Scala's `val` and I found that the closest thing is `const`, which you can use as `const a = 1`

#### Interesting aspects of variable declaration:

```go
var a, b, c = true, 1.5, "hello"
var j = 0
i, j := 0, 1 // reassign j and MUST declare at least one new variable
i, j = j, i //swap variables
```

### Type

`type` is not a type alias, which is different than Scala for example. Which means the new type is not compatible, assignable or comparable to the original type. In my mind this is better than type aliasing because type aliasing, if it is overused, it ends up in many cases to be  redundant and confusing.
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

We can instantiate person without naming the fields, given that we provide all fields in the correct order. If we don't provide any of the fields, i.e. use `{}` we end up with an object having all its fields set to the corresponding zero value.

```go
a1 := person {"name 1", 1}
a2 := person {}
```

There is also the built in function `new` which creates a pointer to a zero allocation of the type, thus the following two declarations are identical
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

We notice that we have two *anonymous* fields, one of type `person` and the other of type `int`. The question now is, how do you access those? The field names are implicit and they are same as the type.

```go
p := player{person{"jack", 23}, "game", 1}

fmt.Println(p)              // {{jack 23} game 1}
fmt.Println(p.person.name)  // jack
fmt.Println(p.int)          // 1
```

Go also gives us another interesting feature, you can access the fields of the embedded type, `person` in our case, directly using `p.name` for example... just like that.

```go
fmt.Println(p.name)         // jack
```

If both the embedded and embedding types have fields holding the same name, then you will need to use `p.name` for the embedding type and `p.person.name` for the embedded type. Let's redefine our type player and create an instance of it.

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

* we can perform conversion from one *struct* to another if they have exactly same fields.
* Anonymous *struct* doesn't require explicit conversion, if types are identical.
* Field order is part of a type. Different field order means different types, even if the fields are identical.
* If we are trying to access a field in a pointer to struct, we don't need to explicitly use `*` to get to the contents of that struct, we can use `.` directly. That is the following two lines of codes will print the exact same thing

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

It is possible to omit the arguments of return statement, if the named results have been assigned within the function, but that is not recommended because it makes code harder to reason about.

```go
func namedResult() (r string){
	r = "Hello Go"
	return
}
```
Go doesn't support tail recursion optimization. However, function stacks in go is dynamic which mitigate the risk of stack over flow, thus most of reasonably bound or terminating recursion calls are safe.

#### Higher order functions
Higher order functions are functions that take functions as their argument. Since functions are first class citizens in Go, they have a type and they can be passed around. A function type is the representation of both its arguments and its return. In this case we are passing behaviour as a parameter.

```go
func consumeBehaviour(f func() string) {
	println(f())
}

consumeBehaviour(namedResult)
```

`consumeBehaviour` takes an argument `f` which takes no arguments and returns a string. The function `newResult` we defined earlier fits the requirement. We don't even need to pass a concrete function, we can pass an *anonymous function*

```go
consumeBehaviour(func () string {return "Hi"})
```

#### Variadic function

A variadic function can receive zero or more of its arguments. The significant thing to note here is the signature of function `func(...int)`, which means it doesn't have the same signature as `func([]int)` which is a function taking a slice of int.
```go
func variadic(x...int){
	println(fmt.Sprintf("The type of this function is: %T", variadic))              //  The type of this function is: func(...int)
	println(fmt.Sprintf("The type of argument is %T and its value is %v", x, x))    //  The type of argument is []int and its value is [1]
}
```

#### Defer

A function marked as defer is executed at the end of the enclosing function, regardless how the enclosing function ends. deferred functions are executed in the reverse order of their declaration inside the enclosing function. For example

```go
func letsDefer(s string) (r string) {
	defer func() {println("defer 1:", s ,r)}()
	r = "r-" + s
	defer func() {println("defer 2:", s ,r)}()
	return r
}
```
will yield
>defer 2: hello r-hello  
>defer 1: hello r-hello

Note that the second deferred function was executed first and both functions had the value of the result, even though it was not assigned at the time defer-1 was declared. A very intriguing use case is where one can capture a changing value at both the point of entry and exit of the function, to a more specific example, how about getting the running time of a function?

```go
func functionAsFinally() {
	defer func() func() {
		before := time.Now()
		return func() { println("Elapsed time:", (time.Now().Sub(before)).String()) }
	}()()
}
```

In summary defer makes a function act like a `finally` block for the enclosing function.

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

The downside with that is the function results are only visible within the scope of the if statement, which might be ok in some cases.

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

#### Panic and recover

How about throwing an exception? Go give us that in the form of `panic` which will halt the execution of a program an exit. It also gives us a mechanism to recover from `panic`, that is a panic in let's say a function we are consuming will not result in halting the program and give us some means to handle the panic (read exception) which is `recover`. One can think of `panic` as a `throw` statement in other programs and `recover` as a `catch` in other programs.

```go
func main() {
	defer func() {
		if p:= recover(); p!= nil{
			fmt.Printf("Recovering from our panic! %v", p)  // Recovering from our panic! It is the end of the world!
		}
	}()
	panic("It is the end of the world!")
}
```

We notice that the `recover()` built in function would result in a value that might be `nil`, that happens if no error has occurred, i.e no `panic`. Notice that recover is called within a deferred anonymous function, so it executes at the end of the function.

We can still have a more traditional error handling if we might be recovering from different types of errors, where we switch on the `panic` itself in a similar fashion to handling multiple exceptions.

```go
func main() {
	defer func() {
		switch p := recover(); p {
		case nil: // no panic has occurred
		case "It is the end of the world!":
			fmt.Printf("Recovering from our panic! %v", p)
		default:
			fmt.Printf("Recovering from unknown error! %v", p)
		}
	}()
	panic("It is the end of the world!")
}
```

### Methods
Let's create a simple type, `account` which has one field `money`, we will create one method `add` which will add some money to the account.

```go
type account struct {
	money int
}

func (a account) add(amount int) {
	a.money += amount
}
```

Let's give it a shot

```go
func main() {
	a:= account{}
	fmt.Println("Account before:", a)        // Account before: %v {0}
	a.add(15)
	fmt.Println("Account after add:", a)     // Account after add: %v {0}
}
```

What just happened? Since by default in Go function arguments are passed by value not by reference, what was passed into the add function was a copy of account. So changes in the value itself doesn't affect the value of `a` in the main function. To use more accurate language, that first argument (before the method name) is called *function receiver*. So, the function receiver is a copy.

There are two ways to go over this, let's take the more obvious if you are coming from OOP background. Instead of having the receiver of type `account` we change it to `*account`.

```go
func (a *account) add(amount int) {
	a.money += amount
}
```

Without changing anything in the code in main, we will get the result we were looking for.

```go
func main() {
	a:= account{}
	fmt.Println("Account before:", a)        // Account before: %v {0}
	a.add(15)
	fmt.Println("Account after add:", a)     // Account after add: %v {15}
}
```

Note that conveniently we didn't need to qualify `a` as `&a` (address of a).

Another solution, if you are like me coming from Functional Programming background (Scala), you'd return a copy of the receiver.

```go
func (a account) add(amount int) account {
	a.money += amount
	return a
}
```

The main function will slightly change:

```go
func main() {
	a := account{}
	fmt.Println("Account before:", a)
	a = a.add(15)
	fmt.Println("Account after add:", a)
}
```

Another point, you might consider a method as syntactic sugar for a function where that function's first argument is the receiver. You can use `account.add(a, 1)` instead of `add` on the instance of that type.

In summary, we can think of pointer function receiver as your regular methods in an OOP language, while value function receivers as static methods. In Scala those would be objects within a method. Using the last approach renders types immutable, immutability down side is memory use, but it comes with lots of benefits.

One other thing to note that the receiver type declaration and the method has to be in the same package.

### Pimp My Library

Pimp my Library is a term coined by *Martin Odersky* explaining how Scala gives you the option of enriching types that you might not own. It comes to mind when you think about how methods can only have receivers in the same package. It turns out it is possible to create methods for receivers you don't own and hence extend existing types or other libraries. There are other advantages of this of course. Let's have a look at enriching int. We will just add a method `Increment` to `int`. This is redundant of course to `+1` or `++` but it is a simple example in one hand, and it is immutable on the other.

```go
func main() {
	a := 1
	a = RichInt(a).Increment().toInt()
	println(a)
}

type RichInt int

func (n RichInt) Increment() RichInt {
	return n + 1
}

func (n RichInt) toInt() int {
	return int(n)
}
```

We had to create a new type, but if we have more methods, not just Increment, this seems like an option.

### Interfaces

There is no inheritance in Go, but there are interfaces! Interfaces have methods and any type that has the same methods *satisfies*  that interface. In that sense, that type is an instance of that interface. So, in Go, you will not find extends or with. It just happens. In my personal opinion, you lose a bit of type self-documentation. Thus, there is no override either. So, unless I discover something interesting, you might rename a method and no longer satisfy an interface that the type used to.

```go
func main() {
	var s Speaker
	s = Person{"Muhammad"}
	s.speak()
}

type Speaker interface {
	speak()
}

type Person struct {
	name string
}

func (p Person) speak(){
	println("My name is", p.name)
}
```

Just to re-iterate, in the above example the type `Person` is a `Speaker`, because it has the method `speak`.

 We can compose interfaces using type embedding.

 ```go
 type Speaker interface {
 	speak()
 }
 
type SoccerPlayer interface {
	Play()
}

type SoccerSpeaker interface {
	Speaker
	SoccerPlayer
} 
 ```

So, if some type satisfies *SoccerSpeaker* it satisfies the other two as well.

Without explicit extension there is a probability that some one would refactor a piece of code without realizing that it satisfies one interface or the other. Turns out there is an *ugly* work around. By declaring a dummy variable in your type you can make it more explicit. It might be useful as documentation. But also, it helps your tooling to scream at you when you rename a function.

```go
var _ Speaker = Person{}
```


#### any or object

In other languages such as Java or Scala there is a type that represents anything. In Java it is `object` and in Scala it is `any`. You can have the same effect by using `interface {}` in Go, since any type satisfies this interface. You will need type matching or some reflection magic to make use of the value you have based though, which is the same as any other typed language.


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

To add an item to a map we use `map1["Samsung"] = 3`, to delete an item from a map `delete(map1, "Samsung")`. Finding an item in a map is more interesting, as the result is two variables. The first is the value and the second is a boolean that indicates whether or not the element was found.
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

This is intentionally a subset of Maps. There is no sets in Go, but since map keys can't duplicate, maps can be used to represent sets. For example to represent a set of strings we can use `map[string]bool` This is a "set of strings"

### Go Routines

We can think of a go routine as a lightweight thread. Everything in go is executed in a `goroutine`, note it is written as one word. When you start the main function it is on a goroutine, the main goroutine. To span a new goroutine all you need to do is use the magic word `go`.

Let's start with a simple program to print count up with a delay.

```go
func main() {
	printCountUp("A")
	printCountUp("B")
}

func printCountUp(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s-%d ", prefix,  i)
		time.Sleep(100 *  time.Millisecond)
	}
}
```
Running this will output the expected result: `A-0 A-1 A-2 A-3 A-4 A-5 A-6 A-7 A-8 A-9 B-0 B-1 B-2 B-3 B-4 B-5 B-6 B-7 B-8 B-9`

Let's introduce the `go` keyword

```go
go printCountUp("A")
printCountUp("B")
```

Now we get `B-0 A-0 A-1 B-1 A-2 B-2 A-3 B-3 A-4 B-4 A-5 B-5 A-6 B-6 A-7 B-7 A-8 B-8 A-9 B-9` Which is expected and the program would run almost at half the time. Now let's go a bit further

```go
go printCountUp("A")
go printCountUp("B")
```

The program will terminate immediately and we will get no output. The reason being that the parent goroutine, the main goroutine in this case, has terminated and took its children down with it.

There are many reasons why you would in a real world program to go to multiple routines in the same time and waiting for them all to finish. For that we have a `WaitGroup`. Let's look at a better implementation.

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		printCountUp("A")
	}()
	go func() {
		defer wg.Done()
		printCountUp("B")
	}()
	
	wg.Wait()
}
```
This will print again `B-0 A-0 A-1 B-1 A-2 B-2 A-3 B-3 A-4 B-4 A-5 B-5 A-6 B-6 A-7 B-7 A-8 B-8 A-9 B-9 ` as expected. Note that we can call `wg.Add` twice with the argument set to `1` in each time... It is just the number of goroutines to wait for.

If we would have called `wg.Add(1)` only once, we would have finished execution whenever one of the two routines finished executing. If we would have used `wg.Add(3)`, in this case at least, we will get a deadlock exception `fatal error: all goroutines are asleep - deadlock!`.

Reference:  [How to Wait for All Goroutines to Finish Executing Before Continuing](https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/)

### Channels

GoRoutines mechanism of parallelism is already a leap forward compared to many other languages since it is a core concept of the language. But Go goes a step further by giving a mechanism for goroutines to communicate. This is as dangerous as it is useful. I would highly suggest not to use this to synchronize go routine execution. Once you do that for the most part you lose the benefits of parallelism but keep the complexity. Of course, there might be a reason for it, but you really need to think hard and find a justification for the cost of the complexity in that case.

#### Unbuffered Channels

Let's look at the simple printCountUp example from the previous section. Let's split counting up and printing into two different functions. One to count up and the other to print. The only new part here will be the `countChannel`. We construct a channel by using the built int function ` make(chan int)`. We have seen the `make` built in function before, the new thing here is the `chan` keyword which creates a channel and has to be followed by the type of the channel. We'll see more sophisticated examples later. Both `countUp` and `printOutput` functions take the channel as an argument. The first channel sends the current countUp value to the channel using `channelName <- value` and the latter receives the value using `value <- channelName` in an infinite loop.

```go
func main() {
	countChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printOutput(countChannel)
	}()

	go func() {
		defer wg.Done()
		countUp(countChannel)
	}()

	wg.Wait()
}

func countUp(counts chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		counts <- i
	}
}

func printOutput(counts chan int) {
	for {
		i := <-counts
		fmt.Printf("%d ", i)
	}
}
```

Executing the above code will print the numbers as expected but will panic at the end with `all goroutines are asleep - deadlock!`. `printOutput` is still waiting for a value from the channel, but nothing is being sent. It is a deadlock indeed!

Go has a mechanism to handle this issue. The sender can close the channel and thus signal the receiver that it no longer has any more information to send. We will use the built in `close` function by deferring that call before we are done waiting.

```go
go func() {
	defer close(countChannel)
	defer wg.Done()
	countUp(countChannel)
}()
```
 The receiver can use an optional second boolean result value to know whether the channel was closed or not `i, ok := <- counts` By convention, we will assign the value and check if the channel is closed in the same step `if i, ok := <- counts; ok`. All what we need to do is to close the channel.

```go
func printOutput(counts chan int) {
	for {
		if i, ok := <-counts; ok {
			fmt.Printf("%d ", i)
			continue
		}
		return
	}
}
```

Turns out that Go has a more concise way to listen for a channel until it is closed using our friend `range` which we used before to iterate over a collection.

```go
func printOutput(counts chan int) {
	for i := range counts{
		fmt.Printf("%d ", i)
	}
}
```

Now, that we have seen the code above, it is simple to see that the same could have been done by passing one function to the other. With no need for using unbuffered channel. The code would have been easier to reason about. Hence, you need to justify the use of channels fo synchronization as mentioned in the beginning.

#### Notes
- Sending a message to a closed channel will panic
- Closing an already closed channel will panic

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
8. [Pimp My Library ~ Martin Odersky](https://www.artima.com/weblogs/viewpost.jsp?thread=179766)
