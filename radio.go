package main

import (
        "time"

  )

func playList() []string{
  var songList []string
  //TODO - Read from file directory and add into the array
  return songList
}

func findSongDuration(songName string) int{
  //TODO - using "ffmpeg" for each song calculate the song length in sec
  return 10
}

func playFile(songName string, wsocket Websockets, songLength int){
  //TODO - Read the music file and write into websocket
}

func playRadio(wsocket Websockets){
  for {
  	playList := playList()
  	for _, songName := range playList {
  		songlength := findSongDuration(songName)
  		playFile(songName, wsocket, songlength)
  		time.Sleep( time.Duration(songlength) * time.Second)
  	}
  }
}
