package main

import (
        "io/ioutil"
        "strings"
        "fmt"
        "strconv"
        "os/exec"
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
  out, err := exec.Command("./script/find_song_duration.sh", songName).Output()
    if err != nil {
			fmt.Fprintf(os.Stderr, "Error in script command %s\n", err.Error())
    }
    duration, _ := strconv.Atoi(strings.TrimSpace(string(out)))
		return duration
}
