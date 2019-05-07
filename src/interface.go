package main
import (
	"fmt"
	"math"
)

func main(){
	fmt.Println("-----Interfaces -----")

	rect:=Rectangle{50,60}
	circ :=Circle{7}

	fmt.Println(" Rectangle ",rect)
	fmt.Println(" Circle ",circ)

	fmt.Println(" Rectangle Area ",getArea(rect))
	fmt.Println(" Circle Area",getArea(circ))

}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
} 

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius,2)
}

func getArea(shp Shape) float64 {
	return shp.area()
}