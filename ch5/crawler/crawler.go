package main

import (
	"bufio"
	"fmt"
	"gopl/ch5/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if !exists("data") {
		os.Mkdir("data", 0777)
	}
	hostURL, _ := url.Parse(os.Args[1])

	crawl := func(fullURL string) []string {

		fmt.Printf("Found %s\n", fullURL)

		if u, _ := url.Parse(fullURL); sameDomain(u, hostURL) {
			fmt.Printf("Fetching %s\n", fullURL)
			//1. Fetch data
			resp, err := http.Get(fullURL)
			if err != nil {
				fmt.Printf("fetch %s: %s", fullURL, err)
			}

			var relD string
			if esc := u.EscapedPath(); esc == "" {
				relD = "data/"
			} else {
				relD = "data" + u.EscapedPath()
			}

			if relF := relD + "index.html"; !exists(relF) {

				os.MkdirAll(relD, 0777)

				f, err := os.Create(relF)
				if err != nil {
					fmt.Printf("create file %s: %s", relF, err)
				}

				writer := bufio.NewWriter(f)
				if _, err := io.Copy(writer, resp.Body); err != nil {
					fmt.Printf("write html %s: %s", relF, err)
				} else {
					fmt.Printf("HTML written to %s\n", relF)
				}

				writer.Flush()
				resp.Body.Close()
			}

			resp.Body.Close()
		}

		list, err := links.Extract(fullURL)
		if err != nil {
			log.Print(err)
		}

		return list
	}

	breadthFirst(crawl, os.Args[1:])
}

//Calls f for each item in worklist
//Items returned by f are added to worklist
//f is called atmost once for each item
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...) //Append is a variadic function.  adds all returned items
			}
		}
	}
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func sameDomain(u1, u2 *url.URL) bool {

	parts1 := strings.Split(u1.Hostname(), ".")
	parts2 := strings.Split(u2.Hostname(), ".")

	if parts1[len(parts1)-1] == parts2[len(parts2)-1] && parts1[len(parts1)-2] == parts2[len(parts2)-2] {
		return true
	}
	return false
}
