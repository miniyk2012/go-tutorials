## There are 3 important applications of this feature.
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

__2. Developers of different pacakges can work on the same interfaces without knowing each other__
Combining number 0 and 1, this is one of the most useful consequence of Go interface.

For example, let's say a networking package `A` defines a
```go
type Sender interface {
  Send() (done bool, err error)
}
```
Another network pacakge `B` may define the same `Sender` in their own source code. They don't need to mention each other at all. As long as their interfaces preserve similar semantics, as a user, we can define our own type implementing a `Sender` and be compatible with both libraries.

From the opposite perspective, if we are the definer of an interface, let' say we have a function
```go
type A interface {
  Method()
}

func F(a A) {}
```
We can use any third party libraries whose types implemnts `Method()`, these libraries' authors don't need to know us first. We also don't need to couple `F` with concrete types in their libraries.
