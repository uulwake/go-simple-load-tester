package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const DEFAULT_USERS = 10

func main() {
	var wg sync.WaitGroup

	var numberOfCalls int
	flag.IntVar(&numberOfCalls, "c", DEFAULT_USERS, "number of concurrent users")
	flag.Parse()

	fmt.Printf("Concurrent Users: %d\n\n", numberOfCalls)

	for i := 0; i < numberOfCalls; i++ {
		wg.Add(1)
		go HttpCall(&wg)
	}

	wg.Wait()
}

func HttpCall(wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get("http://localhost:3030")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Response: %s", string(body))
}
