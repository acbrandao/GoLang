package main

import (
	 "fmt"
	 "time"
)

func main() {
var msg="hello"

	fmt.Println("hello, GoRoutines")

go func(msg string){
	fmt.Println("GoROutine Anonymous func: ",msg)
}(msg)

msg="Goodbye"
time.Sleep(100* time.Millisecond)

go sayHello()
time.Sleep(100* time.Millisecond)
}
 
  func sayHello() {
	fmt.Println("func sayHello")

  }