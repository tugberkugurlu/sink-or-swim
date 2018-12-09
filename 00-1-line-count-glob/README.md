## Run

```bash
➜  00-1-line-count-glob git:(master) ✗ go build
➜  00-1-line-count-glob git:(master) ✗ ./00-1-line-count-glob
```

## Concurrency

 - [Using goroutines on loop iterator variables](https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables) easy mistake to make but very common in every lang. It was a surprise that Goland didn't give a warning for this.
 - [Go by Example: Range over Channels](https://gobyexample.com/range-over-channels) you can range over a channel and makes it more reliable to get things this way as it handles the closed channels appropriately as far as I experienced. For instance, with a `for` loop and reading from the channel manually by using `<-ch` syntax, I started getting it zeros at the end. This is probably further explain here on why: [Why are there nil channels in Go?](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308).

## Side Learnings

 - [What do 'real', 'user' and 'sys' mean in the output of time(1)?](https://stackoverflow.com/questions/556405/what-do-real-user-and-sys-mean-in-the-output-of-time1/556411#556411)
 - You can create an error object using [`errors.New`](https://golang.org/pkg/errors/#example_New)
 - Go doesn't have the distinctive looping constructs such as `foreach`, `while`. `for` can handle all cases. See [Is there a foreach loop in Go?](https://stackoverflow.com/a/7782507/463785)
 - [Communicating Sequential Processes (CSP)](http://www.usingcsp.com/)
 - [cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
 - [Golang how to import local packages without gopath?](https://stackoverflow.com/questions/17539407/golang-how-to-import-local-packages-without-gopath)
 - [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules)
 - [go build vs go build file.go](https://stackoverflow.com/questions/19234445/go-build-vs-go-build-file-go)
 - [Golang how to import local packages without gopath?](https://stackoverflow.com/questions/17539407/golang-how-to-import-local-packages-without-gopath)
