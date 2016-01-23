package main

import (
        "time"
        "strings"
        "io"
        "fmt"
        "os"
  )

type ControlMsg struct{
	Name string
	Duration int
	Command	string
}


func writeInSocket(source io.Reader, wsocket, wsCntrl Websockets, size int64, songName string) {
  content := make([]byte, size)
	n, _ := io.ReadFull(source, content)
	fmt.Println("reader size: ", n)
	wsocket.Write(content)
  msg := ControlMsg{Name: songName, Duration: 0, Command: "play"}
	wsCntrl.WriteText(msg)
}

func sendToSocket(fileName string, wsocket, wsCntrl Websockets, songLength int){
  songName := strings.SplitAfter(fileName, "/")[1]
	fmt.Printf("Name of the song: %s \n", songName)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", fileName, err.Error())
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in file info %s: %s\n", fileName, err.Error())
		return
	}
  songStartTime = time.Now()
	fmt.Println("song strt time: ", songStartTime)
	writeInSocket(f, wsocket, wsCntrl, fi.Size(), songName)
}

func playRadio(wsocket, wsCntrl Websockets){
  for {
  	playList := playList()
  	for _, fileName := range playList {
  		songlength := findSongDuration(fileName)
      fmt.Println(songlength)
  		sendToSocket(fileName, wsocket, wsCntrl, songlength)
  		time.Sleep( time.Duration(songlength) * time.Second)
  	}
  }
}
