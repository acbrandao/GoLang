package main
import "fmt"

func main(){
	x:=10 

	fmt.Println("---------- Pointers in Go ------------")
	fmt.Println("Value of x = ", x)
	fmt.Println("Value of &x = ", &x)

	changeVal(&x)
	fmt.Println("Value of x = ", x)
}

func changeVal(x *int) {
	*x=7
}