package main

import (
      "fmt"
      "flag"
      "net/http"
      "golang.org/x/net/websocket"
      "html/template"
      "os"
      "time"
)

var PORT string
var wsocket Websockets
var wsCntrl Websockets
var waiting_time int
var songStartTime time.Time
var songLength int

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

func ctrlHandler(ws *websocket.Conn) {
	wchan := make (chan string)
	wsCntrl.Add(ws, wchan)
  waiting_time = calculateTimeDiff()
	cntrlmsg := ControlMsg{Name: "", Duration: waiting_time, Command: "wait"}
	websocket.JSON.Send(ws, cntrlmsg)
	<- wchan
}


func rootHandler(w http.ResponseWriter, r *http.Request){
  hostname := getLocalIP()
  musicUrl := "ws://" + hostname + ":" + PORT + "/radio"
	ctrlDataUrl := "ws://" + hostname + ":" + PORT + "/ctrl"

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
  wsCntrl.Init()
  go playRadio(wsocket)
  http.Handle("/radio", websocket.Handler(handler))
  http.Handle("/ctrl", websocket.Handler(ctrlHandler))
  http.HandleFunc("/", rootHandler)
  err := http.ListenAndServe(":" + PORT, nil)
  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }

}
