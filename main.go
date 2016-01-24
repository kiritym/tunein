package main

import (
	"flag"
	_ "fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
	_ "os"
	"time"
)

var (
	PORT          string
	wsocket       Websockets
	wsCntrl       Websockets
	waiting_time  int
	songStartTime time.Time
	songLength    int
)

func init() {
	const (
		defaultPort = "8080"
		portMessage = "The port of the Tunein Server"
	)
	flag.StringVar(&PORT, "p", defaultPort, portMessage)
}

type TuneInPage struct {
	MusicUrl       string
	ControlDataUrl string
}

func handler(ws *websocket.Conn) {
	wchan := make(chan string)
	wsocket.Add(ws, wchan)
	<-wchan
	//fmt.Fprintf(os.Stdout, "Music Go Routine exited: %s\n", msg)
}

func ctrlHandler(ws *websocket.Conn) {
	wchan := make(chan string)
	wsCntrl.Add(ws, wchan)
	waiting_time = calculateTimeDiff()
	//fmt.Println("waiting time: ", waiting_time)
	cntrlmsg := ControlMsg{Name: "", Duration: waiting_time, Command: "wait"}
	websocket.JSON.Send(ws, cntrlmsg)
	<-wchan
	//fmt.Fprintf(os.Stdout, "Ctrl Go Routine exited: %s\n", msg)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	hostname := getLocalIP()
	musicUrl := "ws://" + hostname + ":" + PORT + "/radio"
	ctrlDataUrl := "ws://" + hostname + ":" + PORT + "/ctrl"

	t, err := template.ParseFiles("tmpl/index.tmpl")
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Unable to load template file: "+err.Error())
		return
	}
	p := TuneInPage{MusicUrl: musicUrl, ControlDataUrl: ctrlDataUrl}
	t.Execute(w, p)
}

func main() {
	flag.Parse()
	wsocket.Init()
	wsCntrl.Init()
	go playRadio(wsocket, wsCntrl)
	http.Handle("/radio", websocket.Handler(handler))
	http.Handle("/ctrl", websocket.Handler(ctrlHandler))
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
