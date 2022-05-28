package main

import ( 
  "fmt" 
  "os"
  "net/http" ) 

func handler(w http.ResponseWriter, r *http.Request) {
  hostName, _ := os.Hostname() 
  _, _ = fmt.Fprintf(w, "<h1>Server was created by using Golang<h2> Version 1<h1>Hostname = %s\n", hostName)
}
func main() { 
  http.HandleFunc("/", handler) 
  http.ListenAndServe(":80", nil)
}
