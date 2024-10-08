// use mutex where multiple gorountine access a single memory.
// mutex provide a lock over the memory.
// readwrite Mutex - anyone can read but not write. If writing not complete don't allow to read.
// concurrency - you have multiple task and you do it one at a time. You can stop and continue latter.
// parallelism - do all task together.

// gorountines - it is used to achieve parallelism.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // these are pointers
var signals = []string{"test"}

var mut sync.Mutex // pointer

func main() {
	
	websitelist := []string{
		"https://google.com",
		"https://youtube.com",
		"https://go.dev",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web) 

		wg.Add(1)
	}

	wg.Wait() 
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPs in endpoint")
	} else {
		mut.Lock() // where we are performing an operation like read and write
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
