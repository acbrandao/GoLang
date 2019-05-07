package main
import "fmt"

func main(){
	fmt.Println("----- Maps in Go -----")

	studentAge:= make(map[string] int)

	studentAge["Janice"]=22
	studentAge["Bob"]=24
	studentAge["Pam"]=22
	studentAge["Rich"]=24
	studentAge["Wanda"]=22
	studentAge["Carlos"]=33

	fmt.Println(studentAge)
	fmt.Println("Student count:",len(studentAge))
	fmt.Println("Carlos Age: ",studentAge["Carlos"])


}