package main

import ("fmt"
	"net/http")

func home_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Server was made by Go")
}

func main() {
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":80", nil)
}

