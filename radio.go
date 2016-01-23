package main

import (
        "time"
        "strings"
        "io"
        "io/ioutil"
        "fmt"
        "os"
  )

func playList() []string{
  var songList []string
  files, _ := ioutil.ReadDir("music")
  i := 0
    for _, f := range files {
          if strings.Contains(f.Name(), ".mp3"){
            fname := "music/" + f.Name()
            songList = append(songList, fname)
          }
          i++
    }
  return songList
}

func findSongDuration(songName string) int{
  //TODO - using "ffmpeg" for each song calculate the song length in sec
  return 10
}

func copy(source io.Reader, wsocket Websockets, size int64, songName string) {
  content := make([]byte, size)
	n, _ := io.ReadFull(source, content)
	fmt.Println("reader size: ", n)
	wsocket.Write(content)
}

func playFile(fileName string, wsocket Websockets, songLength int){
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
	copy(f, wsocket, fi.Size(), songName)
}

func playRadio(wsocket Websockets){
  for {
  	playList := playList()
  	for _, fileName := range playList {
  		songlength := findSongDuration(fileName)
  		playFile(fileName, wsocket, songlength)
  		time.Sleep( time.Duration(songlength) * time.Second)
  	}
  }
}
