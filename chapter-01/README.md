# First ever program

## Getting started

1 Create a hello directory for your first Go source code

```zsh
cd <path/you-want/golang>
mkdir hello
cd hello
```

2 Enable dependency tracking for your code.

When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.

To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path.

In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be github.com/mymodule. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module.

```zsh
go mod init seamoonpandey/hello
```

3 Setting basic program

- Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
- Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
- Implement a main function to print a message to the console. A main function executes by default when you run the main package.

4 Running

```zsh
go run .
```

## Updating with the quote library _(Learning importing modules)_

- In pkg.dev.go search for the quote, then at the top of this page, note that package quote is included in the rsc.io/quote module.

- In your Go code, import the rsc.io/quote package and add a call to its Go function

- Add new module requirements and sums

```zsh
go mod tidy
```

- Run the code

```zsh
go run .
```
