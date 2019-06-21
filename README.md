# Introduction To Go for Non-Go Developers

_DinosaurJS 2019_

## Plan Of Attack

1. Background: Why Is Go A Thing?
2. Installation/Environment Setup
3. Basic Go Syntax and Structures
4. Error Handling
5. Concurrency with Goroutines and Channels
6. Build A Baby App

## Background: Why is Go A Thing?

Straight from [the docs](https://golang.org/doc/faq#creating_a_new_language), Go was created to address performance and scalability issues at Google.

The options were either speedy compiling, ease of programming, or speedy execution, without the option to have all three.

Go fixes these problems by offering a strictly typed language that is pleasant to write in, with a super fast build time.

To do this, Go employs a lightweight but thorough type system, concurrency for handling multiple threads of work at the same time, automatic garbage collection, and strict dependency specs.

The syntax of Go comes mainly from the C family, but works reduce the amount of code on the page and complexity around that code.

Obvious differences, coming from a language like JavaScript, include the following:

- Go uses types. Vanilla JS and ES6+ do not.
- Go does not have implicit error handling. If your program errors out, it will do so silently unless you explicitly tell it to behave differently. You won't get the console shouting at you if you don't tell it to.
- Go uses goroutines to allow multi-threading, vs the single thread event call stack from Node/JS.
- Go has opinions. There are rules that are strictly enforced and the compiler won't allow for deviations, vs JavaScript's flexible "whatever makes you happy" approach.

## Installation & Setup

Follow the instructions to [install go](https://golang.org/dl/) based on your system requirements.

You can also use an only sandbox like [Repl.it](https://repl.it/languages/go), or the [Go Playground](https://play.golang.org/).

Once you have Go installed and your GOPATH set up, create a `src` directory within your `go` directory.

You can find the base repository we will be starting with [here](https://github.com/martensonbj/intro-to-go).

## Basic Syntax And Structures

### Anatomy Of A Go File

Let's take a high level look at the structure of a basic go file.

```go
package main // 1

import "fmt" // 2

func main() { //3
   fmt.Println("Talk to me Goose") // 4
} // 5
```

1. Package Declaration

- Packages are go's way of organizing and reusing code within an application
- These packages, like classes, expose different variables that can be used within a file
- Every go program must have at least a `main` package
- When you run `go install` Go creates a binary file which will call the `main()` function of your program

2. Imported packages and libraries

- Exposes code from built in go libraries, third party packages, or internally created packages your program needs
- In this case, the package `fmt` comes with go and provides formatting methods for standard input and output
- Also, it's pronounced "fumpt". 🙄

3. The `main()` go function signature

- Starts with a keyword `func`
- The name of the function - `main`
- A list of optional parameters - not present
- An optional return statement - not present
- Opening curly brace

4. Code to be executed

- Here we're using the `fmt` package and calling the method `Println`
- This is similar to calling `console.log()` in JavaScript
- Notice that Go uses double quotes

5. Closing curly brace

- Indicates the end of a function
- Is often preceded by a `return` statement, similar to JS

> **_Try It_**
> Throw the above baby function into the go playground and run the program.

### Documentation

To access the help docs in your terminal, run `go help`.

Here, you can find a list of available commands to manage your go source code.

To check the docs on a particular package or symbol run `go doc <pkg> <method>`.

```go
go doc fmt
go doc fmt Println
```

### Types

Go is a _statically typed_ language. Unlike JavaScript, variables must have a
specified type and cannot change once they have been declared. If there is a
discrepancy in types, your program won't compile and you'll get an error.

Let's talk about Go's built in types.

#### Integer: `1`, `2`, `44`

- Example: `var age int = 21`
- `int`,`uint8`,`uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64`
- The `u` indicates an "unsigned" integer, which can only contain positive numbers or zero
- `byte`: alias for `uint8`
- `rune`: alias for `int32`

#### Float: `1.5`, `3.14`, `2100`

- Example: `var distanceInMiles float = 22.7`
- Floats can contain a decimal point
- `float32`, `float64`
- To see a visual of the difference between how Go stores data in these two types, visit [this playground](https://play.golang.org/p/ZqzdCZLfvC)

#### String: `"The need for speed"`

- Example: `var callSign string = "Maverick"`
- Note that Go uses double quotes to indicate a string

#### Booleans: `true`, `false`

- Example: `var isLessThan bool = 1 < 5`
- Can be generated using the expected logical and comparison operators
- `&&`, `||`, `!`, `<, <=, >, >=, ==, !=`

#### Error: `error`

- Example: `error.Error()` ==> string
- `log.Fatal(err)` will print the error message and stop execution

#### TypeOf

- To check the type of a variable, you can use the `TypeOf()` method from the package `reflect`.

> **_TRY IT_**
> Add `import reflect` under the `import fmt` statement, and then add `fmt.Println(reflect.TypeOf("hello"))`

- You can also use string formatting within `fmt` to print out a type:

```go
	var f float64 = 52.2
	fmt.Printf("f has value %v and type %T", f, f)
```

#### Type Conversion

Sometimes you'll want to convert one type to another. Go lets you do this using the type name as a function.

```go
var age int = 21
var floatAge = float64(age)
fmt.Println(reflect.TypeOf(age))
fmt.Println(reflect.TypeOf(floatAge))
```

### Variables

Variables in Go are written in camel case, similar to JavaScript, and can be defined a few different ways.

**Option 1:**

- Initialize a variable with a type and the keyword `var`
- Looks similar to JavaScript with an additional type declaration

```go
var pilot string = "Iceman"
```

**Option 2:**

- Go can infer the type of initialized variables and the type declaration can be
  omitted.

```go
var pilot = "Iceman"
```

**Option 3:**

- Feels similar to using `let` in JavaScript to instantiate and empty variable, but instead of storing the value as `undefined`, Go will default to the zero value for that type
- string: `""`
- int: `0`
- float: `0`
- bool: `false`

```go
var pilot string
// This will have a default value of "" that can be modified later
```

**Option 4:**

- Declare multiple variables at once

```go
var pilot1, pilot2 string = "Goose", "Maverick"
```

**Option 5:**

- Can only be used within a function
- Most commonly used pattern
- To declare and assign a default value at the same time
- You can omit the keyword `var`, and go will infer the value

```go
pilot := "Iceman"
```

#### Var vs Const

Consts, like in JavaSript, are variables who's values cannot be changed.
Attempting to modify a const will result in a run-time `panic`.

**Sidebar**:

- Panic vs Error:

  - There are two main types of errors in Go. Panic, and Error.

    Panic is used when the error is fatal to your program and there is nothing else that can be done to move forward.

  - Errors indicate that something bad happened, but it might be possible to continue running the program. More on this later.

- RunTime vs CompileTime
  - Reminder that compile time errors are events like syntax or typechecking errors. They occur when you as the developer are compiling your code.
  - Run time errors happen after the program is compiled and a user or browser is trying to execute the code.
  - These events are things like trying to open a file or url that doesn't exist, running out of memory, trying to do something that syntactically is legit but isn't valid in real life (like dividing by 0).

### Data Structures

1. For Loops
2. Ifs & Switches
3. Arrays
4. Slices
5. Maps

#### For Loops

The only loop in Go is the `for` loop. It looks nearly identical to our friendly for loop in JavaScript.

Like variables, we'll start with the most verbose version and work our way to idiomatic go.

```go
func main() {
  i := 1
  for i <= 100 {
    fmt.Println(i)
    i += 1
  }
}
```

Like in JS, we instantiate our counter `i` at 1, and then for every iteration as long as our counter `i` is below 100, we print its value and then add one.

A shorter way to write this that looks more JS familiar is:

```go
  func main() {
    for i := 1; i <= 100; i++ {
    fmt.Println(i)
    }
  }
```

#### If & Switch Statements

An example if statement:

```go
  if i > 100 {
    fmt.Println("Greater than 100")
  } else if i == 100 {
    fmt.Println("Equals 100")
  } else {
    fmt.Println("Less than 100")
  }
```

An example switch statement:

```go
  switch {
  case i < 10: fmt.Println("Less than 10")
  case i < 20: fmt.Println("Less than 20")
  case i < 30: fmt.Println("Less than 30")
  default: fmt.Println("WTF")
  }
```

#### Arrays

Arrays in Go have some significant differences compared to those in JavaScript.

Let's compare notes.

**JavaScript:**

```javascript
const grabBag = [];
// or
const grabBag = ["banana", 45, true, "2"];
```

**Go:**

```go
var scores [5]float64
```

In JS, defining an array is pretty chill.

You can make it whatever length you want, you can modify that length, and you can throw any combination of data
types into it.

Go, however, has _opinions_. You must define a fixed length of elements, all of which must be of the same type. In the above example, our variable `scores` is an array of 5 elements, each of which is a float.

> **_TRY IT_**
> In the Go playground, print out the variable `scores` as is. What do you see?

To add elements into this array, both JavaScript and Go allow you to insert elements into a specific index location.

```javascript
grabBag[0] = "hello"; // => ["hello"]
```

```go
scores[0] = 5 // => [5, 0, 0, 0, 0]
```

To add multiple elements, you could start by doing something like this:

```go
  var scores [5]float64
  scores[0] = 9
  scores[1] = 1.5
  scores[2] = 4.5
  scores[3] = 7
  scores[4] = 8
```

Or, to avoid setting each index individually:

```go
scores := [5]float64{9, 1.5, 4.5, 7, 8}
```

> _TRY IT_
> Using the array above and a for loop, print out the average score within the array.
> Hint: To find the length of an array, use `len(array)`.

What unexpected errors did you run into? How did you fix them?

When looping over arrays, we can implement a different version of the basic for loop using the `range` keyword:

```go
  var total float64
  for i, value := range scores {
    total += value
  }
```

Let's break it down:

- `i` Still represents the index position we are pointing to in our array
- `value` represents the value of the element at that index (aka `scores[i]`)
- `range` is a keyword that is followed by the variable representing the array we are looping over

> _TRY IT_
> What happens when you replace the for loop from finding the average score with
> this one?

As we've seen, Go is very particular about types and making sure any declared variable is used within a program. In this case, we aren't using the `i` variable representing the index of the array and we get this error:

```bash
i declared and not used
```

Go uses a single underscore, `_`, to replace an unused variable within a program. Make this modification and watch the program pass.

Defining an array with a specific length can cause some obvious problems.

Adding or removing an element would require redefining the entire variable signature to specify a different length, which often is unknown.

This brings us to Slices.

#### Slice

A slice is a segment of an array. You still need a single type, access elements within it by their index, and a slice has a length, but unlike arrays their lengths can change.

Instantiating a slice looks almost identical to instantiating an array, just without the specific length.

```go
  var scores[]float64
```

Because slices are segments of an array, they must always be associated with an underlying array.

To create an empty slice, you must use the `make` function with takes two arguments - what you are instantiating, and a length.

Alternatively, if you know what will be in your slice, you can instantiate it using the curly braces.

```go
  scores := make([]float64, 5)
  // OR
  scores := []float64{1, 2, 3}
```

Here, we are associating our slice called `scores` with an array containing 5 elements of type `float64`.

Note that slices can never be longer than their associated array, but they can
be smaller.

We can add a third argument to `make()` to tell go how much of an underlying
array to take up, and what length the underlying array should be:

```go
 scores := make([]float64, 5, 10)
```

This would create a slice of 5 elements of type `float64`, associated with an
underlying array that has a length of 10.

##### Slice Helper Methods

Append:

- `append(originalSlice, newEl1, newEl2)`

```go
  slice1 := []int{1, 2, 3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)
```

Copy:

- `copy(destination, source)`

```go
    destination := make([]int, 2)
    source := []int{1, 2, 3}

    fmt.Println("before copy")
    fmt.Println(destination, source)

    copy(destination, source)

    fmt.Println("after copy")
    fmt.Println(destination, destination)
```

> _TRY IT_
> Run the example code from Copy block.
> What do you notice about the resulting data object?

#### Maps

Maps look and behave similarly to Objects in JavaScript, using a set of key
value pairs.

A map starts with the keyword `map`, followed by the key type, and the
value type.

```go
 var pilotsAges map[string]int
```

Here we are telling Go to create a map called `pilotsAges` which will have keys as
strings, and values as integers.

Let's build one with data in it:

```go
  var pilotsAges map[string]int
  pilotsAges["maverick"] = 30
  pilotsAges["goose"] = 27
  fmt.Println(pilotsAges)
```

> _TRY IT_
> Throw that code into the Go playground. What happens? What are we missing?

If we run the above code as-is, we'll see the following error:

```go
panic: assignment to entry in nil map
```

The (run-time!) error here, indicated by the keyword `panic`, is telling us that
we're trying to assign entries to a map that doesn't exist yet.

Look at where we are declaring our `pilotsAges` variable. Here, we are telling
go to create a variable called `pilotsAges` that will have a type of map with
particular key value pair types. We haven't actually set that variable to a
value of any kind.

Remember the `make` function from our slice examples? Let's do that now.

```go
  var pilotsAges map[string]int
  pilotsAges = make(map[string]int)
  // 	var pilotsAges = map[string]int{}

  pilotsAges["maverick"] = 30
  pilotsAges["goose"] = 27
  fmt.Println(pilotsAges)
```

Or, with the shorthand syntax:

```go
  pilotsAges := make(map[string]int)
  // pilotsAges := map[string]int{}
  pilotsAges["maverick"] = 30
  pilotsAges["goose"] = 27
  fmt.Println(pilotsAges)
```

And finally, creating the map in place:

```go
  pilotsAges := map[string]int{
    "maverick": 30,
    "goose": 27,
  }
  fmt.Println(pilotsAges)
```

Now we should see:
`map[goose:27 maverick:30]`

To ask for a specific value from a map, you'd use syntax similar to interacting
with objects in JS:

```go
pilotsAges["maverick"] // ==> 30
```

> _TRY IT_
> What happens if you ask for a key that doesn't exist?

Go has a clever way to handle checking for null values. Looking up a value
within a map actually returns two elements: the value (if it exists), and a
boolean for whether or not that value existed.

```go
  icemansAge, ok := pilotsAges["iceman"]
  fmt.Println(icemansAge, ok)
```

This makes it easy to add logic around checking for null values within our
program using an if statement:

```go
  if icemansAge, ok :=  pilotsAges["iceman"]; ok {
    fmt.Println(icemansAge)
  } else {
    fmt.Println("I don't know what you want from me")
  }
```

First, we create the two variables `icemansAge` and `ok` inside our if
statement and set them to the two elements we get back from a lookup in our map.

Then, we check what boolean we got back in our `ok` variable. The code within
the first block will fire only if `ok` returns true, otherwise our if block will
continue to the else block.

As with arrays, we can simplify creating our map using abbreviated syntax:

```go
pilotsAges := map[string]int{
  "maverick": 30,
  "goose": 27,
}
```

#### Functions

Functions in Go, like in most languages, map a specified set of inputs
to a specified output.

In Go specifically, the input parameters also require a type definition, and there is an optional return type definition.

```go
func printAge(age int) int {
  fmt.Println(age)
  return age
}
```

A function can also return multiple values, which both need to be specified in
the return type definition:

```go
  func myFunc() (int, int) {
    return 5, 6
  }

  func main() {
    x, y := myFunc()
  }
```

> _TRY IT_
> Create a function called `average` that takes an array of floats.
> The function should return the average of the provided array.
> Make sure to call this function from within `main()`

Similar to the spread/rest operator in JavaScript, you can pass an unspecified number
of arguments to a function.

```go
  func doThings(args ...int) int {
    total := 0
    for _, num := range args {
      total += num
    }
    return total
  }

  func main() {
    fmt.Println(doThings(1, 2, 3))
  }
```

##### A Note About Scope

Functions do not have access to variables that are not global, or defined within
themselves.

In the variadic example above, our `main()` function would not have access to
the variable `total`.

#### Defer/Panic/Recover

Go has a few built in helper functions that help guide a go program to behave a
certain way:

- `defer`: executes a line of code last within a function
- `panic`: called during a run time error, halts execution of the program
- `recover`: tells go what to do after a panic

**Defer**

Waits until everything else within a function is complete before firing.

Multiple defers are executed in a `LIFO` order.

```go
func doThings() {
  defer fmt.Println("Do this last!")
  defer fmt.Println("Do this second to last!")
  fmt.Println("Things And Stuff")
}

func main() {
  doThings()
}
```

An example use case would be when opening and closing a file.

```go
f, _ := os.open(filename)
defer f.Close()

// a bunch of other code you want to run once you've opened the file
```

This is helpful because it keeps two functions that are closely related physically close to each other for readability.

Deferred functions are also run even if a runtime panic occurs.

**Panic/Recover**

Panic will be called during a run time error and fatally kill execution of a
program.
Recover will tell Go what to do when that happens, returning what was passed to `panic`.

Note that in order for recover to do its job it must be paired with `defer`, which will fire even after a `panic`, otherwise `panic` completely shuts down the execution of a program.

```go
func doThings() {
  for i := 0; i< 3; i++  {
    fmt.Println(i)
    if i ==2 {
      panic("PANIC!")
    }
  }
}

func main() {
  doThings()
}
```

With the above code, once we hit the `panic()` function, our program will stop
executing. Adding a `recover()` cleanup function will tell our program what to do
when this happens.

```go
func handlePanic() {
  // recover() will only return a value if there has been a panic
  if r := recover(); r != nil {
    fmt.Println("We panicked but everything is fine. Panic message received:", r)
  }
}

func doThings() {
  defer handlePanic()
  for i := 0; i< 3; i++  {
    fmt.Println(i)
    if i ==2 {
      panic("PANIC!")
    }
  }
}

func main() {
  doThings()
}
```

> _TRY IT_
> Add a line that tries to print something to the console after the panic
> function.
> Comment out the defer function
> What happens?

#### Pointers

When we pass an argument to a function in Go without any additional symbols, we
are passing the function a **copy** of the original variable.

This means that modifying the argument will not modify the original variable in
memory.

> _TRY IT_
> Paste the following code snippet into the console. Before running it, what do
> you expect to see as an output?

```go
  func makeTopGun(name string) {
    fmt.Println("in makeTopGun")
    name = "Maverick"
  }

  func main() {
    name := "Iceman"
    makeTopGun(name)
    fmt.Println(name)
  }
```

As you noticed, the variable `name` within the `main()` function is not modified.
It still prints out `Iceman` even though our `makeTopGun()` function is reassigning
that variable to be `Maverick`.

Check out what the `name` argument prints to the console within the `makeTopGun`
function. You should see the string `Iceman`, which is a COPY of the value of the
variable `name.`

Sometimes the intended behavior is not to modify the original variable - but what if we DO want to modify that variable?

This is when we would want to use something called a `Pointer`, which is
indicated by adding a `*` in front of the type definition in the argument, and tells Go to
expect the actual data stored in memory.

```go
  func makeTopGun(name *string) {
    fmt.Println("in makeTopGun")
    *name = "Maverick"
  }

  func main() {
    name := "Iceman"
    makeTopGun(&name)
    fmt.Println(name)
  }
```

You'll notice that within the `main()` function we also added a `&`
character in front of the `name` variable as we pass it to our `makeTopGun`
function.

When we run it this time, we see that the variable has been permanently modified
and now prints out `Maverick`. WTF.

#### \* and &

So what's actually happening here?

In Go, a function indicates it's expecting a `pointer` by placing an asterisk `*` in front of the type definition within the argument parentheses.

```go
func myFunc(arg *string)
```

Here, our function is expecting `arg` to be a POINTER, pointing to the actual
location in memory that stores this variable.

Similarly (but different), an asterisk `*` is ALSO used to _dereference_ a pointer variable.

Within the `makeTopGun()` function above, we wrote `*name = "Maverick"`. This line is saying: store the string "Maverick" in the _memory location_ we just passed to our function.

> _TRY IT_
> In the `makeTopGun` function, print out the argument `name` when using a
> pointer.
> Try removing the `*` from the last line in `makeTopGun()` to turn it back into a non dereferenced variable

This is because _Go has rules_, and it is expecting the incoming string to be a
pointer value and will ardently refuse to work with anything else.

Lastly - the `&` tells Go to find the memory location of that variable instead
of making a copy of it.

Once go finds that memory location, it will return the pointer version of that variable and pass that along, allowing us to modify the original variable.

#### Structs and Interfaces

Go provides data structures that allow you as a developer to define your own
types for commonly used objects and sets of methods.

A `struct` is a collection of fields that have defined types and is reusable
across your program.

```go
type Pilot struct {
  firstName string
  lastName string
  callsign string
  aircraft string
}
```

Another way to write this is:

```go
type Pilot struct {
  firstName, lastName, callsign, aircraft string
}
```

We can now treat the `Pilot` struct as a type of variable and instantiate it
with or without default values.

```go
var p Pilot
p.firstName = "Pete"
p.lastName = "Mitchell"
p.callsign = "Maverick"
p.aircraft = "f14"
```

Or in shorthand

```go
p := Pilot{firstName: "Pete", lastName: "Mitchell", callsign: "Maverick",
aircraft: "f14"
```

Then, in order to access the fields, we use dot notation similar to working with
Objects in JS.

```go
p := Pilot{firstName: "Pete", lastName: "Mitchell", callsign: "Maverick",
fmt.Println(p.firstName) // => "Pete"
```

To use a struct within a function, we simply add the struct name within the
function signature as the type of argument we are passing in.

```go
func introducePilot(p Pilot) string {
	intro := fmt.Sprintf("Introducing %s %s, callsign %s, flying the %s", p.firstName, p.lastName, p.callsign, p.aircraft)
	return intro
}
```

In `main()`, add a few lines to call this function.

```go
func main() {
	p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
  pilotIntro := introducePilot(p)
  fmt.Println(pilotIntro)
}
```

Let's create another struct for a squadron of pilots

```go
type Squadron struct {
	name           string
	numberOfPilots int
	currentTopGun  string
}
```

Next, create a function that introduces the squadron:

```go
func introduceSquadron(s Squadron) string {
	intro := fmt.Sprintf("Introducing squadron %s, with %d pilots. Current top gun: %s", s.name, s.numberOfPilots, s.currentTopGun)
	return intro
}
```

In main we can call create instances of our structs and call each function:

```go
func main() {
	p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
	s := Squadron{name: "Top Gun", numberOfPilots: 10, currentTopGun: "Iceman"}
  pilotIntro := introducePilot(p)
  squadronIntro := introduceSquadron(s)
  fmt.Println(pilotIntro)
  fmt.Println(squadronIntro)
}
```

What if in `introduceSquardon` we wanted to modify who was Top Gun?

> _TRY IT_
> In `introduceSquadron`, update `currentTopGun` to be "Maverick"
> Then, print out the squadron variable `s` at the end of `main()`. What do you
> notice?

Remember in Go that arguments are **COPIES** of the variables we pass in. In order to permanently modify the original variable, we need to use pointers.

Modify the arguments to use the `&` and `*` symbols that tell Go we want access
to the original variable in memory.

```go
func introduceSquadron( *Squadron) string {...

func introducePlayer(p *Pilot) string {...

// in main()
pilotIntro := introducePilot(&p)
squadronIntro := introduceSquadron(&s)
```

#### Methods

Right now we have two very similar functions that print out introduction
sentences about the structs we pass in.

We can reduce some repetition by implementing a `method` that we can call on our
structs, instead of an explicit function.

```go
 func (p *Pilot) introduce() string {
	p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
	return intro
 }
```

The main difference between the function signatures is that now, instead of
passing in our Pilot struct as a variable, we insert it as a `receiver` between the `func`
keyword and the name of our more generic function, `introduce`.

This tells Go that when the `introduce` function receives (is called on) a struct with the shape
Pilot, it should execute the code within the curly braces.

We call the method a little differently:

```go
// former syntax:
p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
pilotIntro := introducePilot(&p)

// syntax using a method:
p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
pilotIntro := p.introduce()
```

Note that we no longer need to include the `&` operator - Go automatically
passes a pointer to a method call.

> _TRY IT_
> Modify the `introduceSquadron` function to be a method called `introduce`
> Modify the main function to reflect those changes

The entire program at this point should now look like this:

```go
package main

import (
	"fmt"
)

type Pilot struct {
	firstName string
	lastName  string
	callsign  string
	aircraft  string
}

type Squadron struct {
	name           string
	numberOfPilots int
	currentTopGun  string
}

func (p *Pilot) introduce() string {
	intro := fmt.Sprintf("Introducing %s %s, callsign %s, flying the %s", p.firstName, p.lastName, p.callsign, p.aircraft)
	return intro
}

func (s *Squadron) introduce() string {
  s.currentTopGun = "Maverick"
	intro := fmt.Sprintf("Introducing squadron %s, with %d pilots. Current top gun: %s", s.name, s.numberOfPilots, s.currentTopGun)
	return intro
}

func main() {
	p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
	s := Squadron{name: "Top Gun", numberOfPilots: 10, currentTopGun: "Iceman"}
  pilotIntro := p.introduce()
  squadronIntro := s.introduce()
  fmt.Println(pilotIntro)
  fmt.Println(squadronIntro)
  fmt.Println(s)
}
```

#### Interfaces

If you think of structs as a set of properties within a type, interfaces can be
thought of as a set of methods that define a type.

```go
type TopGunEntity interface {
  introduce() string
}
```

Once again we start with the `type` keyword, followed by the name of the
interface, and then the actual keyword `interface`.

Within the curly braces we list a set of methods that are associated with our interface type and the type
those methods should return.

Our interface `TopGunEntity` includes any type that has a method named `introduce()`, which in our program includes `Pilot` and `Squadron`. Any type that defines this method is said to `satisfy` the `TopGunEntity` interface.

Next, create a function that can be called on any types that satisfy our
`TopGunEntity` interface.

```go
func introduceSomething(t TopGunEntity) {
  fmt.Println(t)
  fmt.Println(t.introduce())
}
```

Then, let's use our interface to introduce all of the entities in our program so far.

```go
func main() {
  p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
	s := Squadron{name: "Top Gun", numberOfPilots: 10, currentTopGun: "Iceman"}
  entities := []TopGunEntity{&p, &s}
  for _, entity := range entities {
    introduce(entity)
  }
}
```

You can also use the empty interface type to indicate that Go should accept
anything.

```go
interface{}
```

```go
func DoSomething(v interface{}) {
  // This function will take anything as a parameter.
}
```

Once again, our entire program should now look like this:

```go
package main

import (
	"fmt"
)

type Pilot struct {
	firstName string
	lastName  string
	callsign  string
	aircraft  string
}

type Squadron struct {
	name           string
	numberOfPilots int
	currentTopGun  string
}

type TopGunEntity interface {
  introduce() string
}

func introduceSomething(t TopGunEntity) {
  fmt.Println(t)
  fmt.Println(t.introduce())
}

func (p *Pilot) introduce() string {
	intro := fmt.Sprintf("Introducing %s %s, callsign %s, flying the %s", p.firstName, p.lastName, p.callsign, p.aircraft)
	return intro
}

func (s *Squadron) introduce() string {
  s.currentTopGun = "Maverick"
	intro := fmt.Sprintf("Introducing squadron %s, with %d pilots. Current top gun: %s", s.name, s.numberOfPilots, s.currentTopGun)
	return intro
}

func main() {
	p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
	s := Squadron{name: "Top Gun", numberOfPilots: 10, currentTopGun: "Iceman"}
    entities := []TopGunEntity{&p, &s}
  for _, entity := range entities {
    introduceSomething(entity)
  }
}
```

## Packages

Packages are directories with one or more Go source files that allow Go to reuse code across a program and across files.

Every go file must belong to a package.

So far, the packages we've seen are `main`, the package we wrote, `fmt` and `reflect`, which are built
into the Go source code. There are [tons](https://golang.org/pkg/) of packages
that come out of the box with Go.

To import packages, you list them at the top of your file either like this:

```go
import "fmt"
import "math"
import "reflect"
```

Or more commonly, like this:

```go
import (
  "fmt"
  "math"
  "reflect"
)
```

Some of the (many) available go packages are:

- `strings` - simple functions to manipulate strings

  - Example Methods: `Contains`, `Count`, `Index`, `HasPrefix`

- `io` - handles input/output methods related to the `os` package (like reading
  files)

  - Example Methods: `Copy`, `Reader`, `Writer`

- `os` - methods around operating system functionality

  - Example Methods: `Open`, `Rename`, `CreateFile`

- `testing` - Go's build in test suite

  - Example Methods: `Skip`, `Run`, `Error`

- `net/http` - provides http client and server implementations
  - Example Methods: `Get`, `Post`, `Handle`

### Exported vs Unexported Names

When you import a package, you can only access that package's exported names
(functions). Anything (variable, type, function) that starts with a capital
letter is exported and is visible outside the package.

Anything that starts with a lowercase letter is NOT exported, and is only
visible within the same package.

```go
fmt.Println()
```

Within the `fmt` package, the function `Println` is exported (starts with a
capital letter) and can be accessed within our `main` package.

### Custom Packages

Until now we've been writing all of our go code in the `main` package and only
using functionality from Go's imported libraries.

Let's create a folder called `utils`, and within that create a file called
`math.go`.

In math.go add the following:

```go
package utils

func Add(nums ...int) int {
  total := 0
  for _, v := range nums {
    total += v
  }
  return total
}
```

And then in our `main.go` file, lets import and use this package.

Note that the import name is a path that is relative to the `src` directory.
Often, this will be a github location as your files will live in a github
repository.

```go
package main

// Use whatever path gets you from your src directory to the utils directory
utils directory
import (
	"fmt"
	"intro-to-go/utils"
)

func main() {
	sum := utils.Add(1, 2, 3)
	fmt.Println(sum)
}
```

## Testing

Go includes a package, `testing`, that contains all the functions necessary to
run a test suite and command to run those tests, `go test`.

Test file names must end with `_test.go` and the Go compiler will know to ignore
these unless `go test` is run.

Let's create a test for our `Add` method in the `utils` directory.

`touch math_test.go`

```go
package utils

import "testing"

func TestAdd(t *testing.T) {
  total := Add(2, 3, 4)
  if total != 9 {
    t.Errorf("Sum was incorrect! Received: %d, Got: %d", total, 9)
  }
}
```

In tests, the test name should start with `Test` followed by the name of the
function you are testing. The only parameter passed into the test should be `t *testing.T`.

Navigate to the `utils` directory and run `go test` and watch it pass!

// To see a completed version of this section, checkout branch
`02-packages-and-testing`

## Concurrency

Goroutine: From [the docs](https://tour.golang.org/concurrency/1), a goroutine is a lightweight threat managed by the Go runtime.

A goroutine is indicated by adding the keyword `go` before the name of a
function. This will tell Go to fire up a new gorouting running that function on
a separate thread.

```go
  import (
    "time"
    "fmt"
  )

  func say(s string) {
    for i := 0; i < 3; i++ {
      fmt.Println(s)
      time.Sleep(time.Millisecond*100)
    }
  }

  func main() {
    go say("Hello")
    say("There")
  }
```

If all of the code within a function is a go routine:

```go
  func main() {
    go say("Hello")
    go say("There")
  }
```

Everything will be non blocking, and nothing will finish execution. In order to fix this, we need to synchronize our goroutines, using the package [sync](https://golang.org/pkg/sync/).

To start, create a variable that defines a `WaitGroup` - meaning a set of go routines you want execute to completion before moving forward in your program.

```go
package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func handlePanic() {
  if r := recover(); r != nil {
    fmt.Println("We panicked! But its fine. Message received: ", r)
  }
}

func printStuff(s string) {
  // Decrement the wait group counter
  // Use defer so that if the function panics we aren't waiting forever
  // Also figure out what to do if the program panics
	defer wg.Done()
  defer handlePanic()

  for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
  // Increment the wait group counter
	wg.Add(1)
  // Launch a goroutine
	go printStuff()
  // Increment the wait group counter
	wg.Add(1)

	go printStuff()
  // Wait for the waitgroup counter to be 0 before continuing
	wg.Wait()
}
```

Here, we're defining a variable that represents our WaitGroup. The `Add()`
function takes an integer and adds it to a behind the scenes counter. If the
counter becomes zero, all goroutines are released and the program will continue.

## Channels

Send and receive values between your goroutines.

Create a channel but using the `make()` function, and the `chan` keyword,
including the type declaration.

```go
myChannel := make(chan int)
```

Then, fire off a couple goroutines form which you are expecting values.

```go
package main

import "fmt"

func mathThings(c chan int, someValue int) {
	// From right to left, calculate the value, send it to the channel
	c <- someValue * 5
}

func main() {
// Instantiate a new channel
	resultChan := make(chan int)

	// Kick off a couple go routines
	// Notice that we don't need the WaitGroup here, because channels are blocking.
	// Our channel is expecting a value back so it will block until it finishes execution
	go mathThings(resultChan, 5)
	go mathThings(resultChan, 3)

	// Store whatever we get back from our channel, in a var
	num1 := <-resultChan
	num2 := <-resultChan
  // or num1, num2 := <-resultChan, <-resultChan
	fmt.Println(num1, num2)
}
```

We can shorten this up a bit by iterating over our channel, adding back in
WaitGroups, and including a buffer value when instantiating our channel:

```go
package main

import
  "fmt"
  )

var wg sync.WaitGroup

func mathThings(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	// second argument is the buffer to tell our channel how many times we plan on using it
	resultChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mathThings(resultChan, i)
	}

	// Wait for all goroutines to complete before closing the channel
	wg.Wait()

	// Close wont wait for the goroutines to be done without additional synchronization
	close(resultChan)


	for item := range resultChan {
		fmt.Println(item)
	}
}
```

## Resources

- [Offical Golang Docs](https://golang.org/doc/)
- [Golang FAQ](https://golang.org/doc/faq#methods_on_values_or_pointers)
- [How To Use Interfaces In
  Go](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
- [Introducing Go](http://shop.oreilly.com/product/0636920046516.do), Caleb
  Doxsey, O'Reilly Publications
- [Web Applications With
  Go](https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html)
- [Go Language Programming Practical Basics
  Tutorial](https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX)
- [Star Wars API](https://swapi.co/)
- Additional shout outs to my colleague Justin Holmes, and former colleagues
  Mike McCrary and Steven Bogacz for their patience with my endless questions.
