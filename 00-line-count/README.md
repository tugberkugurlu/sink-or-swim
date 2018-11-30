## Run

```bash
go run main.go
```

## Command line arguments

[Go by Example: Command-Line Arguments](https://gobyexample.com/command-line-arguments)

 - import `os` package
 - `os.Args` will get us access to command line args. These will be raw command-line arguments which will include path to the program as the first value. `os.Args[1:]` will give us arguments to the program.

## Side Learnings

 - It is possible to slice the arrays by specifying a half-open range with two indices separated by a colon. based on `b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}`, `b[1:4] == []byte{'o', 'l', 'a'}` will be true, sharing the same storage as b. More info: https://blog.golang.org/go-slices-usage-and-internals
