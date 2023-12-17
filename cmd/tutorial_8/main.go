package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	wg      = sync.WaitGroup{}
	results = []string{}
	mu      = sync.Mutex{}
)

func main() {
	websitesList := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	t0 := time.Now()

	for _, web := range websitesList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Println(results)
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Error when fetching the endpoint: %v\n", err)
	} else {
		fmt.Printf("%v status code for %v\n", res.StatusCode, endpoint)

		mu.Lock()
		results = append(results, endpoint)
		mu.Unlock()
	}

	wg.Done()
}
