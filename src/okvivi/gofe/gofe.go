package main

import (
  "fmt"
  "net/http"
  "okvivi/models"
)

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, models.HelloWorld)
}

func main() {
  fmt.Println("Listening on 3001")

  http.HandleFunc("/", hello)
  err := http.ListenAndServe(":3001", nil)
  if err != nil {
    fmt.Println(err)
  }
}
