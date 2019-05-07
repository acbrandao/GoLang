package main
import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main(){
	fmt.Println("----- File I/O -----")

	 file,err :=os.Create("sample.txt")

	 if err != nil{
		 log.Fatal(err)
	 }

	 file.WriteString("File created using GoLang")

	 file.Close()

	 fmt.Println("Reading smaple.txt")
	 stream,err :=ioutil.ReadFile("sample.txt")
	 if err!=nil{
		 log.Fatal(err)
	 }

	 s1 :=string(stream)
	 fmt.Println(s1)

}