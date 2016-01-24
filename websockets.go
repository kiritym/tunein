package main

import (
	"container/list"
	"fmt"
	"golang.org/x/net/websocket"
	"os"
)

type WebsocketElement struct {
	conn     *websocket.Conn
	ch       chan string
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

func (w *Websockets) Write(buff []byte) {
	l := w.ws
	i := 0
	fmt.Fprintf(os.Stderr, "Number of connections: %d\n", l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		i++
		wsElem := e.Value.(WebsocketElement)
		err := websocket.Message.Send(wsElem.conn, buff)
		remoteAddr := wsElem.conn.RemoteAddr().String()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%d :: Error Sending music data to %s[ErrCount:%d]: %s\n", i, remoteAddr, wsElem.errCount, err.Error())
			wsElem.ch <- remoteAddr + ":" + err.Error()
			l.Remove(e)
		}
	}
}

func (w *Websockets) WriteText(cntrl_msg ControlMsg) {
	l := w.ws
	for e := l.Front(); e != nil; e = e.Next() {
		wsElem := e.Value.(WebsocketElement)
		err := websocket.JSON.Send(wsElem.conn, cntrl_msg)
		remoteAddr := wsElem.conn.RemoteAddr().String()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Sending text data to %s[ErrCount:%d]: %s\n", remoteAddr, wsElem.errCount, err.Error())
			wsElem.ch <- remoteAddr + ":" + err.Error()
			l.Remove(e)
		}
	}
}
