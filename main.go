package main

import (
      "fmt"
      "flag"
      "net/http"
      "golang.org/x/net/websocket"
      "html/template"
      "os"
      "net"
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

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}


func handler(ws *websocket.Conn) {
  wchan := make (chan string)
  wsocket.Add(ws, wchan)
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
  go playRadio(wsocket)
  http.Handle("/radio", websocket.Handler(handler))
  http.HandleFunc("/", rootHandler)
  err := http.ListenAndServe(":" + PORT, nil)
  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }

}
