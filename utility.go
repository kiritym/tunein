package main

import (
      "fmt"
      "net"
      "math"
      "time"
)

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

func calculateTimeDiff() int{
  fmt.Println("calculation")
  fmt.Println("song start time: ", songStartTime)
  fmt.Println("length ...", songLength)
	diff := time.Now().Sub(songStartTime)
	diff_time := int(math.Ceil(diff.Seconds()))
	fmt.Println("song length: ", songLength)
	waiting_time = songLength - diff_time
	fmt.Println("waiting time: ", waiting_time)
	return waiting_time
}
