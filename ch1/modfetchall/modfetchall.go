//A concurrent program. Note: Concurrency is not parallelism - https://talks.golang.org/2012/waza.slide#57
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() { //Main itself runs in a goroutine
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs	elapsed\n", time.Since(start).Seconds()) //Note the compiler converts \n to \r\n for windows
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch return; Sprint prints using default formats.
		return
	}
	timestamp := fmt.Sprintf("%d_%02d_%02dT%02d_%02d_%02d", start.Year(), start.Month(), start.Day(),
		start.Hour(), start.Minute(), start.Second())
	fileName := strings.Join([]string{timestamp, "txt"}, "")
	//ch <- fmt.Sprintf(fileName)
	fo, _ := os.Create(fileName)
	w := bufio.NewWriter(fo)
	nbytes, err := io.Copy(w, resp.Body) //Discarded but bytes are counted
	resp.Body.Close()
	w.Flush()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%0.2fs %7d %s", secs, nbytes, url)
}
