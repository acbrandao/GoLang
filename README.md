# GoLang

![alt text](https://github.com/acbrandao/GoLang/blob/master/img/golang_ride.png "GoLang Mascot")
Go Lang code and scripts. Mostly for my learning and educational purposes.
Go [[https://golang.org/]] is a compiled programming language developed at Google and popular for server side applications.
The language is cross-platform and features a number of modern capabilities over traditional languages like C/C++ or Java. It's primarily intended to be run on server side applications, which makes its typical use case on par with languages like Java, C++ and Python.

_Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson. Go is a statically typed compiled languals
ge in the tradition of C, with memory safety, garbage collection, structural typing, and CSP-style concurrent programming features added. The compiler, tools and source code are all free and open source_ (source Wikipedia)

# Sample Golang program

Below is the simplest Go lang program , saved as **hello.go**

```go
package main
import "fmt"

func main() {
        fmt.Printf("Hello\n")
}
```

In order to test run this command issue:

```
go run hello.go
```

If there are no syntax issues it shoudl run cleanly

# Compiling in Windows or Linux

To create a windows executable (.exe) issue th ecommand below

```
> GOOS=windows GOARCH=386 go build -o hello.exe hello.go
```

For linux

```
GOOS=linux go build -o hello hello.go
```

# Reduce Go binary sizes

Go binaries tend to be large , even for smallish source apps, this is primarily because Go applications contain a runtime which provides garbage collection, reflection and many other features (which your program might not really use, but it's there). And the implementation of the fmt package which you used to print the "Hello World" text (plus its dependencies). So most binaries start off in the 2MB to 6MB in size depending on what import (packages you include).

Some approaches to reduce binary size (mostly Linux)
**use the -s and -w linker flags** to strip the debugging information like , reduces about 28% of binary size
`go build -ldflags "-s -w" -o hello hello.go`

Next run **upx** a Linux Ultimate Packer for eXecutables utilitity , which essentially zips the executable into a running format, this shoudl yield another 15% reduction in size.

```
go build -ldflags "-s -w" -o hello hello.go
upx --brute hello
```

For comparison on _go version go1.10.4 linux/amd64_

| Program (Source)       | Adjustment  | Size (MB) |
| ---------------------- | ----------- | --------- |
| hello.go               | source code | 87 bytes  |
| hello                  | executable  | 2.0 MB    |
| hello -ldflags "-s -w" | execuable   | 1.2 MB    |
| hello upx --brute      | executable  | 350kb     |

# Go References wen sites and links

- [Go Lang :: Official Site ](https://www.golang.com)
- [Learning Go Lang ](https://blog.learngoprogramming.com/)
- [Golang Online Book ](https://www.golang-book.com/books/intro)
- [Reddit Go Programming Language ](https://www.reddit.com/r/golang/)
- [GoLang Code Snippets ](https://golangcode.com/)
- [GoLang bot Tutorial ](golangbot.com)
- [The Go Programming Language Blog](https://blog.golang.org/)
- [Applied Go Beyond Tutorials](https://appliedgo.net/)
