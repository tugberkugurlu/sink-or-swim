## Run

```bash
go run main.go
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

## Side Learnings

 - It is possible to slice the arrays by specifying a half-open range with two indices separated by a colon. based on `b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}`, `b[1:4] == []byte{'o', 'l', 'a'}` will be true, sharing the same storage as b. More info: https://blog.golang.org/go-slices-usage-and-internals
 - In Go, the reference types need to be dereferenced. `*` in front of the variable name allows to dereference the pointers. More info: https://www.golang-book.com/books/intro/8
    - OK, not reference types maybe. Go apperently doesn't have a concept of a reference type, see https://www.tapirgames.com/blog/golang-has-no-reference-values and http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
 - In order to exit from a command line app, you use `os.Exit(int32)`. More info: https://stackoverflow.com/a/18969976/463785
