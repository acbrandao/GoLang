package main
import "fmt"

func main(){
	fmt.Println("----- Arrays -----")
	
	var EvenNum[5] int
	OddNum :=[5]int {1,3,5,7,9}
	EvenNum[0]=0
	EvenNum[1]=2
	EvenNum[2]=3
	EvenNum[3]=6
	EvenNum[4]=8

	fmt.Println("Array[4] value :",EvenNum[4])
	fmt.Println("Array EvenNum :",EvenNum)
	fmt.Println("Array EvenNum :",OddNum)

	for i,value := range EvenNum{
		fmt.Println(value,i)
	}

	numslice:=[]int{5,4,3,2,1}
	sliced :=numslice[0:]


	fmt.Println(sliced)
	

}