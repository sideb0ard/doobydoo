package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	doobyreply := ""
	if rand.Intn(10) > 8 {
		doobyreply = "What a Fool Believes"
	} else {
		doobyreply = "I Keep Forgettin'"
	}
	fmt.Fprintf(w, doobyreply+"\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
