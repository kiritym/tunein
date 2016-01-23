package main

import (
      "fmt"
      "flag"
      "net/http"
      "golang.org/x/net/websocket"
      "html/template"
      "os"
)

var PORT string
var wsocket Websockets

func init() {
  const (
          defaultPort = "8080"
          portMessage = "The port of the Tunein Server"
        )
  flag.StringVar(&PORT, "p", defaultPort, portMessage)
}

type TuneInPage struct {
	MusicUrl string
	ControlDataUrl string
}

func handler(ws *websocket.Conn) {
  wchan := make (chan string)
  wsocket.Add(ws, wchan)
  <- wchan
}

func rootHandler(w http.ResponseWriter, r *http.Request){
  musicUrl := "ws://" + "localhost" + ":" + PORT + "/radio"
	ctrlDataUrl := "ws://" + "localhost" + ":" + PORT + "/ctrl"

	t, err := template.ParseFiles("tmpl/index.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load template file: " + err.Error())
		return
	}
	p := TuneInPage{MusicUrl : musicUrl, ControlDataUrl : ctrlDataUrl}
	t.Execute(w, p)
}

func main(){
  fmt.Println("Hello Tunein")
  flag.Parse()
  wsocket.Init()
  go playRadio(wsocket)
  http.Handle("/tunein", websocket.Handler(handler))
  http.HandleFunc("/", rootHandler)
  err := http.ListenAndServe(":" + PORT, nil)
  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }

}
