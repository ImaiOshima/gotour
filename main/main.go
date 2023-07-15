package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", hiWorld)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatalln(err)
	}
}

func hiWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi,Web World!")
}
