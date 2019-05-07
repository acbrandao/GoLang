package main
import "fmt"

func main(){
	var age int =20

	fmt.Println("----- Logic Operations -----")

	if age > 18 {
		fmt.Println("Voting Age YES ",age)
	} else {
		fmt.Println("No Voting age ",age)
	}

	switch age {
	case 16:fmt.Println("In High School ",age)
	case 18:fmt.Println("a New car ",age)

	case 20:fmt.Println("in college ",age)
	default: fmt.Println("A Default case",age)
	
	}

}