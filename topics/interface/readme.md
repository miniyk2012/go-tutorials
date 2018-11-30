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

I believe there are 2 most applications of this feature.

__1. Developers can define pacakge private interface and function local interface__
```go
package a

type unexportedInterface {
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
