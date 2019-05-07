# GoLang

Go Lang code and scripts. Mostly for my learning and educational purposes.
Go [[https://golang.org/]] is a compiled programming language developed at Google and popular for server side applications.
The language is cross-platform and features a number of modern capabilities over traditional languages like C/C++ or Java


*Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson. Go is a statically typed compiled language in the tradition of C, with memory safety, garbage collection, structural typing, and CSP-style concurrent programming features added. The compiler, tools and source code are all free and open source* (source Wikipedia)

# Sample Golang program

Below is the simplest  Go lang program , saved as **hello.go**

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

# Compiling in Windows

To create a windows executable (.exe) issue th ecommand below

```
> GOOS=windows GOARCH=386 go build -o hello.exe hello.go
```

# Go References wen sites and links
- [Go Lang :: Official Site ](https://www.golang.com)
- [Learning Go Lang ] (https://blog.learngoprogramming.com/ )
- [Reddit Go Programming Language] (https://www.reddit.com/r/golang/ ) 
- [GoLang Code Snippets] (https://golangcode.com/)
- [GoLang bot Tutorial ] (golangbot.com  )
- [The Go Programming Language Blog] ( https://blog.golang.org/ )
- [Applied Go  Beyond Tutorials] (https://appliedgo.net/)


