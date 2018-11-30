## The Empty Interface
```
interface{} says nothing.
```
from [Go Proverbs](https://go-proverbs.github.io/)

It's the base object `object` in Java. It's either any object or `nil` pointer. It carries no semantic meaning at all.

It's really not the good way to write generic code. You won't write generic code in Java with `object`, will you?

Because `interface{}` has 0 method and any type in `Go` has at least 0 method, any type in `Go` satisfy empty `interface{}` as a consequence.

Therefore, do not use `interface{}` unless you have no alternatives. Everytime you want to use `interface{}`, think really hard. A good example is the [`json` pacakge in std lib](https://golang.org/pkg/encoding/json/#Marshal).
```go
func Marshal(v interface{}) ([]byte, error)
```
