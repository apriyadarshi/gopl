package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	//"strconv"
	"strings"
	"sync"
	"time"
)

//working directory
var wd string
var mu sync.RWMutex

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
		wd = "/users/skywalker/go/src/github.com/gopl"
		handleFTPCtrlConn(conn)
	}

	//data conn
}

func openDataConn(port string, dataListener *net.Listener, dataConn *net.Conn) {
	fmt.Println("openDataConn called")
	if port == "" {
		port = "20"
	}
	addr := "localhost:" + port
	//fmt.Println("attempting to listen on: " + addr)

	var errListen error
	var listener net.Listener
	if *dataListener != nil {
		listener = *dataListener
	} else {

		listener, errListen = net.Listen("tcp", addr)
		if errListen != nil {
			log.Fatal(errListen)
		}
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Print(err)
	}
	mu.Lock()
	*dataConn = *(&conn)
	*dataListener = *(&listener)
	mu.Unlock()
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
	var dataListener net.Listener
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

		resp := ""
		switch strings.ToLower(cmd) {
		case "user":
			resp = "331 Enter Password \r\n"
		case "pass":
			resp = "230 Login Successful \r\n"
		case "syst":
			resp = "215 LINUX Ubuntu 16.04\r\n"
		case "feat":
			resp = "211 No extra features. Use PASV for ls.\r\n"
		case "pwd":
			resp = fmt.Sprintf("257 %s \r\n ", wd)
		case "pasv":
			resp = "227 Entering Passive Mode (127,0,0,1,0,20) \r\n"
			mu.Lock()
			fmt.Println(dataConn)
			if dataConn != nil {
				dataConn.Close()
			}
			go openDataConn("20", &dataListener, &dataConn)
			mu.Unlock()
			time.Sleep(2000 * time.Millisecond)
		case "list":
			var dir string
			var dataResp string
			if len(raw) == 1 {
				dir = wd
			} else {
				dir = filepath.Join(wd, raw[1])
			}
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
			mu.Lock()
			if dataConn != nil {
				fmt.Fprint(dataConn, dataResp)
				dataConn.Close()
				resp = "226 Directory send ok \r\n"

			} else {
				resp = "500 Dataconn not established \r\n"
			}
			mu.Unlock()

		case "cwd":
			dir := filepath.Join(wd, raw[1]) //assuming incremental paths are given
			file, err := os.Stat(dir)
			if err != nil {
				resp = fmt.Sprintf("500 server error : %s\r\n", err)
				break
			}
			if os.IsNotExist(err) {
				resp = fmt.Sprintf("550 Dir %s doesn't exist \r\n", raw[1])
				break
			}
			if !file.IsDir() {
				resp = fmt.Sprintf("550 %s is a file. \r\n", raw[1])
				break
			}

			wd = dir
			resp = fmt.Sprintf("250 Working directory changed to %s \r\n", wd)
		case "retr":
			fName := filepath.Join(wd, raw[1])
			file, err := os.Stat(fName)
			if err != nil {
				resp = fmt.Sprintf("500 server error : %s\r\n", err)
			}
			if os.IsNotExist(err) {
				resp = fmt.Sprintf("550 File %s doesn't exist \r\n", raw[1])
				break
			}
			if file.IsDir() {
				resp = fmt.Sprintf("550 %s is a directory. Can retrieve files only. \r\n", raw[1])
				break
			}
			fmt.Fprintf(conn, "150 Fetching file. Transferring %s bytes ....\r\n", file.Size())
			f, errOpen := os.Open(fName)
			if errOpen != nil {
				resp = fmt.Sprintf("550 open error: %s\r\n", err)
			}
			mu.RLock()
			n, errCopy := io.Copy(dataConn, f)
			if errCopy != nil {
				resp = fmt.Sprintf("550 retr error: %s\r\n", err)
				f.Close()
			}
			f.Close()
			dataConn.Close()
			mu.RUnlock()
			fmt.Fprintf(conn, "226 File sent ok. %d bytes transferred. \r\n", n)
			//type "200 A N \r\n"
		case "quit":
			exit = true
			resp = "221 Goodbye.\r\n"
		default:
			resp = "501 method not implmented \r\n"
		}
		if resp != "" {
			fmt.Fprintf(conn, resp)
		}
		if exit {
			break
		}
	}
}
