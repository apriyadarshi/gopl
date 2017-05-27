package ch5

import (
	//"fmt"
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {

	var nodes []*html.Node
	if names == nil {
		return nil
	}

	var search func(n *html.Node, tags []string)
	search = func(n *html.Node, tags []string) {
		curr := n
		//fmt.Printf("curr value: %s\n", curr.Data)
		for _, name := range names {
			if n.Data == name {
				nodes = append(nodes, curr)
				break
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			//Search only in element nodes, else a data field of textnode with value h1 will match tagname h1
			if c.Type == html.ElementNode {
				search(c, tags)
			}
		}
	}

	search(doc, names)

	return nodes
}
