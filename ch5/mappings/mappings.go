package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Errorf("Error during parsing: %v\n", err)
		os.Exit(1)
	}
	counts := make(map[string]int)
	visit(&counts, doc)

	fmt.Println("Element Type\tCount")
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}

}

func visit(counts *map[string]int, n *html.Node) {

	if n != nil {

		getMap(counts, n)

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(counts, c)
		}
	}
}

func getMap(counts *map[string]int, n *html.Node) {
	if n != nil && n.Type == html.ElementNode {
		countMap := *counts
		countMap[n.Data]++
	}
}
