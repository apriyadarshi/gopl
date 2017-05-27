package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	//"os"
	"path/filepath"
	//"strconv"
	"strings"
	"time"
)

//working directory
var wd string

func main() {

	//ctrl conn
	listener, err := net.Listen("tcp", "localhost:21")

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		wd = "/home/skywalker/Documents/go/src/gopl"
		handleFTPCtrlConn(conn)
	}

	//data conn
}

func openDataConn(port string, dataConn *net.Conn) {
	fmt.Println("openDataConn called")
	if port == "" {
		port = "20"
	}
	addr := "localhost:" + port
	//fmt.Println("attempting to listen on: " + addr)
	listener, errListen := net.Listen("tcp", addr)
	if errListen != nil {
		log.Fatal(errListen)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Print(err)
	}

	*dataConn = *(&conn)
}

func handleFTPDataConn(conn net.Conn) {
	defer conn.Close()

	fmt.Fprint(conn, "220 Login successful \r\n")
	b := make([]byte, 1024)
	conn.Read(b)

	/*Read exact number of bytes that are filled, else invalid characters will come at end*/
	n := bytes.IndexByte(b, 0)
	fmt.Printf("dataconn n: %d\n", n)
	fmt.Printf("dataconn bytes: %s\n", string(b[:n]))

}

func handleFTPCtrlConn(conn net.Conn) {
	defer conn.Close()
	//bufio.NewReader(conn).ReadString(delim)
	fmt.Fprint(conn, "220 Login successful \r\n")

	var exit bool = false
	var dataConn net.Conn
	for {
		b := make([]byte, 1024)
		conn.Read(b)

		/*Read exact number of bytes that are filled, else invalid characters will come at end*/
		n := bytes.IndexByte(b, 0)
		fmt.Printf("n: %d\n", n)
		fmt.Printf("bytes: %s\n", string(b[:n]))

		//trimspace used to remove the last CRLF present in standard FTP/telnet commands
		raw := strings.Split(strings.TrimSpace(string(b[:n])), " ")

		cmd := raw[0]
		//dir := filepath.Join(wd, fName)

		var resp string
		switch strings.ToLower(cmd) {
		case "user":
			resp = "331 Enter Password \r\n"
		case "pass":
			resp = "230 Login Successful \r\n"
		case "syst":
			resp = "215 LINUX Ubuntu 16.04\r\n"
		case "pasv":
			resp = "227 Entering Passive Mode (127,0,0,1,0,20) \r\n"
			go openDataConn("20", &dataConn)
			//PORT 127,0,0,1,148,80
			/*argraw := strings.Split(raw[1], ",")
			fmt.Println(argraw)
			portLBits, _ := strconv.Atoi(argraw[5])
			portHBits, _ := strconv.Atoi(argraw[4])
			serverDTPport := portHBits*256 + portLBits

			/*if err != nil {
				resp = fmt.Sprintf("500 error opening data conn : %s\r\n", err)
			} else {
				resp = "200 Data connection open \r\n"
			}*/
			//resp = "200 Data connection open \r\n"
			time.Sleep(1000 * time.Millisecond)
		case "list":
			var dir string
			var dataResp string
			if len(raw) == 1 {
				dir = wd
			} else {
				dir = filepath.Join(wd, raw[1])
			}
			if dataConn != nil {
				fmt.Fprint(conn, "150 Fetching file list \r\n")
				files, err := ioutil.ReadDir(dir)
				if err != nil {
					dataResp = fmt.Sprintf("server: cmd %s : %v", "server", cmd, err)
					break
				}
				for _, file := range files {
					dataResp += file.Name()
					dataResp += " "
				}
				dataResp += "\r\n"
				fmt.Fprint(dataConn, dataResp)
				dataConn.Close()
				resp = "226 Directory send ok \r\n"
			} else {
				resp = "500 Dataconn not established \r\n"
			}
			/*case "cd":
			file, err := os.Stat(dir)
			if err != nil {
				if !file.IsDir() {
					resp = fmt.Sprintf("server error: %s is a file", fName)
				}
			}
			if os.IsNotExist(err) {
				resp = fmt.Sprintf("server error: no dir with name %s", fName)
			}
			resp = fmt.Sprintf("success: current directory : %s", dir)*/
		case "quit":
			exit = true
			resp = "221 Goodbye.\r\n"
		default:
			resp = fmt.Sprintf("%s: %s", cmd, "server: method not supported")
		}
		if resp != "" {
			fmt.Fprintf(conn, resp)
		}
		if exit {
			break
		}
	}
}
