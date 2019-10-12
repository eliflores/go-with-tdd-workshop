# Notes

* [Unicode Consortium](http://unicode.org) 
* The syntax of Go is minimalistic. 
* It only has 1 construct for looping using the `for` and `range` **reserved words**.
    * 1st get the index and 2nd the thing at the index position.
* The compiler does not let me put variables that I do not use, so you can use `_` for the compiler to ignore.     
* The idea of TDD is to be humble. Start with _baby steps_.
* Train yourself to be humble, then you won't stumble as often. We are always on overdrive. 
* Do not fall into the habit of ignoring error messages.
* `Example` is a kind of test.
* `=` is the regular assignment operator. `:=` is used for type inference.
    * The type of the variable is inferred from what it is assigned to.
* `%FOO` is called _formatting verbs_.
    * Adding a `#` is a more strict (format) representation. 
* Use short assignments when you can, unless you need to do a more explicit declaration.
* Single quotes are a character, Gophers call them runes. e.g. `A` (code for `A` in the ASCII table)
* Arrays in Go are like arrays in Pascal, they are fixed.
* Everything in `go` has a zero value by default.
* `slices` are flexible arrays (closer to the idea of `lists` in python)
* `rune` is an alias of `int32`.
* A conversion: `rune(myInt)`
* `Example` is the most illustrative example / test of the package.
* The toolchain of Go is so good :)
* To install packages:

```bash
go get golang.org/x/text/unicode/runenames
go install golang.org/x/text/unicode/runenames
```

* We are going to need a data structure that has a character and a name.
* There is no inheritance in Go.
* No special assertion functions in Go.
* The name of the variable is in proportion to the scope
* The functions in `fmt` that start with `S` they build / generate strings and return it. The ones without `S`, they print to the output.
* A method is a function that is bound to a specific type.
* You should start using only the standard library if you can.
* `==` and `=!` only work when comparing slices to `nil`
* Assertions that work with complex data structures.
* Trailing comma at the end for slices.
* Testify assert package: Uses the `t` structure will be used internally to report errors.
* In C, arrays are syntactic sugar for pointers.
* Go has no generics.
* It's better that the tests are simple.
* Table driven tests, subtests. 
    * (Fixtures)
* The main reason for mocking is to enable TDD.
* You implement a test that exposes the problem.
    * We are trained as professionals to know how to do stuff.
        * You immediately start thinking about the how.
    * The _what_ is more important. 
        * If the _what_ is not clear the _how_ is probably going to be wrong.
* The main function in Go does not take arguments.
* No optional params in go, but we can use `...`
* `defer func() { os.Args = oldArgs }()` - it won't be called now.
    * This call is going to happen when the function finishes, even if the function terminates abruptly with a panic, in a reverse order of how they were called (to use for: close flies, I/O, etc)
    * Simpler and general.
`go build` -> takes the name of the directory as default.
* Go does not have true itrators. 
    * We have to fake it, there are to ways if doing that:
        * Implement a classic iterator pattern.
* Same of number of threads as the number of cores than the machine where the program is running has.
* Conditions in go:
```go
// Below, the v > 0 is the actual condition
if v := p.target[p.index]; v > 0 {
 /// ...
}
```
           
* Channel where my clients can only take stuff, not put stuff. 
    * It's like a queue, it's a safe-thread queue.
    * FIFO data structure
* The only place where I can put things on the channel, it's when I declare it. After, it's not possible.     
    * It's crucial that it is done inside of a go routine.
    * `ch <- v` - in this case the `<-` is the receive operator. 
        * This operation is blocking.
    * Channels with no buffers are useful to synchronize stuff.    
    * The send and the read operations are blocking.        
* Go's work-stealing scheduler (low stealing).
* The solution in the repo [here](https://github.com/standupdev/runes2019/blob/master/runes/main.go#L22) is using iterators.
           
 ## References     
* [Test Driven Development: By Example](https://www.goodreads.com/book/show/387190.Test_Driven_Development)
* [Growing Object-Oriented Software, Guided by Tests](https://www.goodreads.com/book/show/4268826-growing-object-oriented-software-guided-by-tests)
* [Mocks Aren't Stubs](https://martinfowler.com/articles/mocksArentStubs.html)
* [Awesome Go](https://github.com/avelino/awesome-go)
* [Go playground](https://play.golang.org/)








