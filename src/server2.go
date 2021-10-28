package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("127.0.0.1:9999", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q", r.URL.Path)
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count = %d", count)
	mu.Unlock()
}
