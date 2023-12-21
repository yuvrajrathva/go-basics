package main
import (
	"fmt"
	"sync"
)
func main(){
	myCh :=  make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup){ // <-chan mean receiver
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)

		defer wg.Done()
	}(myCh, wg)

	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup){ // chan<- mean sender
		myCh <- 5
		close(myCh)
		
		defer wg.Done()
	}(myCh, wg)
	wg.Wait()
}