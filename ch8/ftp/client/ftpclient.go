package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, cmd, f := os.Args[1], os.Args[2], os.Args[3]

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	io.WriteString(conn, cmd+" "+f)

	b := make([]byte, 1000, 1000)
	conn.Read(b)
	fmt.Printf("%s\n", string(b))
}
