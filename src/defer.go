package main
import "fmt"

func main(){
	fmt.Println("----- Defer / Recover / Panic -----")

  defer firstRun()
   secondRun() 

 

   fmt.Println("Add numbers ",addn(1,2,3,4,5))
  
   fmt.Println("Divide 3/0 ",div(3,0))
   fmt.Println("Divide 3/5 ",div(3,5))
   demPanic()
 
}

func addn(args ...int) int  {
  sum :=0 
	for _, value:=range args {
     sum +=value
	}

	return sum
}

func div(n1,n2 int)int {

	defer func() {
		fmt.Println(recover())
	}()

	result := n1/n2
	return  result
}

func demPanic() {
	defer func(){
		fmt.Println(recover())
	}()
}

func firstRun(){
	fmt.Println("1st Run! ")
}

func secondRun(){
	fmt.Println("2nd Run! ")
}