# Working with Go modules

- Actually creating one of em and learning to integrate those things

## List of things I learned

```go
message := fmt.Sprintf("Hi, %v. Welcome", name)
```

In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:

```go
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```
