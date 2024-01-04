package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var n int
	n = 100000000
	for i := 2; i <= n; i++ {
		isPrime := true
		
	 	for j := 2; j <= i/2; j++ {
			if i%j == 0 {
		   		isPrime = false
		   		break
			}
	 	}	

	 	if isPrime {
		   fmt.Fprintln(w, i)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6666", nil))
}
