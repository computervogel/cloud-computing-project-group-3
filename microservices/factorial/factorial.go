package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var n uint64
	n = 65
	var result uint64
	result = 1
	factorial := "! = "
	//the bad efficiency for printing is intentional, in order to see the autoscaler work properly
	for i := uint64(0); i <= n; i++ {
		result = result * i
		printer := fmt.Sprint(i, factorial, result)
		if(i == 0) {
			result = 1
		}
		fmt.Fprintln(w, printer)
	}
	
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5555", nil))
}
