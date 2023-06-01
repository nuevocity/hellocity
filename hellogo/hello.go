package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hola server\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("", nil)
}
