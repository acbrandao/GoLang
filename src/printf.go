package main
import "fmt"

func main(){
	const pi float64 = 3.1415926
	x:=5
	isbool := true
	var name string ="Bob "
	last,color :="Ross", "blue"

	fmt.Println("---- PrintF -----")

	fmt.Printf("Pi as float %f \n ", pi);
	fmt.Printf("Type of boolean  %T \n ", isbool);
	fmt.Printf("Pi as boolean  %t \n ", pi);
	fmt.Printf("X as integer %d \n ", x);

	fmt.Println("String length  \n ", len(name) );
	fmt.Println( name+" "+last+" Paints with  color:"+color );

	fmt.Printf("%f \n",pi)
	fmt.Printf("%.3f \n",pi)
	fmt.Printf("%T \n",pi)
	fmt.Printf("%T \n",name)
	fmt.Printf("%t \n",isbool)

	fmt.Printf("%b \n",25)
	fmt.Printf("%c \n",35)
	fmt.Printf("%x \n",15)
	fmt.Printf("%e \n",pi)

}
