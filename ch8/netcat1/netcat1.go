package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", os.Args[1]+":"+os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	now := time.Now()
	zone, _ := now.Zone()

	l := len([]byte(zone + ": " + now.Format("15:04:05")))
	p := make([]byte, l, l)
	for {
		conn.Read(p)
		fmt.Printf("\r%s", string(p))
	}
}
