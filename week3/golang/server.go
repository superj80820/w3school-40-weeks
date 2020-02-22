package main

import (
  "fmt"
  "net/http"
)

func cpu(w http.ResponseWriter, r *http.Request) {
  sum := 1
  for sum < 100000 {
      sum ++
      fmt.Println(sum)
  }
  fmt.Fprintln(w, "Hello, world.")
}

func io(w http.ResponseWriter, r *http.Request) {
  fmt.Println("let't go")
  _, err := http.Get("http://localhost:9999/sleep")
  if err != nil {
      fmt.Fprintln(w, "shit")
  }
  fmt.Println("done")
  fmt.Fprintln(w, "Hello, world.")
}

func main() {
  http.HandleFunc("/cpu", cpu)
  http.HandleFunc("/io", io)
  fmt.Println("Golang sever listening on port 1234!")
  http.ListenAndServe(":1234", nil)
}