package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//Structure of control message consists of song details
type ControlMsg struct {
	Name     string
	Duration int
	Command  string
}

//Write the song and its info(name, duration etc.) into socket
func writeInSocket(source io.Reader, wsocket, wsCntrl Websockets, size int64, songName string) {
	content := make([]byte, size)
	io.ReadFull(source, content)
	//fmt.Println("reader size: ", n)
	wsocket.Write(content)
	msg := ControlMsg{Name: songName, Duration: 0, Command: "play"}
	wsCntrl.WriteText(msg)
}

//Open the song and write in socket
func sendToSocket(fileName string, wsocket, wsCntrl Websockets, songLength int) {
	songName := strings.SplitAfter(fileName, "/")[1]
	songName = strings.Split(songName, ".")[0]
	fmt.Printf("Name of the song: %s \n", songName)
	f, err := os.Open(fileName)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", fileName, err.Error())
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Error in file info %s: %s\n", fileName, err.Error())
		return
	}
	songStartTime = time.Now()
	//fmt.Println("song strt time: ", songStartTime)
	writeInSocket(f, wsocket, wsCntrl, fi.Size(), songName)
}

//read one song from the playlist and send to read the file and write it into the socket
func playRadio(wsocket, wsCntrl Websockets) {
	for {
		playList := playList()
		for _, fileName := range playList {
			songlength := findSongDuration(fileName)
			//fmt.Println(songlength)
			sendToSocket(fileName, wsocket, wsCntrl, songlength)
			time.Sleep(time.Duration(songlength) * time.Second)
		}
	}
}
