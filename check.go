package main

import (
	"time"
	"fmt"
	"net"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration( 5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachabel, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachabel, \n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
