package main

import (
	"flag"
	"fmt"
	"net/http"
	"foobar.com/shared"
)

var p = flag.Int("p", 3001, "The port to listen to")

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, shared.HelloWorld)
}

func main() {
	flag.Parse()

	fmt.Println("Listening on", *p)

	http.HandleFunc("/", hello)
	err := http.ListenAndServe(fmt.Sprint(":", *p), nil)
	if err != nil {
		fmt.Println(err)
	}
}
