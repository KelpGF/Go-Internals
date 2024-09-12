package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var leakedGoRoutines int32

func leakGoRoutine() {
	atomic.AddInt32(&leakedGoRoutines, 1)
	for {
		// infinite loop to simulate a goroutine leak
	}
}

func leakHandler(w http.ResponseWriter, r *http.Request) {
	go leakGoRoutine()

	fmt.Fprintf(w, "Leaked goroutines")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current number of leaked goroutines: %d", leakedGoRoutines)
}

func main() {
	http.HandleFunc("/leak", leakHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}
