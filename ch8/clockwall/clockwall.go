package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(os.Args[1:])
	var conns []net.Conn
	for _, v := range os.Args[1:] {
		parts := strings.Split(v, "=")
		_, add := parts[0], parts[1]

		conn, err := net.Dial("tcp", add)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		conns = append(conns, conn)
	}

	now := time.Now()
	zone, _ := now.Zone()
	l := len([]byte(zone + ": " + now.Format("15:04:05")))
	p := make([]byte, l, l)
	for {
		for i, conn := range conns {
			conn.Read(p)
			if i == 0 {
				fmt.Printf("\r%s  ", string(p))
			} else {
				fmt.Printf("%s  ", string(p))
			}
		}
	}
}
