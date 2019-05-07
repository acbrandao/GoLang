package main
import "fmt"

func main(){
	fmt.Println("----- Loops -----")

	for i:=1; i <= 5 ; i++ {
		fmt.Println(i)
	}

	fmt.Println("----- While -----")
	i:=1
	for i<=5 {
		fmt.Println(i)
		i++
	}

	for i:=5; i >0 ; i-- {
		for j:=1; j < 3 ; j++ {
		fmt.Println(i,j)
		}
	}

}