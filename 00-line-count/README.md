## Run

Basic usage:

```
➜  00-line-count git:(master) ✗ go run main.go -path ./main.go
==========================
Count:  87
==========================
```

You can also ignore empty lines by passing the `-ignore-empty-lines` command line boolean flag:

```
➜  00-line-count git:(master) ✗ go run main.go -path ./main.go -ignore-empty-lines
==========================
Count:  77
==========================
```

## Command line arguments

[Go by Example: Command-Line Arguments](https://gobyexample.com/command-line-arguments)

 - import `os` package
 - `os.Args` will get us access to command line args. These will be raw command-line arguments which will include path to the program as the first value. `os.Args[1:]` will give us arguments to the program.

[Go by Example: Command-Line Flags](https://gobyexample.com/command-line-flags)

 - Go provides a built-in package for command-line flag parsing
 - The flags can be defined using the `Var` variants on the `flag` package and default values must be provided.
 - Once all defined, we need to parse the flags by calling `flag.Parse` which exists if the args are not specified as expected in a faulted way.
 - One can validate afterwards and can call `flag.Usage()` to print to usage if the validation fails (maybe also exit with `os.Exit` afterwards).

 [How to get the current working directory in golang](https://gist.github.com/arxdsilva/4f73d6b89c9eac93d4ac887521121120)

 - `os.Getwd()` gets the current directory,
 - It returns two values: directory as `string` and error as `error`. I am not sure what would be the case where this cannot return a value. This might give an idea when: https://github.com/golang/go/blob/f70bd914353b2331a48eedb84aceb458982eaac0/src/os/getwd.go#L26

[How can I open files using relative paths in Go?](https://stackoverflow.com/q/17071286/463785)

[How to check whether a file or directory exists?](https://stackoverflow.com/a/10510783/463785)

[Golang Determining whether *File points to file or directory](https://stackoverflow.com/a/25567952/463785)

[Go by Example: Reading Files](https://gobyexample.com/reading-files)

[reading file line by line in go](https://stackoverflow.com/a/16615559/463785)

## Side Learnings

 - It is possible to slice the arrays by specifying a half-open range with two indices separated by a colon. based on `b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}`, `b[1:4] == []byte{'o', 'l', 'a'}` will be true, sharing the same storage as b. More info: https://blog.golang.org/go-slices-usage-and-internals
 - In Go, the reference types need to be dereferenced. `*` in front of the variable name allows to dereference the pointers. More info: https://www.golang-book.com/books/intro/8
    - OK, not reference types maybe. Go apperently doesn't have a concept of a reference type, see https://www.tapirgames.com/blog/golang-has-no-reference-values and http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
 - In order to exit from a command line app, you use `os.Exit(int32)`. More info: https://stackoverflow.com/a/18969976/463785
 - When multi-value return is the case from a function call, all values needs to be referenced (if that's the right word to use here). Combining this with golang's compiler error when a declared variable is unused, this provides us a safe ability evaluate all the options before proceeding with the happy path. However, it turn this creates a very messy code where you have to probably write a same level of handling for each case. This seems to be being addressed by [Go 2 error handling proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md).
 - `fmt.PrintLn` doesn't accept formating directives. VS Code started giving the following warning: `Println call has possible formatting directive %s`. However, it prints the things out when you pass them into the call. For example, `fmt.Println("path does not exist", absoluteFilePath)` prints `path does not exist /Users/tugberk/apps/sink-or-swim/00-line-count/main`.
 - `for` look can also take a `bool` value which would act as a `while` loop in C#
 - used `defer` for the first time. [More formar definition](https://tour.golang.org/flowcontrol/12): "A defer statement defers the execution of a function until the surrounding function returns.". When used as an operator to a function statement, it defers the execution till the current function scope exists. The below code, for example, prints the numbers in the following order `1, 3, 2`:
   ```go
      func () {
         fmt.Println("1")
         defer fmt.Println("2")
         fmt.Println("3")
      }()
   ```
