package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	fmt.Println("Server is listening on :3000...")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
