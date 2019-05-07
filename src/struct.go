package main
import "fmt"

type rect struct {
	height float64
	width float64
}

func main(){
	fmt.Println("----- Struct in Go -----")

	rect1 := rect{height:10, width:5}
	rect2 := rect{10,6}

	fmt.Println("Rect 1: ", rect1)
	fmt.Println("Rect2: ", rect2)

	fmt.Println("Rect2 Area: ", rect1.area() )

}

func (r *rect) area() float64 {
	return r.height*r.width
}