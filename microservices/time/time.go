package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	result := "Current date and time is: " + dt.String()
	fmt.Fprintln(w, result)
	
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
