package main

import (
	"fmt"
	"sync"
	"time"
)

const ChannelSize = 3

func main() {
	start := time.Now()
	rc := make(chan any, 3)
	wg := &sync.WaitGroup{}

	wg.Add(ChannelSize)

	go searchProvider1(rc, wg)
	go searchProvider2(rc, wg)
	go searchProvider3(rc, wg)

	wg.Wait()
	close(rc)
	for resp := range rc {
		fmt.Println("resp:", resp)
	}

	fmt.Println("took:", time.Since(start))

}

func searchProvider1(rc chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 500)

	rc <- "searchResult - provider1"
	wg.Done()
}

func searchProvider2(rc chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 1500)

	rc <- "searchResult - provider2"
	wg.Done()
}

func searchProvider3(rc chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 900)

	rc <- "searchResult - provider3"
	wg.Done()
}
