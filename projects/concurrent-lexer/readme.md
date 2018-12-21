# Concurrent Lexer
In this project, we will write a concurrent lexer/tokenizer for a simple arithmetic syntax
```
(* ( + 1 2) 3)
```
Tokens are defined as
```
(
)
+
-
*
/
integers
EOF
INVALID
```
The API should be
```go
var input string = readInput()
lexer := lexer.New(input)
token := lexer.NextToken() // token is either a valid token or the special EOF or INVALID
```
