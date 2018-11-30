# Interface
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

You can also define unexported method in an interface,
```go
type Something interface {
  do()
}
```
although this is not useful.
