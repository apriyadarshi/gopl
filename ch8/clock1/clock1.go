package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:"+os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		now := time.Now()
		zone, _ := now.Zone()
		_, err := io.WriteString(c, zone+": "+now.Format("15:04:05"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}
