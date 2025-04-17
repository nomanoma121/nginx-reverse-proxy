package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// リクエストヘッダーをlogに出力
	fmt.Println("Request Headers:")

	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
	
	fmt.Fprintf(w, "Hello from server-1!\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
