Interface might be one of the most powerful and confusing feature in Go.

First, we know that interfaces are implicitly satisfied.
```go
type Something interface {
  Do()
}
```
As long as a type implements this method, it is of `Something` type.
```go
type Int int

func (i Int) Do() {}
```
You don't need to declare `type Int implements Something` like the way you will do in Java. This makes interface in Go less verbose and gives a dynamic language feeling. But this is not the key point.

I believe there are 3 important applications of this feature.

__0. Interfaces with same methods are the same interface.__
```
type A interface {
  Method()
}
type B interface {
  Method()
}
```
`A` and `B` are the same interface, any type which implements one implements the other.

In a language like Java or with similar OO behvarior, `type X implements A` and `type Y implements B` are implementing 2 different interfaces, which means if we have a function
```
func F(arg A)
```
`F` can only accept the instance of `X` and can not accept `Y` as arguments. Then, the programmer will write a wrapper type
```
type Y2 inherites Y implements A
```
which is purely for the sake of pleasing the compiler.

In Go, you just do `F(instanceOfY)`. You don't waste your time writing wrapper types.

__1. Developers can define pacakge private interface and function local interface__
#### Package private interface
```go
package a

type unexportedInterface interface {
  Method()
}
```
Now in another package
```
package b

type B struct{}

func (b B) Method() {}
```
Then you can pass any instance of B to functions which accept `unexportedInterface` as an argument.

This is very powerful, combining number 0, because it enables you to extend the functionality of your program without source level coupling. We will see an example later.

#### Function local interface
```go
func F(...) {
  ...
  type X interface {
    Method()
  }
  ...
}
```
Here we defined an interface locally in the scope of this function. You might think this is absolutely crazy. But, let look at an example from the famous [pkg/errors](https://github.com/pkg/errors/blob/2233dee583dcf88f3c8b22cb7a33f05a499800d8/errors.go#L269-L282)
```go
// pkg/errors/errors.go
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
```
How can we make of this function local interface? Let's define an error type
```go
package my

type Error struct{}

func (e *Error) Error() string { return "" }
func (e *Error) Cause() error { return nil }
```
Now, you can do
```go
import "pkg/errors"
import "my"

errors.Cause(&my.Error{}) // returns nil
```
Well, it's up to you of defining the semantics of a causation of an error. What I want to show you is that, you can write code which utilize private interfaces of other packages and extend their capability without having to couple your code with their source code. 

We haven't mentioned `pkg/errors` in our `my` package at all. This is powerful.
