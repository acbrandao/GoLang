# GoLang

Go Lang code and scripts. Mostly for my learning

# Sample GoLange program

Below is a sample go lang program , saved as **hello.go**

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

#Compiling in Windows
To create a windows executable (.exe) issue th ecommand below

```
> GOOS=windows GOARCH=386 go build -o hello.exe hello.go
```
