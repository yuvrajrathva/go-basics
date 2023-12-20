package main
import (
	"fmt"
	"sync"
	"net/http"
)

var wg sync.WaitGroup 
var mut sync.Mutex
var signals = []string{"start"}
func main(){
	websiteLists:=[]string{
		"https://google.com",
		"https://prometeo.in",
		"https://fb.com",
	}

	for _, w := range websiteLists{
		go getStatusCode(w);
		wg.Add(1) // responsible for saying how many go routines out there here it is 1
	}
	wg.Wait() // responsible to tell main() to wait do not complete main() function // Wait blocks until the WaitGroup counter is zero

	fmt.Println(signals)
}

func getStatusCode(w string){
	defer wg.Done() // Done decrements the WaitGroup counter by one.
	res, err := http.Get(w)
	if err!=nil {
		fmt.Println("this is error")
	} else { 
		mut.Lock()
		signals = append(signals, w)
		mut.Unlock()
		fmt.Println(res.StatusCode) 
	}
}
