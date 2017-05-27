package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin) //doc is a pointer to html node
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {

	if n != nil {

		if n.Type == html.TextNode {
			fmt.Println(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Data != "script" && c.Data != "style" {
				visit(c)
			}
		}
	}
}
