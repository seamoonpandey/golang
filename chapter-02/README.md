# Working with Go modules

- Actually creating one of em and learning to integrate those things

## Setting up the task

- Create a directory hello/ and greetings/ at the same level

```zsh
mkdir hello greetings
cd hello
go mod init seamoonpandey.com/hello
cd ..
cd greetings
go mod init seamoonpandey.com/greetings
```

- In greetings.go we use the package greetings as it is a foreign package/module we're building. _which mean the go run . won't execute this file directly_

- In hello.go we import the seamoonpandey.com/greetings and invoke the function _but it won't work rn_

- In the same hello directory which we assume to be our main package _we specify the `package main` at_

```zsh
go mod edit -replace seamoonpandey.com/greetings=../greetings
```

- Make it clean and aquirable don't forget to

```zsh
go mod tidy
```

- Run

```zsh
go run .
```

## List of things I learned

```go
message := fmt.Sprintf("Hi, %v. Welcome", name)
```

In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:

```go
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

**_Note_**: Make sure you have one package main started files or the `go run .` won't do any good

## Return and handle error

Handling errors is an essential feature of solid code. In this section, you'll add a bit of code to return an error from the greetings module, then handle it in the caller.

- In greetings.go add some error handling since there's no sense sending a greeting back if you don't know who to greet. Return an error to the caller if the name is empty.

- Change the function so that it returns two values: a string and an error. Your caller will check the second value to see if an error occurred. (Any Go function can return multiple values)

- Import the Go standard library errors package so you can use its errors.New function.

- Add an if statement to check for an invalid request (an empty string where the name should be) and return an error if the request is invalid. The errors.New function returns an error with your message inside.

- Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.
