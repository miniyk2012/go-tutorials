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

Now, let's dive deep into it.

[The Basics](basics.md)  
[Syntax Sugar](syntax.md)  
[The Empty Interface `interface{}`](interface{}.md)  
 
### Other Read:
If you want to understand evaluation of interfaces, read [The Laws of Reflection](https://blog.golang.org/laws-of-reflection)

[How To Use Go Interfaces](https://blog.chewxy.com/2018/03/18/golang-interfaces/)
