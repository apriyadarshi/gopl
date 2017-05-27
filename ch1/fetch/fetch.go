package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = strings.Join([]string{"http://", url}, "") //This is an array/slice literal in Go
		}
		resp, GetErr := http.Get(url)
		if GetErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", GetErr)
			os.Exit(1) //Exit causes the current program to exit with the given status code. Deferred commands dont run
		}
		//Create Output file : Currently not checking whether it exists
		fo, createErr := os.Create("output.txt")
		if createErr != nil {
			fmt.Fprintf(os.Stderr, "Error while creating output file %s", createErr)
			os.Exit(1)
		}

		writer := bufio.NewWriter(fo)
		httpStatus := resp.Status
		fmt.Printf("%s", "Http response status : ", httpStatus)

		_, err := io.Copy(writer, resp.Body)
		resp.Body.Close()
		writer.Flush()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading and writing to output file for url  %s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", "File Successfully copied") //%s - uninterpreted bytes for string or slice
	}
}
