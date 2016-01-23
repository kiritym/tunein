package main

import (
      "fmt"
      "flag"
      "net/http"
      "golang.org/x/net/websocket"
)

var PORT string

func init() {
  const (
          defaultPort = "8080"
          portMessage = "The port of the Tunein Server"
        )
  flag.StringVar(&PORT, "p", defaultPort, portMessage)
}

func handler(ws *websocket.Conn) {
  
}

func rootHandler(w http.ResponseWriter, r *http.Request){

}

func main(){
  fmt.Println("Hello Tunein")
  flag.Parse()
  http.Handle("/tunein", websocket.Handler(handler))
  http.HandleFunc("/", rootHandler)
  err := http.ListenAndServe(":" + PORT, nil)
  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }

}
