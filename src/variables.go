package main

import (
	"fmt"
)

func main() {
	var a int =5
	var b float32 = 4.32
	const pi float64 = 3.14159
	x,y:=5,6

	fmt.Println("Variables!")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(y)

	//Arithmetic ops
	fmt.Println("x + y =",x+y)
	fmt.Println("x - y =",x-y)
	fmt.Println("x * y =",x*y)
	fmt.Println("x / y =",x/y)
	fmt.Println("x mod y =",x%y)

	//Boolean
	bool1:=true
	var bool2 bool=false
	fmt.Println("Boolean")
	fmt.Println("Boolean bool1 =",bool1)
	fmt.Println("Boolean bool2 =",bool2)
	fmt.Println("Boolean bool1 && bool2 =",bool1 && bool2)

}