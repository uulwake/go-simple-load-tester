package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintln(w, "Hello World")

		duration := time.Since(now).Milliseconds()
		fmt.Printf("duration: %d\n", duration)
	})

	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal(err)
	}
}
