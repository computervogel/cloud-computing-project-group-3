package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var n int
	n = 40
	//the bad efficiency for printing is intentional, in order to see the autoscaler work properly
	for i := 0; i < n; i++ {
		fmt.Fprintln(w, FibonacciRecursion(i))
	}
	
}

func FibonacciRecursion(n int) int {
    if n <= 1 {
        return n
    }
    return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
