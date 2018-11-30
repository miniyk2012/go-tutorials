## Several Syntax Sugar
1. Composition
```go
type X interface {
  M1()
}

type Y interface {
  X
  M2()
}
```
`Y` has all methods of `X`, there is no relationship of inheritance or composition in an inclusion sense. You just write less words.

Examples from std libs: https://golang.org/pkg/io/#ReadCloser
```go
type ReadCloser interface {
        Reader
        Closer
}
```


2. Overlap Doesn't Matter
```go
type X interface {
  M1()
  M2()
}

type Y interface {
  X
  M2()
  M3()
}
```
