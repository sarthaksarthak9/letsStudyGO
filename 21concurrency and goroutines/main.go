// concurrency - you have multiple task and you do it one at a time. You can stop and continue latter
// parallelism - do all task together
// gorountines - it is used to achieve parallelism.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // these are pointers

func main() {
	// go greeter("Hello") // way to implement gorountine to
	// greeter("world")
	websitelist := []string{
		"https://google.com",
		"https://youtube.com",
		"https://go.dev",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web) // getStatusCode alone taking a lot of time to execute completly.
		// lets solve the gorountine problem with the help of sync package

		// check docs of sync package
			// you find WaitGroup which is a adv version of time.sleep
				// In WaitGroup we have Add(), Done(), Wait() methods

		wg.Add(1) // it keep on adding go routine number that how many freinds are been out there.
		// In this case we have one as we have one call only.
	}

	wg.Wait() // this always come in the end. It does not allow main method to finish.
}

// In above (greeter()) we don't instruct the thread when to come back
// so codes end but we don't see the thread. One way to resolve use time package
// ideal way to resolve use sync package

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond) // not a ideal way to resolve the issue
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() // waitgroup pass a signal done

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPs in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
