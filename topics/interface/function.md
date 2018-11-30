## One Method Interface vs Function Type
We know that in Go, we can define function signature as type.
```go
package op
type Adder func(x int, y int) int
```
Then any function with the same signature could be considered as an instance of this type.

But, in order to do this, you have to convert your function to this type as shown below
```go
func Calculator(adder op.Adder) {...}

func Add(x, y int) int {
  return x + y
}
func main() {
  Calculator(op.Adder(Add))
}
```
One potential problem is that we now have source level coupling. You need the namespace accessing a type to do the type conversion.

The alternative would be
```go
package op
type Adder interface {
  Add(x, y int) int
}
```
and
```go
import op

type Add func(x, y int) int

func (a Add) Add(x, y int) int {
  return a(x, y)
}

func Calculator(adder op.Adder) {...}

func main() {
  Calculator(Add)
}
```

A good example is in the std lib [`net/http`](https://github.com/golang/go/blob/2012227b01020eb505cf1dbe719b1fa74ed8c5f4/src/net/http/server.go#L1987-L1996)
```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```
