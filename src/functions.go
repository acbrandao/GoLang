package main
import "fmt"

func main(){
	fmt.Println("----- Functions -----")

	x,y:=5,6

	fmt.Println("x +y =",add(x,y) )
	fmt.Println("Factorial 5 =",fact(5) )
}

func add (a,b int) int {
	return a+b
}

func fact(num int) int {
	if num==0 {
		return 1
	}

	return num * fact(num-1)
}