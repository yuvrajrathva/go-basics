package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func getEvents(event string, ch chan<- int, wg *sync.WaitGroup) int {
	url := "https://apiv.prometeo.in/events/event/?type=" + event
	resp, err := http.Get(url)
	defer resp.Body.Close()
	defer wg.Done()

	if err != nil {
		fmt.Printf("error")
		return resp.StatusCode
	}

	ch <- resp.StatusCode
	return resp.StatusCode
}

func main() {
	startTime := time.Now()
	events := []string{"technical", "talk", "pre_events"}

	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, e := range events {
		wg.Add(1)
		go getEvents(e, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println("ch", result)
	}
	fmt.Println("end:", time.Since(startTime))
}
