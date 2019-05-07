package main

import (
	 "fmt"
	 "errors"
	 "math"
)

type person struct {
   name string
   age int

}

func main() {
	x := 5 
	var y int =7
	var a[5]int
	a[2]=7 

	p:=person{name:"john", age:23}

	fmt.Println(x)
	fmt.Println(&x) 

	 fmt.Println(p.name)
	

  for i:=0; i < 5; i++ {
	fmt.Println(i)
  }
 
  
	fmt.Println(x,y )
	fmt.Println(a)
	fmt.Println("hello, world")

    if x > 6 {
		fmt.Println(" > 6 ")
	}
	result := sum(2,3)
	fmt.Println(result)

    res64, err:=sqrt(8)

 if (err != nil) {
	fmt.Println(err)
} else
{
	fmt.Println(res64)
 }

  fmt.Println("Before inc")
  fmt.Println(x)
  inc(&x)
  fmt.Println("After inc")
  fmt.Println(x)
}

func inc(x *int) {
	*x++
}

func sum(x int,y int) int {
	return x + y
}

func sqrt(x float64) (float64,error) {
  
	if x < 0 {
		return  0, errors.New("Undefined for neg")
	}

	return math.Sqrt(x),nil 
}