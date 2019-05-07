package main
    
import (
	"fmt"
    . "github.com/kkdai/pubsub"
)
    
func main() {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("topic1")
	c2 := ser.Subscribe("topic2")
	ser.Publish("test1", "topic1")
	ser.Publish("test2", "topic2")
	fmt.Println(<-c1)
	//Got "test1"
	fmt.Println(<-c2)
	//Got "test2"


    // Add subscription "topic2" for c1.          
	ser.AddSubscription(c1, "topic2")

    // Publish new content in topic2
	ser.Publish("test3", "topic2")

	fmt.Println(<-c1)
	//Got "test3"
	
    // Remove subscription "topic2" in c1
	ser.RemoveSubscription(c1, "topic2")
	
    // Publish new content in topic2
	ser.Publish("test4", "topic2")
    
	select {
	case val := <-c1:
		fmt.Printf("Should not get %v notify on remove topic\n", val)
		break
	case <-time.After(time.Second):
	    //Will go here, because we remove subscription topic2 in c1.         
		break
	}
}