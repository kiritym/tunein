package main

import (
      "os"
      "golang.org/x/net/websocket"
      "fmt"
      "container/list"
)


type WebsocketElement struct {
	conn *websocket.Conn
	ch chan string
	errCount int
}

type Websockets struct {
	ws *list.List
}

func (w *Websockets) Init() {
	w.ws = list.New()
}

func (w *Websockets) Add(conn *websocket.Conn, ch chan string) {
	w.ws.PushBack(WebsocketElement{conn: conn, ch: ch, errCount: 0})
}


func (w *Websockets) Write (buff []byte) {
	l := w.ws
	fmt.Fprintf(os.Stderr, "Number of connections: %d\n", l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		wsElem := e.Value.(WebsocketElement)
		err := websocket.Message.Send(wsElem.conn, buff)
		if err != nil {
			l.Remove(e)
		}
	}
}

func (w *Websockets) WriteText (cntrl_msg ControlMsg){
	l := w.ws
	for e := l.Front(); e != nil; e = e.Next() {
		wsElem := e.Value.(WebsocketElement)
		err := websocket.JSON.Send(wsElem.conn, cntrl_msg)
		if err != nil {
			l.Remove(e)
		}
	}
}
