package main

import (
	 "fmt"
	 "time"
	 "sync"
	 "runtime"
)

var wg=sync.WaitGroup{}
var counter=0
var m=sync.RwMutex{}

func main() {

	var msg="hello"

	var threads= runtime.GOMAXPROCS(-1);
	fmt.Printf("%d GoRoutines Trheads %v \n",ts(),threads )

	for i:= 0; i < 10 ; i++ {
		wg.Add(2)
		m.Rlock()
		go sayHello(msg)
		m.Lock()
		go increment()
	}

	wg.Wait()

fmt.Printf("%d GoRoutines END. \n",ts() )
}
 



func sayHello(msg string) {

	fmt.Printf("%d  SayHello msg=%s #%d  \n",ts(),msg,counter )
	m.Runlock()
	wg.Done()
}

func increment() {
	
	counter++
	m.Unlock()
	wg.Done()
}

func ts() int64 {
    return time.Now().UnixNano() / 1e6
}
